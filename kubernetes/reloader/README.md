# Introduction to K8s Reloader 

In this demo we will walk through a project called [reloader](https://github.com/stakater/Reloader)

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

Deploy a `Secret`

```
kubectl apply -f kubernetes/secrets/secret.yaml
```

## Install Reloader

```
helm repo add stakater https://stakater.github.io/stakater-charts
helm repo update

# see versions
helm search repo stakater/reloader -l

#pick a version
CHART_VERSION="2.1.5"


helm install reloader stakater/reloader \
--create-namespace \
--namespace reloader \
--version $CHART_VERSION \
--values kubernetes/reloader/values.yaml
```

### Usage

Deploy an application that consumes the 'ConfigMap' and 'Secret' with auto reload enabled:

```
kubectl apply -f kubernetes/reloader/deployment.yaml
```

We can now update the `ConfigMap` object and observe our `example-app` pods restarting when the `ConfigMap` changes.
We can also run `kubectl logs -l app=example-app` in a separate terminal and observe the changing values inside pods. 
