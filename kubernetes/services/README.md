# Introduction to Kubernetes: Services

<a href="https://youtu.be/hiiiGLnPdq4" title="k8s-services"><img src="https://i.ytimg.com/vi/hiiiGLnPdq4/hqdefault.jpg" width="20%" alt="k8s-services" /></a> 

## Create a Kubernetes cluster 

Firstly, we will need a Kubernetes cluster and will create one using [kind](https://kind.sigs.k8s.io/)

```
kind create cluster --name services --image kindest/node:v1.31.1
```

Test the cluster 

```
kubectl get nodes
NAME                     STATUS   ROLES           AGE     VERSION
services-control-plane   Ready    control-plane   5m48s   v1.31.1
```

## Deploy a few pods

Let's deploy a few pods using a Kubernetes Deployment. To understand services, we need to deploy some pods that become our "upstream" or "endpoint" that we want to access. </br>

```
kubectl apply -f kubernetes/deployments/deployment.yaml 
```

## Deploy a service

We can now deploy a Kubernetes service that targets our deployment:


```
kubectl apply -f kubernetes/services/service.yaml
```