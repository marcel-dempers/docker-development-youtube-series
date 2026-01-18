# Introduction to Linkerd: Gateway API

<a href="https://youtu.be/B9gKwJdWW7A" title="linkerd"><img src="https://i.ytimg.com/vi/B9gKwJdWW7A/hqdefault.jpg" width="40%" alt="linkerd" /></a>

## Prerequisites

To get started, you will need to follow the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs
* Deploy example apps to our cluster (below)

In this guide, we will deploy our videos application which has both north\south and east\west traffic: 

```shell
kubectl apply -f monitoring/opentelemetry/applications/playlists-api/
kubectl apply -f monitoring/opentelemetry/applications/playlists-db/
kubectl apply -f monitoring/opentelemetry/applications/videos-web/
kubectl apply -f monitoring/opentelemetry/applications/videos-api/
kubectl apply -f monitoring/opentelemetry/applications/videos-db/

```
* Have Domains for our traffic

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

## What is Linkerd

[Linkerd](https://linkerd.io) is a service mesh technology for Kubernetes that has [gateway api support](https://linkerd.io/2.19/features/gateway-api/) </br>
Linkerd is the first service mesh to achieve CNCF graduation status and is highly matured. </br>

In a distributed system with service-to-service communication, a service mesh uses various methods to intercept all inbound and outbound communication to and from a given service. </br>

So when a service makes an outbound call or has an inbound network call, the mesh is aware of it. </br>

Doing so, a service mesh provides many features:

* <b>observability</b>:  monitor network calls (metrics \ traces)
* <b>traffic</b> management: traffic routing features
* <b>resillience</b>: automatic retry, circuit breaker, rate limiting etc
* <b>security</b>: mutual TLS, authentication, policies etc

Each service mesh is different in terms of how it achieves this. Some meshes use side-cars to intercept and proxy traffic that go in and out of pods. Others use node level proxies and even eBPF to proxy at kernel level. </br>

Linkerd uses the sidecar approach with their proxy written in Rust. </br>

There are a few things that make Linkerd very different to other Gateway APIs we looked at. </br>
Linkerd is a big player in the [The GAMMA Initiative](https://gateway-api.sigs.k8s.io/mesh/gamma/) for Gateway API. </br>
Linkerd does not focus on north/south traffic and therefore we will use an alternative gateway API implementation for that. </br>
Linkerd's focus is more about mesh traffic, which is east\west. </br>

## Linkerd: Gateway API

Linkerd provides Gateway API support included within its single installation. </br>
Unlike some other gateways, that provide Gateway API support through a separate `helm` chart, Linkerd is all one single simple install. </br>

The `helm` chart relies on the Gateway API CRD's we installed in our cluster already </br>

We can checkout the official [install via helm](https://linkerd.io/2.19/tasks/install-helm/) for more instructions. </br>

### Installation 

We firstly need to create mTLS certificates that Linkerd will use for secure internal communication in its service mesh

```shell
apk add step-cli #alpine linux

#use the linkerd CLI to check the install
export LINKERD2_VERSION=edge-25.12.3
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/install-edge | sh
export PATH=$PATH:/root/.linkerd2/bin

linkerd check --pre

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

helm install linkerd-control-plane \
  --version $CHART_VERSION \
  --namespace linkerd \
  --create-namespace \
  --set-file identityTrustAnchorsPEM=/tmp/ca.crt \
  --set-file identity.issuer.tls.crtPEM=/tmp/issuer.crt \
  --set-file identity.issuer.tls.keyPEM=/tmp/issuer.key \
  --values kubernetes/gateway-api/linkerd/values.yaml \
  linkerd-edge/linkerd-control-plane

```

### Check Installation

Now we should have the Envoy gateway API controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implement some gateway API CRDs. </br>

```shell
# check the control plane pods
kubectl -n linkerd get pods

linkerd check

```

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [default-values.yaml](./default-values.yaml)

Linkerd also provide a High Availability YAML file for HA deployments which I found [here on Github](https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-control-plane/values-ha.yaml). </br>
So in this guide I will base my demo off that with the `controllerReplicas` set to 1 for demo purpose as I am running a single node cluster. </br>


## Install a Gateway Class & Gateway

Since Linderd does not have a gateway, we will be using Envoy from our [Envoy Gateway Guide](../envoy/README.md#envoy-gateway-api-controller)

```shell
# deploy gateway class
kubectl apply -f kubernetes/gateway-api/envoy/02.1-gateway-config.yaml

# deploy a gateway
kubectl apply -f kubernetes/gateway-api/envoy/02-gateway.yaml

```

When we apply the gateway, we get a new gateway api pod. 

```shell
# check the new gateway-api pod
kubectl -n envoy-gateway-system get pods

# we also have a new service
kubectl -n envoy-gateway-system get svc

# port forward for access to the new gateway
kubectl -n envoy-gateway-system port-forward svc/envoy-gateway-default 80
```

## Mesh our Services

To add our services to the Linkerd mesh, we need to annotate each deployment

```shell
kubectl get deploy playlists-api -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy playlists-db -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-api -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-db -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-web -o yaml | linkerd inject - | kubectl apply -f -

```

## Linkerd Dashboard

To access the `linkerd` dashboard, we need to install the viz extension

```
linkerd viz install | kubectl apply -f -
kubectl -n linkerd-viz port-forward svc/web 8084
```

## HTTPRoutes: North \ South Traffic

For north\south traffic, we will use Gateway API and setup a HTTP route

```shell
kubectl apply -f kubernetes/gateway-api/linkerd/03-httproute.yaml
```
## HTTPRoutes - East \ West Traffic 

For east\west traffic features:

Let's deploy a V2 of our videos-api that only returns a video about the Gateway API. </br>

```shell
# new V2
kubectl apply -f kubernetes/gateway-api/linkerd/04-east-west-app.yaml

# mesh the new app
kubectl get deploy videos-api-v2 -o yaml | linkerd inject - | kubectl apply -f -

# HTTPRoute that routes traffic east \ west
kubectl apply -f kubernetes/gateway-api/linkerd/04-httproute-east-west.yaml
```

You could use this for a number of reasons:

* Inject latency
* Inject faults to test applications
* Canary deployments
* Dynamic routing 