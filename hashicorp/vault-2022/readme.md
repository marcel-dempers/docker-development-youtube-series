# Hashicorp Vault Guide

<a href="https://youtu.be/2Owo4Ioo9tQ" title="hashicorp-vault"><img src="https://i.ytimg.com/vi/2Owo4Ioo9tQ/hqdefault.jpg" width="20%" alt="introduction hashicorp vault" /></a>

Requirements:

* Kubernetes 1.21
* Kind or Minikube

For this tutorial, I will be using Kubernetes 1.21.
If you are watching the old guide for Kubernetes 1.17, go [here](..\vault\readme.md)

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
cd hashicorp/vault-2022

kind create cluster --name vault --image kindest/node:v1.21.1 --config kind.yaml
```

Next up, I will be running a small container where I will be doing all the work from:
You can skip this part if you already have `kubectl` and `helm` on your machine.

```
docker run -it --rm --net host -v ${HOME}/.kube/:/root/.kube/ -v ${PWD}:/work -w /work alpine sh
```

Install `kubectl`

```
apk add --no-cache curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

Install `helm`

```
curl -LO https://get.helm.sh/helm-v3.7.2-linux-amd64.tar.gz
tar -C /tmp/ -zxvf helm-v3.7.2-linux-amd64.tar.gz
rm helm-v3.7.2-linux-amd64.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm
```

Now we have `helm` and `kubectl` and can access our `kind` cluster:

```
kubectl get nodes
NAME                  STATUS   ROLES                  AGE   VERSION
vault-control-plane   Ready    control-plane,master   37s   v1.21.1
```

Let's add the Helm repositories, so we can access the Kubernetes manifests

```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

## Storage: Consul

We will use a very basic Consul cluster for our Vault backend. </br>
Let's find what versions of Consul are available:

```
helm search repo hashicorp/consul --versions
```

We can use chart `0.39.0` which is the latest at the time of this demo
Let's create a manifests folder and grab the YAML:

```

mkdir manifests

helm template consul hashicorp/consul \
  --namespace vault \
  --version 0.39.0 \
  -f consul-values.yaml \
  > ./manifests/consul.yaml
```

Deploy the consul services:

```
kubectl create ns vault
kubectl -n vault apply -f ./manifests/consul.yaml
kubectl -n vault get pods
```


## TLS End to End Encryption

See steps in [./tls/ssl_generate_self_signed.md](./tls/ssl_generate_self_signed.md)
You'll need to generate TLS certs (or bring your own)
Remember not to check-in your TLS to GIT :)

Create the TLS secret 

```
kubectl -n vault create secret tls tls-ca \
 --cert ./tls/ca.pem  \
 --key ./tls/ca-key.pem

kubectl -n vault create secret tls tls-server \
  --cert ./tls/vault.pem \
  --key ./tls/vault-key.pem
```

## Generate Kubernetes Manifests


Let's find what versions of vault are available:

```
helm search repo hashicorp/vault --versions
```

In this demo I will use the `0.19.0` chart </br>

Let's firstly create a `values` file to customize vault.
Let's grab the manifests:

```
helm template vault hashicorp/vault \
  --namespace vault \
  --version 0.19.0 \
  -f vault-values.yaml \
  > ./manifests/vault.yaml
```

## Deployment

```
kubectl -n vault apply -f ./manifests/vault.yaml
kubectl -n vault get pods
```

## Initialising Vault

```
kubectl -n vault exec -it vault-0 -- sh
kubectl -n vault exec -it vault-1 -- sh
kubectl -n vault exec -it vault-2 -- sh

vault operator init
vault operator unseal

kubectl -n vault exec -it vault-0 -- vault status
kubectl -n vault exec -it vault-1 -- vault status
kubectl -n vault exec -it vault-2 -- vault status

```
## Web UI

Let's checkout the web UI:

```
kubectl -n vault get svc
kubectl -n vault port-forward svc/vault-ui 443:8200
```
Now we can access the web UI [here](https://localhost/)

## Enable Kubernetes Authentication

For the injector to be authorised to access vault, we need to enable K8s auth

```
kubectl -n vault exec -it vault-0 -- sh 

vault login
vault auth enable kubernetes

vault write auth/kubernetes/config \
token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
kubernetes_host=https://${KUBERNETES_PORT_443_TCP_ADDR}:443 \
kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt \
issuer="https://kubernetes.default.svc.cluster.local"
exit
```

# Summary

So we have a vault, an injector, TLS end to end, stateful storage.
The injector can now inject secrets for pods from the vault.

Now we are ready to use the platform for different types of secrets:

## Secret Injection Guides

### Basic Secrets

Objective:
---------- 
* Let's create a basic secret in vault manually
* Application consumes the secret automatically

[Try it](./example-apps/basic-secret/readme.md)




