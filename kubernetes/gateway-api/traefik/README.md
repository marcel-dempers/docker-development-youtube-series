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

## Install a Traefik: Gateway API controller

Traefik has a lot of features, including Ingress controller functionality. </br>
In this guide we will turn on the Gateway API feature in Traefik and use that specifically. </br>

Install:

```shell
CHART_VERSION="37.3.0" # traefik version v3.6.0
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
helm search repo traefik --versions


helm install traefik traefik/traefik \
  --version $CHART_VERSION \
  --values kubernetes/gateway-api/values.yaml \
  --namespace traefik \
  --create-namespace

```

Upgrade:

```shell
helm upgrade traefik traefik/traefik \
  --version $CHART_VERSION \
  --values kubernetes/gateway-api/values.yaml \
  --namespace traefik
```