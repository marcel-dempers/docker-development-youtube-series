# Hashicorp Vault Guide

# Vault

For this tutorial, I use Kuberentes 1.17
It's critical because we'll need certain [admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) enabled.

To get 1.17 for Linux\Windows, just use `kind` since you can create a 1.17 with admissions all setup.

```
kind create cluster --name vault --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62
```

## TLS End to End Encryption

See steps in `hashicorp/vault/tls/ssl_generate_self_signed.txt`
You'll need to generate TLS certs (or bring your own)
Create base64 strings from the files, place it in the `server-tls-secret.yaml` and apply it.
Remember not to check-in your TLS to GIT :)

## Deployment

```
kubectl create ns vault-example
kubectl -n vault-example apply -f .\hashicorp\vault\server\
```

## Storage

```
kubectl -n vault-example get pvc
```
ensure vault-claim is bound, if not, `kubectl -n vault-example describe pvc vault-claim`
ensure correct storage class is used for your cluster.
if you need to change the storage class, deleve the pvc , edit YAML and re-apply

## Initialising Vault

```
kubectl -n vault-example exec -it vault-example-0 vault operator init
kubectl -n vault-example exec -it vault-example-0 vault operator unseal
```

## Depploy the Injector

Injector allows pods to automatically get secrets from the vault.

```
kubectl -n vault-example apply -f .\hashicorp\vault\injector\
```







