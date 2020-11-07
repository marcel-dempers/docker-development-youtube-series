# Introduction to Fluentd on Kubernetes

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name fluentd --image kindest/node:v1.19.1
```

## Fluentd Manifests

I would highly recommend to use manifests from the official fluentd [github repo](https://github.com/fluent/fluentd-kubernetes-daemonset) <br/>

The manifests found here are purely for demo purpose. <br/>

In this example I will use the most common use case and we'll break it down to get an understanding of each component.