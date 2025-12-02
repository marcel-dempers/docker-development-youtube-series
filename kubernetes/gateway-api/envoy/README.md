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


Check our installation

```shell
# check the pods
kubectl -n envoy-gateway-system get pods

# check the logs 
kubectl -n envoy-gateway-system logs -l app.kubernetes.io/instance=envoy-gateway

# port forward for access
kubectl -n envoy-gateway-system get svc
NAME                                 TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                                    
envoy-default-gateway-api-30a1473e   LoadBalancer   10.96.16.176   <pending>     80:31127/TCP  
envoy-gateway                        ClusterIP      10.96.27.58    <none>        18000/TCP,18001/TCP,18002/TCP,19001/TCP,9443/TCP

kubectl -n envoy-gateway-system port-forward svc/envoy-default-gateway-api-30a1473e 80
```

## Install an Envoy Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/envoy/01-gatewayclass.yaml
```

## Install an Envoy Gateway

```shell
kubectl apply -f kubernetes/gateway-api/envoy/02-gateway.yaml
```

## Security Policy

### CORS 
### Basic Auth

### API Auth 

### External Auth