# Hashicorp Vault Guide

# Vault

For this tutorial, I use Kuberentes 1.17
It's critical because we'll need certain [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) enabled.

To get 1.17 for Linux\Windows, just use `kind` since you can create a 1.17 with admissions all setup.

```
#Windows
kind create cluster --name vault --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62

#Linux
kind create cluster --name vault --kubeconfig ~/.kube/kind-vault --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62
```

## TLS End to End Encryption

VIDEO: ```<Coming-Soon>```  
See steps in [./tls/ssl_generate_self_signed.txt](./tls/ssl_generate_self_signed.txt)
You'll need to generate TLS certs (or bring your own)
Create base64 strings from the files, place it in the `server-tls-secret.yaml` and apply it.
Remember not to check-in your TLS to GIT :)

## Deployment

```
kubectl create ns vault-example
kubectl -n vault-example apply -f ./hashicorp/vault/server/
kubectl -n vault-example get pods
```

## Storage

```
kubectl -n vault-example get pvc
```
ensure vault-claim is bound, if not, `kubectl -n vault-example describe pvc vault-claim`
ensure correct storage class is used for your cluster.
if you need to change the storage class, delete the pvc, edit YAML and re-apply

## Initialising Vault

```
kubectl -n vault-example exec -it vault-example-0 vault operator init
# unseal 3 times
kubectl -n vault-example exec -it vault-example-0 vault operator unseal
kubectl -n vault-example get pods
```

## Deploy the Injector

VIDEO: ```<Coming-Soon>```  
Injector allows pods to automatically get secrets from the vault.

```
kubectl -n vault-example apply -f ./hashicorp/vault/injector/
kubectl -n vault-example get pods
```

## Injector Kubernetes Auth Policy

For the injector to be authorised to access vault, we need to enable K8s auth

```
kubectl -n vault-example exec -it vault-example-0 vault login
kubectl -n vault-example exec -it vault-example-0 vault auth enable kubernetes

kubectl -n vault-example exec -it vault-example-0 sh
vault write auth/kubernetes/config \
token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
kubernetes_host=https://${KUBERNETES_PORT_443_TCP_ADDR}:443 \
kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
exit

kubectl -n vault-example get pods
```

# Summary

So we have a vault, an injector, TLS end to end, stateful storage.
The injector can now inject secrets for pods from the vault.

Now we are ready to use the platform for different types of secrets:

## Secret Injection Guides

I've broken this down into basic guides to avoid this document from becoming too large.

### Basic Secrets

Objective:
---------- 
* Let's create a basic secret in vault manually
* Application consumes the secret automatically

[Try it](./example-apps/basic-secret/readme.md)

### Dynamic Secrets: Postgres

Objective:
---------- 
* We have a Postgres Database
* Let's delegate Vault to manage life cycles of our database credentials
* Deploy an app, that automatically gets it's credentials from vault

[Try it](./example-apps/dynamic-postgresql/readme.md)




