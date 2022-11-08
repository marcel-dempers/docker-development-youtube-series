# Introduction to Sealed Secrets 

Checkout the [Sealed Secrets GitHub Repo](https://github.com/bitnami-labs/sealed-secrets) </br>

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
You will notive it writes a `sealedSecret` to `stdout`. </br>
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
sealedsecret.bitnami.com/mysecret created
```

Now few seconds later, see the secret

```
kubectl -n default get secret
NAME                  TYPE                                  DATA   AGE
mysecret              Opaque                                1      25s
```

## How the encryption key is managed

TODO:  How the encryption key is managed and stored 

TODO:  Set a duration to test `--key-renew-period=<value>`

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

TODO: * backup encryption keys 
TODO:  migrating kubernetes clusters

```
kubectl get secret -n kube-system \
  -l sealedsecrets.bitnami.com/sealed-secrets-key \
  -o yaml \
  >  kubernetes/secrets/sealed-secrets/sealed-secret-keys.key
```

## Migrate your encryption keys to a new cluster 

Delete & Deploy a new Kubernetes cluster

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

restart the controller:
```
kubectl delete pod -n kube-system -l name=sealed-secrets-controller
```

### apply our old sealed secret

```
kubectl apply -f kubernetes/secrets/sealed-secrets/sealed-secret.yaml
```

TODO:  Encrypted sealed secrets across namespaces  https://github.com/bitnami-labs/sealed-secrets#scopes

## Re-encrypting secrets with the latest key

```
kubeseal --re-encrypt <my_sealed_secret.json >tmp.json \
  && mv tmp.json my_sealed_secret.json
```