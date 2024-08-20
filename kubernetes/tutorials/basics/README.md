# Kubernetes Tutorial: The Basics 

<a href="https://youtu.be/JeAHlTYB1Qk" title="How to learn Kubernetes in 2022"><img src="https://i.ytimg.com/vi/JeAHlTYB1Qk/hqdefault.jpg" width="40%" alt="How to learn Kubernetes in 2022" /></a>

This guide is aimed to fast-track your Kubernetes learning by focusing on a practical hands-on overview guide. </br>

When learning Kubernetes, you usually have an idea of some existing system you own and manage, or a website that you are building. </br>

The challenge is understanding which Kubernetes building blocks you need in order to run your workloads on Kubernetes </br>

<hr/>
<b>The problem:</b> "I want to adopt Kubernetes" </br>
<b>The problem:</b> "I have some common existing infrastructure"
<hr/>

<b>Our focus:</b> Solving the problem by learning each building block
in order to port our infrastructure to Kubernetes. 

## Understanding Containers

Before even looking at Kubernetes, you need to have a general understanding of containers like `docker`.
Your workloads need to fit in containers in order to be shipped on Kubernetes. </br>
Containers also have a bunch of assumptions that you need to meet.

* Defining the container - `Dockerfile`
* Serving traffic - Exposing ports
* Configuration - mount config files & secrets or `env` variables
* Data persistence - When a container is terminated, everything inside the container is gone
* Container entrypoint - The main process that runs in the container. Your app


### Docker installation 

* Install Docker [here](https://docs.docker.com/get-docker/)
* Let's take a look at [Wordpress on Docker Hub](https://hub.docker.com/_/wordpress)
* Build our docker file

## Create Network

```
docker network create wordpress
```

## Build & Test container images

 ### Wordpress example
```
cd kubernetes\tutorials\basics\

docker build -f dockerfiles/wordpress.dockerfile . -t aimvector/wordpress-example

```

* Run our Wordpress container

```
docker run -it --rm -p 80:80 --net wordpress aimvector/wordpress-example
```

The wordpress container will be visible on port 80 on `http://localhost/`

### MySQL example

* We need a database, let's take a look at [MySQL on Docker Hub](https://hub.docker.com/_/mysql)
* Build our MySQl container image

```
docker build -f dockerfiles/mysql.dockerfile . -t aimvector/mysql-example
```

* How do we run our MySQL ? 

We need to understand that databases require storage and state
Just like installing software on a server, it will store its files in some 
directory. Mysql stores its files under `/var/lib/mysql`

* We need a volume mount

Let's see how to run this in docker

```
mkdir data

docker run --rm -d `
--name mysql `
--net wordpress `
-e MYSQL_DATABASE=exampledb `
-e MYSQL_USER=exampleuser `
-e MYSQL_PASSWORD=examplepassword `
-e MYSQL_RANDOM_ROOT_PASSWORD=1 `
-v ${PWD}/data:/var/lib/mysql `
aimvector/mysql-example

# we can see the container with
docker ps
CONTAINER ID   IMAGE                     COMMAND                  CREATED         STATUS         PORTS                 NAMES
92cde663a3f5   aimvector/mysql-example   "docker-entrypoint.s…"   5 seconds ago   Up 3 seconds   3306/tcp, 33060/tcp   mysql

```

* Run Wordpress and connect it to MySQL

```
docker run -d `
--rm `
-p 80:80 `
--name wordpress `
--net wordpress `
-e WORDPRESS_DB_HOST=mysql `
-e WORDPRESS_DB_USER=exampleuser `
-e WORDPRESS_DB_PASSWORD=examplepassword `
-e WORDPRESS_DB_NAME=exampledb `
aimvector/wordpress-example
```

### Clean up 

```
docker rm -f wordpress
docker rm -f mysql
docker network rm wordpress
rm data
```

## Kubernetes Tools: kubectl

To manage and work with Kubernetes, you need `kubectl` </br>
Let's grab that from [here](https://kubernetes.io/docs/tasks/tools/)


## Run Kubernetes Locally

* Install `kubectl` to work with kubernetes

We'll head over to the [kubernetes](https://kubernetes.io/docs/tasks/tools/) site to download `kubectl` 

* Install the `kind` binary

You will want to head over to the [kind](https://kind.sigs.k8s.io/) site

* Create a cluster 

```
kind create cluster --image kindest/node:v1.23.5
```

## Namespaces 

```
kubectl create namespace cms
```

## Configmaps

[Environment Variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for pods

[How to use](https://kubernetes.io/docs/concepts/configuration/configmap/) configmaps

```
kubectl -n cms create configmap mysql `
--from-literal MYSQL_RANDOM_ROOT_PASSWORD=1

kubectl -n cms get configmaps
```

## Secrets

[How to use](https://kubernetes.io/docs/concepts/configuration/secret/) secrets in pods

```
kubectl -n cms create secret generic wordpress `
--from-literal WORDPRESS_DB_HOST=mysql `
--from-literal WORDPRESS_DB_USER=exampleuser `
--from-literal WORDPRESS_DB_PASSWORD=examplepassword `
--from-literal WORDPRESS_DB_NAME=exampledb

kubectl -n cms create secret generic mysql `
--from-literal MYSQL_USER=exampleuser `
--from-literal MYSQL_PASSWORD=examplepassword `
--from-literal MYSQL_DATABASE=exampledb


kubectl -n cms get secret

```

## Deployments

* Deployment [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

cd kubernetes\tutorials\basics

```
kubectl -n cms apply -f yaml/deploy.yaml
kubectl -n cms get pods
```

# Services

Services [documentation](https://kubernetes.io/docs/concepts/services-networking/service/)

```
kubectl -n cms apply -f .\yaml\service.yaml
kubectl -n cms get svc
```

# Storage Class

StorageClass [documentation](https://kubernetes.io/docs/concepts/storage/storage-classes/)

```
kubectl get storageclass
```

# Statefulset

Statefulset [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)

Let's deploy our `mysql` using what we learnt above:

```
kubectl -n cms apply -f yaml/statefulset.yaml

kubectl -n cms get pods
```

## Persistent Volumes

[Documentation](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)

## Port Forwarding

We can access private service endpoints or pods using `port-forward` :

```
kubectl -n cms get pods
kubectl -n cms port-forward <pod-name> 80
```

## Public Traffic

In order to make our site public, its common practise to expose web servers via  </br>
a proxy or API gateway. </br>
In Kubernetes, an Ingress is used.

## Ingress

To use an ingress, we need an ingress controller

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.3/deploy/static/provider/cloud/deploy.yaml

kubectl -n ingress-nginx get pods

kubectl -n ingress-nginx --address 0.0.0.0 port-forward svc/ingress-nginx-controller 80
```

Create an Ingress

```
kubectl -n cms apply -f yaml/ingress.yaml
```
