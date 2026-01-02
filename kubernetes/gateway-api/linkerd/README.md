# Introduction to Linkerd: Gateway API

<a href="https://youtu.be/xxxxxx" title="linkerd"><img src="https://i.ytimg.com/vi/xxxx/hqdefault.jpg" width="40%" alt="linkerd" /></a>

#TODO

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

## What is Linkerd

#TODO

## Linkerd: Gateway API controller

Linkerd provides Gateway API support included within its single installation. </br>
Unlike some other gateways, that provide Gateway API support through a separate `helm` chart, Linkerd is all one single simple install. </br>

The `helm` chart relies on the Gateway API CRD's we installed in our cluster already </br>

We can checkout the official [install via helm](https://linkerd.io/2.19/tasks/install-helm/) for more instructions. </br>

### Installation 

We firstly need to create mTLS certificates that Linkerd will use for secure internal communication in its service mesh

```shell
apk add step-cli #alpine linux

#generate CA
step certificate create root.linkerd.cluster.local /tmp/ca.crt /tmp/ca.key \
--profile root-ca --no-password --insecure

#generate intermediate cert for csr's
step certificate create identity.linkerd.cluster.local /tmp/issuer.crt /tmp/issuer.key \
--profile intermediate-ca --not-after 8760h --no-password --insecure \
--ca /tmp/ca.crt --ca-key /tmp/ca.key


```
We can install the helm chart with `helm install` or upgrade it with `helm upgrade`
Please note that I will be using the linkerd edge release in this demo to ensure I have the latest functionality. For production, you may want to use the stable release instead. </br>

```shell

CRD_CHART_VERSION="2025.12.3"
CHART_VERSION="2025.12.3"

helm repo add linkerd-edge https://helm.linkerd.io/edge
helm repo update
helm search repo linkerd-edge --versions | grep linkerd-edge/linkerd-crds | head
helm search repo linkerd-edge --versions | grep linkerd-edge/linkerd-control-plane | head

# get values files
helm show values linkerd-edge/linkerd-control-plane > kubernetes/gateway-api/linkerd/default-values.yaml

helm install linkerd-crds linkerd-edge/linkerd-crds \
  --version $CRD_CHART_VERSION \
  --namespace linkerd \
  --create-namespace

helm upgrade linkerd-control-plane \
  --version $CHART_VERSION \
  --namespace linkerd \
  --create-namespace \
  --set-file identityTrustAnchorsPEM=/tmp/ca.crt \
  --set-file identity.issuer.tls.crtPEM=/tmp/issuer.crt \
  --set-file identity.issuer.tls.keyPEM=/tmp/issuer.key \
  --values kubernetes/gateway-api/linkerd/values.yaml \
  linkerd-edge/linkerd-control-plane

```

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [default-values.yaml](./default-values.yaml)

Linkerd also provide a High Availability YAML file for HA deployments which I found [here on Github](https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-control-plane/values-ha.yaml). </br>
So in this guide I will base my demo off that with the `controllerReplicas` set to 1 for demo purpose as I am running a single node cluster. </br>

### Check Installation

Now we should have the Envoy gateway API controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implement some gateway API CRDs. </br>

```shell
# check the controller pods
kubectl -n linkerd get pods

# check the controller pod logs 
kubectl -n linkerd logs -l app.kubernetes.io/instance=linkerd-gateway

#use the linkerd CLI to check the install
export LINKERD2_VERSION=edge-25.12.3
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/install-edge | sh

linkerd check

```

## Install an Linkerd Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/linkerd/01-gatewayclass.yaml
```

## Install an Linkerd Gateway

```shell
kubectl apply -f kubernetes/gateway-api/linkerd/02-gateway.yaml
```

When we apply the gateway, we get a new gateway api pod. 

```shell
# check the new gateway-api pod
kubectl -n linkerd get pods

# we also have a new service
kubectl -n linkerd get svc
```

### Gateway Configuration

#TODO : gateway and class configuration and overrides

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## Gateway API Extensions

#TODO: custom extensions & policies
