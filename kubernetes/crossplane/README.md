# Introduction to Crossplane

[Crossplane](https://www.crossplane.io/) </br>
[Crossplane Documentation](https://docs.crossplane.io/v1.19/) </br>

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name crossplane --image kindest/node:v1.33.0
```

## Installing Crossplane

In this guide we will reference the official document steps in the links above. </br>
I've recorded the commands we follow in the video too 


```
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm search repo crossplane-stable --versions
```

We'll install version `1.19.1` at the time of this guide 

```
VERSION=1.19.1

helm install crossplane \
crossplane-stable/crossplane \
--namespace crossplane-system \
--version $VERSION \
--create-namespace
```

View our install: 

```
kubectl get pods -n crossplane-system
kubectl get deployments -n crossplane-system
```

## Providers 

[Providers](https://docs.crossplane.io/latest/concepts/providers/)