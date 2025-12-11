# Introduction to Istio: Gateway API

## Prerequisites

To get started, you will need to follow the the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs
* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## What is Istio

Before diving into installing Istio, its very important to understand what Istio is. </br>
<b>Istio is not only a Gateway API</b>, its a service mesh. </br>
In a distributed system with service-to-service communication, a service mesh places a proxy next to each service. </br>
This proxy intercepts all inbound and outbound communication for it's neighboring service. </br>
Doing so, a service mesh provides many features
* <b>observability</b>:  monitor network calls (metrics \ traces)
* <b>traffic</b> management: traffic routing features
* <b>resillience</b>: automatic retry, circuit breaker, rate limiting etc
* <b>security</b>: mutual TLS, authentication, policies etc

Let's take a look at the [Official Site](https://istio.io/latest/docs/overview/what-is-istio/) and jump to the documentation. </br>

## Istio: Gateway API controller

With Istio, we have a few more steps to install the gateway API. We fistly need to install the base `helm` chart which includes Istio specific CRD components required. </br>

### Installation 

We can install the helm chart with `helm install` or upgrade it with `helm upgrade`:

```shell

BASE_CHART_VERSION="1.28.1"
ISTIOD_CHART_VERSION="1.28.1"

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

# get versions for the charts
helm search repo istio/base --versions | head
helm search repo istio/istiod --versions | head

helm install istio-base istio/base \
  --version $BASE_CHART_VERSION \
  --namespace istio-system \
  --create-namespace
  --set defaultRevision=default 

helm install istiod istio/istiod \
  --version $ISTIOD_CHART_VERSION \
  --namespace istio-system \
  --wait

```

### Configuration

```
helm show values istio/istiod \
  --version $ISTIOD_CHART_VERSION > kubernetes/gateway-api/istio/default-values.yaml
```

### Check Installation

Now we should have the Istio control plane / controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implements some gateway API CRD's. </br>

```shell
# check deployed status
helm ls -n istio-system

# check the controller pods
kubectl get pods -n istio-system --output wide

# check the controller pod logs 
kubectl -n istio-system logs -l app.kubernetes.io/instance=istiod

```

## Install an Istio Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/istio/01-gatewayclass.yaml
```

## Install an Istio Gateway

```shell
kubectl apply -f kubernetes/gateway-api/istio/02-gateway.yaml
```

When we apply the gateway, we get a new gateway api pod. </br>
Unlike Traefik and Envoy, in Istio, the gateway pod is created in the same namespace as the gateway

```shell
# check the new gateway-api pod - this time its in the same namespace as the gateway object!
kubectl -n default get pods

# we also have a new service
kubectl -n default get svc
```

### Gateway Configuration

For global Gateway class defaults, we can use a `ConfigMap` with a special label in the `istio-system` namespace. </br>
This allows us to set global defaults for all our gatways. </br>

```shell
kubectl apply -f kubernetes/gateway-api/istio/01.1-gatewayclass-config.yaml
```

For Gateways, we can use the `parametersRef` field under `infrastructure` and use a `ConfigMap` to configure it further. </br>
For example, we can change the `deployment` or `service` of the instance like so:

```shell
kubectl apply -f kubernetes/gateway-api/istio/02.1-gateway-config.yaml

# we should see 2 replicas 
kubectl -n default get pods

# port forward for access (get the correct service)
kubectl -n default port-forward svc/gateway-api-istio 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## Istio Features for Traffic Management

Istio has been around for a lot longer than Gateway API has. Therefore Istio has solved much of the traffic management challenges using it's own CRDs, such as `VirtualService`. </br>

There is a very important difference between Istio and the other Gateway APIs we've covered so far. </br>
Istio has made a strict separation between it's own Gateway (`networking.istio.io/v1`) and Gateway API (`gateway.networking.k8s.io/v1`) </br>

This means, their CRDs such as `VirtualService` does not attach to the new Kubernetes Gateway API like `Gateway` and `HTTPRoute`. </br>

Therefore, to use Gateway API featuers, one should use Gateway API and not Istio Gateways. </br>

[Virtual Services](https://istio.io/latest/docs/reference/config/networking/virtual-service/) are used throughout Istio to manage traffic. </br>

### CORS

As an example, `VirtualService` can be used to deal with CORS (Cross-Origin Resource Sharing), when the browser tries to make an API call to another domain. Like `localhost` calling our DEV environment which is a common challenge for developers. </br>

```shell
kubectl apply -f kubernetes/gateway-api/istio/08-virtualservice-cors.yaml
```

🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧🚧

