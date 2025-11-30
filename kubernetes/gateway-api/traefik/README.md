# Introduction to Traefik: Gateway API

## Prerequisites

To get started, you will need to follow the the [Introduction to Gateway API](../README.md)

<b>In the above guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRD
* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## Install a Traefik Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/traefik/01-gatewayclass.yaml
```

## Install a Traefik Gateway

```shell
kubectl apply -f kubernetes/gateway-api/traefik/02-gateway.yaml
```

Let's start with the [Official Documentation](https://doc.traefik.io/traefik/)

## Traefik: Gateway API controller

Traefik has a lot of features, including Ingress controller functionality. </br>
In this guide we will turn on the Gateway API feature in Traefik and use that specifically. </br>

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [values.yaml](./values.yaml)

Traefik allows us to configure many options. Some I find quite important:

* Access logs - fields and format
* Ports to use for incoming traffic
* Infrastructure annotations \ labels
* Control Gateway API
  * Default CRDs
  * Default GatewayClass
  * Default Gateways

It's always good to get a grip on the [default helm values](https://github.dev/traefik/traefik-helm-chart/blob/master/traefik/values.yaml)

The values file is used to customise Traefik, underlying pods, deployments and services as well as turning features on and off. </br>

We can go ahead and install the controller to start with and use `helm upgrade` to make changes as we need to. </br>

### Installation 
Install:

```shell
CHART_VERSION="37.3.0" # traefik version v3.6.0
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
helm search repo traefik --versions


helm install traefik traefik/traefik \
  --version $CHART_VERSION \
  --values kubernetes/gateway-api/traefik/values.yaml \
  --namespace traefik \
  --create-namespace

```

Upgrade:

```shell
helm upgrade traefik traefik/traefik \
  --version $CHART_VERSION \
  --values kubernetes/gateway-api/traefik/values.yaml \
  --namespace traefik
```

Check our installation

```shell
# check the pods
kubectl -n traefik get pods

# check the logs 
kubectl -n traefik logs -l app.kubernetes.io/instance=traefik-traefik

# port forward for access
kubectl -n traefik port-forward svc/traefik 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## Middlewares

One of the key strengths of using Traefik is that they have a catalogue of middlewares that are plugins to perform popular common actions on traffic. </br>

There are many [Middlewares](https://doc.traefik.io/traefik/reference/routing-configuration/http/middlewares/overview/) available. 

The cool thing here is that these Middlewares are now pluggable into Gateway API HTTPRoutes. </br>

In `HTTPRoute` objects, under `filters`, we can add extensions using something like:

```yaml
filters:
- type: ExtensionRef
  extensionRef:
    group: traefik.io
    kind: Middleware
    name: add-prefix
```

### Prefix

Note: To use `ExtensionRef` in Traefik you need to enable `providers.kubernetesCRD.enabled = true`. We have that enabled in our values.yaml file.

Expectation: </br>

http://example-app.com/anything 👉🏽 http://go-svc:5000/prefix/anything </br>

```shell
kubectl apply -f kubernetes/gateway-api/traefik/08-middleware-addprefix.yaml
```

We can follow the log of our upstream to see the HTTP request:

```shell
kubectl logs -f -l app=go-app
```

### Basic Auth

We can also add HTTP basic authentication to allow users to access secured resources protected by username and passwords. </br>

The Basic Auth middleware points to a secret which supports Kubernetes basic auth secret type. </br>

```shell
kubectl apply -f kubernetes/gateway-api/traefik/09-middleware-basicauth.yaml
```

### Headers

Traefik allows us to manipulate headers too using Middleware component called `headers` </br>
We can use this to handle scenarios like CORS (Cross-Origin Resource Sharing)

Let's access our web app directly over `localhost` which makes a call to `example-app.com` and should be blocked by CORS. </br>

```shell
kubectl port-forward svc/web-app 8000:80
```

Once we apply the Middleware update to our HTTPRoute, we can perform cross origin API calls:

```shell
kubectl apply -f kubernetes/gateway-api/traefik/10-middleware-headers.yaml
```