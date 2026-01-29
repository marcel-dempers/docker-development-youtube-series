# Introduction to Cilium: Gateway API

#TODO
<a href="https://youtu.be/xxxxx" title="xxxxx"><img src="https://i.ytimg.com/vi/xxxxx/hqdefault.jpg" width="40%" alt="xxxxx" /></a>

## Prerequisites

This guide will assume you have a basic understanding of the Gateway API. If it's your first time using Gateway API, please see the [Introduction to Gateway API guide](../README.md) first </br>

This guide has a slight twist from the other Gateway APIs. </br>
It's because Cilium is a network CNI for Kubernetes.</br>
This means we need to disable the default kubernetes CNI networking in our cluster, so we will be creating it slightly differently. </br>

We also need to install the Gateway API CRDs followed by Cilium, before we deploy all our example applications. </br>

## What is Cilium

[Cilium](https://cilium.io/) leverages eBPF technology to provide high-performance networking, security and observability for containerised applications, especially in Kubernetes. </br>

eBPF (Extended Berkeley Packet Filter) is a technology that allows us to run custom programs safely in the Linux Kernel without the need for modules, changing source code. </br>

These programs allow programs like Cilium to add security, performance and observability to the network. </br>

### Create a kubernetes cluster

```shell
kind create cluster --name gatewayapi --image kindest/node:v1.34.0 --config kubernetes/gateway-api/cilium/kind.yaml
```

Test our cluster and makes sure `kubectl` is configured for it:

```shell
kubectl get nodes
NAME                       STATUS     ROLES           AGE   VERSION
gatewayapi-control-plane   NotReady   control-plane   93s   v1.34.0
gatewayapi-worker          NotReady   <none>          83s   v1.34.0
```

The nodes will remain in a `NotReady` state until we unstall our network plugin, which we will do with Cilium. </br>

### Install Gateway API CRDs

Now we can deploy the Gateway API CRDs following our [Introduction to Gateway API guide](../README.md#gateway-api-crds). </br>
The CRDs need to be in place when we install Cilium. </br>

### Configuration

It's always good to get a grip on the [default helm values](./default-values.yaml) to see what we can do with the chart. </br>

In this demo we'll only focus on the `gatewayAPI` section of the `helm` values file. </br>
We created our own [values.yaml](./values.yaml) which allows us to enable Gateway API, and use Cilium as the kube proxy replacement so it takes over our networking. </br>

### Installation 

Install:

```shell
CHART_VERSION="1.18.6"
helm repo add cilium https://helm.cilium.io
helm repo update
helm search repo cilium --versions

# default configuration
helm show values cilium/cilium > kubernetes/gateway-api/cilium/default-values.yaml

helm upgrade --install cilium cilium/cilium \
  --version $CHART_VERSION \
  --namespace kube-system \
  --values kubernetes/gateway-api/cilium/values.yaml
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

We can also check the Cilium Operator logs to ensure we are smooth sailing:
```shell
kubectl -n kube-system logs -l app.kubernetes.io/name=cilium-operator
```

<b>In the [introduction guide](../README.md), you will:</b>

* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## Install a Cilium Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/cilium/01-gatewayclass.yaml
```

The `CiliumGatewayClassConfig` CRD for Cilium allows us to customize the Gateway class. </br>

This allows us to change the deployment and services for gateways that use this class. </br> In our case, we can set the `service.type = NodePort`


```shell
kubectl apply -f kubernetes/gateway-api/cilium/01.1-gatewayclass-config.yaml
```

## Install a Cilium Gateway

```shell
kubectl apply -f kubernetes/gateway-api/cilium/02-gateway.yaml
```

## Install the Cilium CLI

```
CILIUM_CLI_VERSION=v0.19.0
CLI_ARCH=amd64
curl -L --fail --remote-name-all https://github.com/cilium/cilium-cli/releases/download/${CILIUM_CLI_VERSION}/cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check cilium-linux-${CLI_ARCH}.tar.gz.sha256sum
tar xzvf cilium-linux-${CLI_ARCH}.tar.gz -C /usr/local/bin
rm cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>


