# Introduction to HAProxy: Gateway API

#TODO
<a href="https://youtu.be/5clOvqDh-zA" title="haproxy"><img src="https://i.ytimg.com/vi/5clOvqDh-zA/hqdefault.jpg" width="40%" alt="haproxy" /></a>

## Prerequisites

To get started, you will need to follow the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<i><b>Important Note: According to testing and [HAProxy documentation](https://www.haproxy.com/documentation/kubernetes-ingress/gateway-api/enable-gateway-api/#deploy-gateway-api-resources), it currently only supports version v0.5.1</b></i>
<i>Tried `GatewayClass` with `gateway.networking.k8s.io/v1` and does not work</i>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs

```shell
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.5.1/experimental-install.yaml
```

* Deploy example apps to our cluster
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## Introduction

Let's take a look at the [Official Documentation](https://www.haproxy.com/documentation/). </br>

At the time of this guide the HAProxy Gateway API solution is built alongside the HAProxy Kubernetes Ingress controller. So its one solution that can achieve both Ingress and Gateway functionality. </br>

The HAProxy Unified Gateway is still under development at the time of this guide. </br>

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [default-values.yaml](./default-values.yaml)

### Installation 

We can install the helm chart with `helm install` or with `kubectl` as stated in the documentation. Important to know that the below Helm chart installs the ingress controller. In order to enable Gateway API functionality, we need to turn it on using our values file. </br>

```shell

CHART_VERSION="1.48.0"
helm repo add haproxytech https://haproxytech.github.io/helm-charts
helm repo update
helm search repo haproxytech --versions

# see the default configuration
helm show values haproxytech/kubernetes-ingress \
> kubernetes/gateway-api/haproxy/default-values.yaml

#apply the cluster roles and bindings (RBAC) to give HAProxy access to Gateway resources
kubectl apply -f https://raw.githubusercontent.com/haproxytech/kubernetes-ingress/master/deploy/tests/config/experimental/gwapi-rbac.yaml

helm install haproxy-kubernetes-ingress haproxytech/kubernetes-ingress \
  --version $CHART_VERSION \
  --create-namespace \
  --namespace haproxy-controller \
  --values kubernetes/gateway-api/haproxy/values.yaml

```

### Check Installation

Now we should have the HAProxy controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implement some gateway API CRDs. </br>

```shell
# check the controller pods
kubectl -n haproxy-controller get pods

# check the controller pod logs 
kubectl -n haproxy-controller logs -l app.kubernetes.io/name=kubernetes-ingress

```

## Install an HAProxy Gateway Class

Please note `apiVersion: gateway.networking.k8s.io/v1alpha2` is used and `v1` does not work. :

```shell
kubectl apply -f kubernetes/gateway-api/haproxy/01-gatewayclass.yaml
```

## Install an HAProxy Gateway

```shell
kubectl apply -f kubernetes/gateway-api/haproxy/02-gateway.yaml
```

## Traffic Routing 

At the time of this guide, HAProxy only supports TCPRoute, and not HTTPRoute like the other controllers we covered thus far. </br>

```shell
kubectl apply -f kubernetes/gateway-api/haproxy/03-tcproute.yaml
```

```shell
kubectl -n haproxy-controller port-forward svc/haproxy-kubernetes-ingress 80
```