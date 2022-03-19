# Tutorial: The Basics 

This guide is aimed to fast-track your Kubernetes learning by focusing on a practical hands-on overview guide. </br>

<hr/>
<b>The problem:</b> "I want to adopt Kubernetes" </br>
<b>The problem:</b> "I have some common existing infrastructure"
<hr/>

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
