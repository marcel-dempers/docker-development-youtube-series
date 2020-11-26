# Introduction to Fluentd on Kubernetes

## Prerequisites 

You will need a basic understanding of Fluentd before you attempt to run it on Kubernetes.<br/> 
Fluentd and Kubernetes have a bunch of moving parts.<br/> 
To understand the basics of Fluentd, I highly recommend you start with this video: <br/> 

<a href="https://youtu.be/Gp0-7oVOtPw" title="Fluentd"><img src="https://i.ytimg.com/vi/Gp0-7oVOtPw/hqdefault.jpg" width="50%" height="50%" alt="Fluentd" /></a>

The most important components to understand is the fluentd `tail` plugin. <br/>
This plugin is used to read logs from containers and pods on the file system and collect them.

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name fluentd --image kindest/node:v1.19.1
```

## Fluentd Manifests

I would highly recommend to use manifests from the official fluentd [github repo](https://github.com/fluent/fluentd-kubernetes-daemonset) for production usage <br/>

The manifests found here are purely for demo purpose. <br/>
The manifests in this repo are broken down and simplified for educational purpose. </br>
<br/>
In this example I will use the most common use case and we'll break it down to get an understanding of each component.

## Fluentd Docker

I would recommend to start with the official [fluentd](https://hub.docker.com/r/fluent/fluentd/)
docker image. <br/>
You may want to build your own image if you want to install plugins.
In this demo I will be using the `fluentd` elasticsearch plugin <br/>
It's pretty simple to adjust `fluentd` to send logs to any other destination in case you are not an `elasticsearch` user. <br/>

<br/>

Let's build our [docker image](https://github.com/marcel-dempers/docker-development-youtube-series/blob/master/monitoring/logging/fluentd/introduction/dockerfile) in the introduction folder:


```
cd .\monitoring\logging\fluentd\kubernetes\

#note: use your own tag!
docker build . -t aimvector/fluentd-demo

#note: use your own tag!
docker push aimvector/fluentd-demo

```

## Fluentd Namespace

I like to run certain infrastructure components in their own namespaces. <br/>
If you are using the official manifests, they may be using the `kube-system` namespace instead. <br/>
You may want to carefully adjust it based on your preference <br/>
Let's create a `fluentd` namespace: <br/>

```
kubectl create ns fluentd

```
## Fluentd Configmap

In my [fluentd introduction video](https://youtu.be/Gp0-7oVOtPw), I talk about how `fluentd` allows us to simplify our configs using the `include` statement. <br/> 
This helps us prevent having a large complex file.

<br/>

We have 5 files in our `fluentd-configmap.yaml` :
* fluent.conf: Our main config which includes all other configurations
* pods-kind-fluent.conf: `tail` config that sources all pod logs on the `kind` cluster.
  Note: `kind` cluster writes its log in a different format
* pods-fluent.conf: `tail` config that sources all pod logs on the `kubernetes` host in the cloud. <br/>
  Note: When running K8s in the cloud, logs may go into JSON format.
* file-fluent.conf: `match` config to capture all logs and write it to file for testing log collection </br>
  Note: This is great to test if collection of logs works
* elastic-fluent.conf: `match` config that captures all logs and sends it to `elasticseach`

Let's deploy our `configmap`:

```
kubectl apply -f .\monitoring\logging\fluentd\kubernetes\fluentd-configmap.yaml

```

## Fluentd Daemonset

Let's deploy our `daemonset`:

```
kubectl apply -f .\monitoring\logging\fluentd\kubernetes\fluentd-rbac.yaml 
kubectl apply -f .\monitoring\logging\fluentd\kubernetes\fluentd.yaml
kubectl -n fluentd get pods

```

Let's deploy our example app that writes logs to `stdout`

```
kubectl apply -f .\monitoring\logging\fluentd\kubernetes\counter.yaml
kubectl get pods

```

## Demo ElasticSearch and Kibana

```
kubectl create ns elastic-kibana

# deploy elastic search
kubectl -n elastic-kibana apply -f .\monitoring\logging\fluentd\kubernetes\elastic\elastic-demo.yaml
kubectl -n elastic-kibana get pods

# deploy kibana
kubectl -n elastic-kibana apply -f .\monitoring\logging\fluentd\kubernetes\elastic\kibana-demo.yaml
kubectl -n elastic-kibana get pods
```

## Kibana

```
kubectl -n elastic-kibana port-forward svc/kibana 5601
```