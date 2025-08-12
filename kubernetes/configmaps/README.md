# Introduction to Kubernetes: Configmaps

<a href="https://youtu.be/o-gXx7r7Rz4" title="k8s-cm"><img src="https://i.ytimg.com/vi/o-gXx7r7Rz4/hqdefault.jpg" width="20%" alt="k8s-cm" /></a> 

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

Deploy a `ConfigMap` 

```
kubectl apply -f kubernetes/configmaps/configmap.yaml
```

Deploy a pod that consumes the `ConfigMap` 

```
kubectl apply -f kubernetes/configmaps/deployment.yaml
```

Checking the deployment pods
The `kubectl describe` command is essential for checking out kubernetes resources

```
kubectl describe deploy example-deploy
```

We can checkout the config contents inside the pod's container with the `kubectl exec` command to get shell

```
kubectl exec -it <pod-name> -- bash
```

When we make a change to the `ConfigMap`, we will notice that the configuration file changes and reflects inside the pod. </br>
This takes a number of seconds to a minute or so to propogate