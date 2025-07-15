# Introduction to K8s Reloader 

## Create a cluster with Kind

```
kind create cluster --name demo --image kindest/node:v1.33.1
```

Checkout our cluster 

```
kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
demo-control-plane   Ready    control-plane   19m   v1.33.1
```

For this demo, we'll want to deploy a `ConfigMap` and/or a `Secret` and consume these using a `Deployment`

Deploy a `ConfigMap`

```
kubectl apply -f kubernetes/configmaps/configmap.yaml
```

