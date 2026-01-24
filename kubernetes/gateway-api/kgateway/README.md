# Introduction to kgateway: Gateway API

#TODO 
<a href="https://youtu.be/xxxxx" title="kgateway"><img src="https://i.ytimg.com/vi/xxxxx/hqdefault.jpg" width="40%" alt="kgateway" /></a>

## Prerequisites

To get started, you will need to follow the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs
* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

Let's start with the [Official Documentation](https://kgateway.dev/docs/envoy/latest/quickstart/)

## kgateway: Gateway API controller

kgateway is feature rich when it comes to L7 HTTP routing. </br>
In this guide we will turn on the Gateway API feature in kgateway and use that specifically. </br>

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [values.yaml](./values.yaml)

kgateway has a section on [Gateway Customization](https://kgateway.dev/docs/envoy/main/setup/) where we can adjust gateway parameters. </br>

Gateway API allows us to inject parameters using the `infrastructure.parametersRef` field </br>

We can customize our gateway:

```shell
kubectl apply -f kubernetes/gateway-api/kgateway/02-gateway-config.yaml
```


### Installation 

To get the version you want, you can use the kgateway [Github Release page](https://github.com/kgateway-dev/kgateway/releases/tag/v2.1.2). </br>
For this guide, we will use 2.1.2. </br>

Install:

```shell
CHART_VERSION="v2.1.2"

helm show chart oci://cr.kgateway.dev/kgateway-dev/charts/kgateway --version $CHART_VERSION
helm show chart oci://cr.kgateway.dev/kgateway-dev/charts/kgateway-crds --version $CHART_VERSION

helm show values oci://cr.kgateway.dev/kgateway-dev/charts/kgateway --version $CHART_VERSION > kubernetes/gateway-api/kgateway/default-values.yaml

# deploy CRDs
helm upgrade kgateway-crds oci://cr.kgateway.dev/kgateway-dev/charts/kgateway-crds \
  --version ${CHART_VERSION} \
  --namespace kgateway-system \
  --create-namespace \
  --install

# Deploy kgateway control plane
helm upgrade kgateway oci://cr.kgateway.dev/kgateway-dev/charts/kgateway \
  --version $CHART_VERSION \
  --values kubernetes/gateway-api/kgateway/values.yaml \
  --namespace kgateway-system \
  --install \
  --create-namespace
  
```

Check our installation

```shell
# check the pods
kubectl -n kgateway-system get pods

# check the logs 
kubectl -n kgateway-system logs -l app.kubernetes.io/name=kgateway
```

## Install a kgateway Gateway Class

The `helm` chart will create a `GatewayClass` for us as per [docs](https://kgateway.dev/docs/envoy/latest/setup/default/#gatewayclass). </br>

```shell
kubectl get gc kgateway -o yaml
```

## Install a kgateway Gateway

```shell
kubectl apply -f kubernetes/gateway-api/kgateway/01-gateway.yaml

# we can see our gateway dataplane pods
kubectl get pods
kubectl get svc

# port forward for access
kubectl -n default port-forward svc/gateway-api 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## TrafficPolicy

kgateway uses `TrafficPolicy` as their custom CRD to perform traffic management actions that are not supported by native Gateway API. </br>

There are many [TrafficPolicy Features](https://kgateway.dev/docs/envoy/main/reference/api/#trafficpolicy) available. 

The cool thing here is that these TrafficPolicies are now pluggable into Gateway API HTTPRoutes. </br>

In `HTTPRoute` objects, under `filters`, we can add extensions using something like:

```yaml
filters:
- type: ExtensionRef
  extensionRef:
    group: gateway.kgateway.dev
    kind: TrafficPolicy
    name: cors-feature
```

### Basic Auth

We can also add HTTP basic authentication to allow users to access secured resources protected by username and passwords. </br>

The Basic Auth TrafficPolicy points to a secret which supports Kubernetes basic auth secret type. </br>

```shell
kubectl apply -f kubernetes/gateway-api/kgateway/02-trafficpolicy-basicauth.yaml
```
