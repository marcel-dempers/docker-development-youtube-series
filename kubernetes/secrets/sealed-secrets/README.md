# Introduction to Sealed Secrets 

<a href="https://youtu.be/u0qtgUMLua0" title="k8s-sealedsecrets"><img src="https://i.ytimg.com/vi/u0qtgUMLua0/hqdefault.jpg" width="20%" alt="k8s-sealedsecrets" /></a> 

Checkout the [Sealed Secrets GitHub Repo](https://github.com/bitnami-labs/sealed-secrets) </br>

There are a number of use-cases where this is a really great concept. </br>

1) GitOps - Storing your YAML manifests in Git and using GitOps tools to sync the manifests to your clusters (For example Flux and ArgoCD!)

2) Giving a team access to secrets without revealing the secret material.

developer: "I want to confirm my deployed secret value is X in the cluster" </br>

developer can compare `sealedSecret` YAML in Git, with the `sealedSecret` in the cluster and confirm the value is the same. </br>

## Create a kubernetes cluster

In this guide we we'll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name sealedsecrets --image kindest/node:v1.23.5
```

See cluster up and running:

```
kubectl get nodes
NAME                  STATUS   ROLES                  AGE     VERSION
sealedsecrets-control-plane   Ready    control-plane,master   2m12s   v1.23.5
```

## Run a container to work in

### run Alpine Linux: 
```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh
```

### install kubectl

```
apk add --no-cache curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

### install helm

```
curl -o /tmp/helm.tar.gz -LO https://get.helm.sh/helm-v3.10.1-linux-amd64.tar.gz
tar -C /tmp/ -zxvf /tmp/helm.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm
```

### test cluster access:
```
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
sealedsecrets-control-plane   Ready    control-plane,master   3m26s   v1.23.5
```

## Install Sealed Secret Controller 

### download the YAML

In this demo we'll use version [0.19.1](https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.19.1/controller.yaml) of the sealed secrets controller downloaded from the 
[Github releases](https://github.com/bitnami-labs/sealed-secrets/releases) page

```
curl -L -o ./kubernetes/secrets/sealed-secrets/controller-v0.19.1.yaml https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.19.1/controller.yaml

```

### install using Helm

You can also install the controller using `helm`

```
helm repo add sealed-secrets https://bitnami-labs.github.io/sealed-secrets
helm search repo sealed-secrets --versions
helm template sealed-secrets --version 2.7.0 -n kube-system sealed-secrets/sealed-secrets \
> ./kubernetes/secrets/sealed-secrets/controller-helm-v0.19.1.yaml

```
With `helm template` we can explore the YAML and then replace the `helm template` with `helm install`
to install the chart 

### install using YAML manifest

```
kubectl apply -f kubernetes/secrets/sealed-secrets/controller-v0.19.1.yaml
```

### Check the installation

The controller deploys to the `kube-system` namespace by default.

```
kubectl -n kube-system get pods
```

Check the logs of the sealed secret controller 

```
kubectl -n kube-system logs -l name=sealed-secrets-controller --tail -1
```

From the logs we can see that it writes the encryption key its going to use as a kubernetes secret </br>
Example log:

```
2022/11/05 21:38:20 New key written to kube-system/sealed-secrets-keymwzn9
```

## Encryption keys

```
kubectl -n kube-system get secrets
kubectl -n kube-system get secret sealed-secrets-keygxlvg -o yaml
```

## Download KubeSeal

The same way we downloaded the sealed secrets controller from the [GitHub releases](https://github.com/bitnami-labs/sealed-secrets/releases) page,
we'll want to download kubeseal from the assets section 
```

curl -L -o /tmp/kubeseal.tar.gz \
https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.19.1/kubeseal-0.19.1-linux-amd64.tar.gz
tar -xzf /tmp/kubeseal.tar.gz -C /tmp/
chmod +x /tmp/kubeseal
mv /tmp/kubeseal /usr/local/bin/
```

We can now run `kubeseal --help`

## Sealing a basic Kubernetes Secret 

Looks at our existing Kubernetes secret YAML

```
cat kubernetes/secrets/secret.yaml 
```

If you run `kubeseal` you will see it pause and expect input from `stdin`. </br>
You can paste your secret YAML and press CTRL+D to terminate `stdin`. </br>
You will notice it writes a `sealedSecret` to `stdout`. </br>
We can then automate this using `|` characters. </br>

Create a sealed secret using `stdin` :

```
 cat kubernetes/secrets/secret.yaml | kubeseal -o yaml  > kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```

Create a sealed secret using a YAML file:

```
kubeseal -f kubernetes/secrets/secret.yaml -o yaml > kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```

Deploy the sealed secret

```
kubectl apply -f kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```

Now few seconds later, see the secret

```
kubectl -n default get secret
NAME                  TYPE                                  DATA   AGE
mysecret              Opaque                                1      25s
```

## How the encryption key is managed

By default the controller generates a key as we saw earlier and stores it in a Kubernetes secret. </br>
By default, the controller will generate a new active key every 30 days.
It keeps old keys so it can decrypt previous encrypted sealed secrets and will use the active key with new encryption. </br>

It's important to keep these keys secured. </br>
When the controller starts it consumes all the secrets and will start using them </br>
This means we can backup these keys in a Vault and use them to migrate our clusters if we wanted to. </br>

We can also override the renewal period to increase or decrease the value. `0` turns it off </br>

To showcase this I can set `--key-renew-period=<value>` to 5min to watch how it works. </br>

```
apk add nano
export KUBE_EDITOR=nano
```
Set the flag on the command like so to add a new key every 5 min for testing:

```
spec:
  containers:
  - command:
    - controller
    - --key-renew-period=5m

kubectl edit deployment/sealed-secrets-controller --namespace=kube-system
```

You should see a new key created under secrets in the `kube-system` namespace 

```
kubectl -n kube-system get secrets
```

## Backup your encryption keys

To get your keys out for backup purpose, it's as simple as grabbing a secret by label using `kubectl` :

```
kubectl get secret -n kube-system \
  -l sealedsecrets.bitnami.com/sealed-secrets-key \
  -o yaml \
  >  kubernetes/secrets/sealed-secrets/sealed-secret-keys.key
```
This can be used when migrating from one cluster to another, or simply for keeping backups. </br>

## Migrate your encryption keys to a new cluster 

To test this, lets delete our cluster and recreate it. </br>

```
kind delete cluster --name sealedsecrets
kind create cluster --name sealedsecrets --image kindest/node:v1.23.5

# check the cluster
kubectl get nodes

# redeploy sealed-secrets controller
kubectl apply -f kubernetes/secrets/sealed-secrets/controller-v0.19.1.yaml

kubectl -n kube-system get pods

```

### restore our encryption keys

```
kubectl apply -f kubernetes/secrets/sealed-secrets/sealed-secret-keys.key
```

### apply our old sealed secret

```
kubectl apply -f kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```

### see sealed secret status

To troubleshoot the secret, you can use the popular `kubectl describe` command. </br>
Note that we're unable to decrypt the secret. </br>
Why is that ? </br>

We'll this is because the encryption key secrets are read when the controller starts. </br> 
So we will need to restart the controller to that it can read ingest the encryption keys:

```
kubectl delete pod -n kube-system -l name=sealed-secrets-controller
```

## Re-encrypting secrets with the latest key

We can also use `kubeseal --re-encrypt` to encrypt a secret again. </br>
Let's say we want to encrypt with the latest key. </br>
This will re-encrypt the sealed secret without having to pull the actual secret to the client </br>

```
cat ./kubernetes/secrets/sealed-secrets/sealed-secret.yaml \
| kubeseal --re-encrypt -o yaml
```

I can then save this to override the original old local sealed secret file:

```
cat ./kubernetes/secrets/sealed-secrets/sealed-secret.yaml \
| kubeseal --re-encrypt -o yaml \
> tmp.yaml && mv tmp.yaml ./kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```