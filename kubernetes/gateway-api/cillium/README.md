# Introduction to Cillium: Gateway API

<a href="https://youtu.be/xxxxx" title="xxxxx"><img src="https://i.ytimg.com/vi/xxxxx/hqdefault.jpg" width="40%" alt="xxxxx" /></a>

## Prerequisites

This guide will assume you have a basic understanding of the Gateway API. If it's your first time using Gateway API, please see the [Introduction to Gateway API guide](../README.md) first </br>

This guide has a slight twist from the other Gateway APIs. </br>
It's because Cillium is a network CNI for Kubernetes.</br>
This means we need to disable the default kubernetes CNI networking in our cluster, so we will be creating it slightly differently. </br>

We also need to install the Gateway API CRDs followed by Cillium, before we deploy all our example applications. </br>

## What is Cillium

#TODO

### Create a kubernetes cluster

```shell
kind create cluster --name gatewayapi --image kindest/node:v1.34.0 --config kubernetes/gateway-api/cillium/kind.yaml
```

Test our cluster and makes sure `kubectl` is configured for it:

```shell
kubectl get nodes
NAME                       STATUS     ROLES           AGE   VERSION
gatewayapi-control-plane   NotReady   control-plane   93s   v1.34.0
gatewayapi-worker          NotReady   <none>          83s   v1.34.0
```

The nodes will remain in a `NotReady` state until we unstall our network plugin, which we will do with Cillium. </br>

### Install Gateway API CRDs

Now we can deploy the Gateway API CRDs following our [Introduction to Gateway API guide](../README.md#gateway-api-crds). </br>
The CRDs need to be in place when we install Cillium. </br>

### Installation 

Install:

```shell
CHART_VERSION="1.18.5"
helm repo add cilium https://helm.cilium.io
helm repo update
helm search repo cilium --versions

helm upgrade --install cilium cilium/cilium \
  --version $CHART_VERSION \
  --namespace kube-system \
  --set kubeProxyReplacement=true \
  --set gatewayAPI.enabled=true
```

Check our installation

```shell
# check the pods
kubectl -n kube-system get pods

# nodes should be Ready once above is Running
kubectl get nodes
NAME                       STATUS   ROLES           AGE   VERSION
gatewayapi-control-plane   Ready    control-plane   38m   v1.34.0
gatewayapi-worker          Ready    <none>          37m   v1.34.0
```


<b>In the introduction guide, you will:</b>

* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## Install a Cillium Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/cillium/01-gatewayclass.yaml
```

## Install a Cillium Gateway

```shell
kubectl apply -f kubernetes/gateway-api/cillium/02-gateway.yaml
```

Let's start with the [Official Documentation](https://cilium.io)

## Cillium: Gateway API controller

#TODO

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>

#TODO values.yaml

* Access logs - fields and format
* Ports to use for incoming traffic
* Infrastructure annotations \ labels
* Control Gateway API
  * Default CRDs
  * Default GatewayClass
  * Default Gateways

It's always good to get a grip on the [default helm values]()


## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>


