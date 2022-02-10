# Learn Kubernetes <img src="https://www.shareicon.net/data/128x128/2017/04/11/883708_media_512x512.png" alt="YouTube" width="5%" height="5%"> :hammer::wrench:

This guide is aimed to fast-track your Kubernetes learning by focusing on a practical hands-on overview guide. </br>
This means, not too much of a deepdive, but enough to get a feel for the required building blocks of Kubernetes so you can align it with a real world problem that you are trying to solve. </br>

<b>The problem:</b> "I want to adopt Kubernetes" </br>
<b>The problem:</b> "I have some common existing infrastructure"

<b>Our focus:</b> Solving the problem by learning each building block
in order to port our infrastructure to Kubernetes. 

## Docker installation 

* Install Docker [here](https://docs.docker.com/get-docker/)

## Run Kubernetes

* Install `kubectl` to work with kubernetes

We'll head over to the [kubernetes](https://kubernetes.io/docs/tasks/tools/) site to download `kubectl` 

* Install the `kind` binary

You will want to head over to the [kind](https://kind.sigs.k8s.io/) site

* Create a cluster 

```
kind create cluster
```

## Namespaces 

```
kubectl create namespace cms
```

## Deployments

* Deployment [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

cd kubernetes\tutorial 

```
kubectl -n cms apply -f deploy.yaml
kubectl -n cms get pods

kubectl -n cms port-forward <pod-name> 80
```

[Environment Variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for pods

## Secrets

```
kubectl -n cms create secret generic wordpress `
--from-literal WORDPRESS_DB_HOST=mysql `
--from-literal WORDPRESS_DB_USER=exampleuser `
--from-literal WORDPRESS_DB_PASSWORD=examplepassword `
--from-literal WORDPRESS_DB_NAME=exampledb

kubectl -n cms get secret

```
[How to use](https://kubernetes.io/docs/concepts/configuration/secret/) secrets in pods

Apply changes to our deployment

```
kubectl -n cms apply -f deploy.yaml
```

We can `port-forward` again, and notice an error connecting to the database because the database does not exist

# Statefulset

Statefulset [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)

# Storage Class

StorageClass [documentation](https://kubernetes.io/docs/concepts/storage/storage-classes/)

# Services

Services [documentation](https://kubernetes.io/docs/concepts/services-networking/service/)

Let's deploy our `mysql` using what we learnt above:

```
kubectl -n cms apply -f .\statefulset.yaml
```

## Full playlist

<a href="https://www.youtube.com/playlist?list=PLHq1uqvAteVvUEdqaBeMK2awVThNujwMd" title="Kubernetes"><img src="https://i.ytimg.com/vi/8h4FoWK7tIA/hqdefault.jpg" width="50%" height="50%" alt="Kubernetes Guide" /></a>

## Getting Started with Kubernetes on Windows for beginners

<a href="https://www.youtube.com/watch?v=8h4FoWK7tIA" title="Kubernetes"><img src="https://i.ytimg.com/vi/8h4FoWK7tIA/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Kubectl

Checkout [kubectl](./kubectl.md) for detailed steps

<a href="https://www.youtube.com/watch?v=feLpGydQVio" title="Kubernetes"><img src="https://i.ytimg.com/vi/feLpGydQVio/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Deployments

Checkout [deployments](./deployments/readme.md) for detailed steps

<a href="https://www.youtube.com/watch?v=DMpEZEakYVc" title="Kubernetes"><img src="https://i.ytimg.com/vi/DMpEZEakYVc/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Configuration Management

<a href="https://www.youtube.com/watch?v=o-gXx7r7Rz4" title="Kubernetes"><img src="https://i.ytimg.com/vi/o-gXx7r7Rz4/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Secret Management

<a href="https://www.youtube.com/watch?v=o36yTfGDmZ0" title="Kubernetes"><img src="https://i.ytimg.com/vi/o36yTfGDmZ0/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Services

<a href="https://www.youtube.com/watch?v=xhva6DeKqVU" title="Kubernetes"><img src="https://i.ytimg.com/vi/xhva6DeKqVU/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Ingress

<a href="https://www.youtube.com/watch?v=u948CURLDJA" title="Kubernetes"><img src="https://i.ytimg.com/vi/u948CURLDJA/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>
<a href="https://www.youtube.com/watch?v=izWCkcJAzBw" title="Kubernetes"><img src="https://i.ytimg.com/vi/izWCkcJAzBw/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## StatefulSets

<a href="https://www.youtube.com/watch?v=zj6r_EEhv6s" title="Kubernetes"><img src="https://i.ytimg.com/vi/zj6r_EEhv6s/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Persistent Volumes

<a href="https://www.youtube.com/watch?v=ZxC6FwEc9WQ" title="Kubernetes"><img src="https://i.ytimg.com/vi/ZxC6FwEc9WQ/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>

## Manage YAML

<a href="https://www.youtube.com/watch?v=5gsHYdiD6v8" title="Kubernetes"><img src="https://i.ytimg.com/vi/5gsHYdiD6v8/hqdefault.jpg" width="25%" height="25%" alt="Kubernetes Guide" /></a>


