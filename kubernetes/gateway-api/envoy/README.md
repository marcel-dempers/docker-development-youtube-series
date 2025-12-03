# Introduction to Envoy: Gateway API

## Prerequisites

To get started, you will need to follow the the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRD
* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## What is Envoy

It's important to take a step back and fully understand and appreciate what Envoy is. </br>
Envoy is not only a Gateway API, that's just one of it's features. </br>
Let's take a look at the [Official Site](https://www.envoyproxy.io/) and jump to the documentation. </br>
Envoy has a separate web page for the Gateway API feature. </br>

## Envoy: Gateway API controller

Envoy has a `helm` chart specifically for the gateway product. </br>
The `helm` chart relies on the Gateway API CRD's we installed in our cluster already as well as Envoy Proxy. </br>

### Installation 

We can install the helm chart with `helm install` or upgrade it with `helm upgrade`:

```shell

CHART_VERSION="v1.6.0"
helm show chart oci://docker.io/envoyproxy/gateway-helm
helm show values oci://docker.io/envoyproxy/gateway-helm > kubernetes/gateway-api/envoy/default-values.yaml

helm install envoy-gateway oci://docker.io/envoyproxy/gateway-helm \
  --version $CHART_VERSION \
  --namespace envoy-gateway-system \
  --create-namespace
```

Upgrade:

```shell
helm upgrade envoy-gateway oci://docker.io/envoyproxy/gateway-helm \
  --version $CHART_VERSION \
  --namespace envoy-gateway-system \
```

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [default-values.yaml](./default-values.yaml)

### Check Installation

Now we should have the Envoy gateway API controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implements some gateway API CRD's. </br>

```shell
# check the controller pods
kubectl -n envoy-gateway-system get pods

# check the controller pod logs 
kubectl -n envoy-gateway-system logs -l app.kubernetes.io/instance=envoy-gateway

```

## Install an Envoy Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/envoy/01-gatewayclass.yaml
```

## Install an Envoy Gateway

```shell
kubectl apply -f kubernetes/gateway-api/envoy/02-gateway.yaml
```

When we apply the gateway, we get a new gateway api pod. 

```shell
# check the new gateway-api pod
kubectl -n envoy-gateway-system get pods

# we also have a new service
kubectl -n envoy-gateway-system get svc
```

### Gateway Configuration

`EnvoyProxy` is the CRD that allows us to configure each gateway API instance. </br>
For example, we can change the `deployment` or `service` of the instance like so:

```shell
kubectl apply -f kubernetes/gateway-api/envoy/02.1-gateway-config.yaml

# we should see 2 replicas 
kubectl -n envoy-gateway-system get deploy
kubectl -n envoy-gateway-system get pods

# port forward for access (get the correct service)
kubectl -n envoy-gateway-system get svc
kubectl -n envoy-gateway-system port-forward svc/envoy-default-gateway-api-30a1473e 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>


## Security Policy

### CORS 
### Basic Auth

### API Auth 

### External Auth


kubectl -n envoy-gateway-system logs -l app.kubernetes.io/managed-by=envoy-gateway