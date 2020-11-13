# Introduction to cert-manager for Kubernetes

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name certmanager --image kindest/node:v1.19.1
```


## Issuer 

https://cert-manager.io/docs/concepts/issuer/


## Certificate

https://cert-manager.io/docs/concepts/certificate/
