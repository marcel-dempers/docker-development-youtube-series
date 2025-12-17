# Introduction to NGINX Gateway Fabric: Gateway API

<a href="https://youtu.be/xxxxxxxx" title="nginx-fabric"><img src="https://i.ytimg.com/vi/xxxxxxxx/hqdefault.jpg" width="40%" alt="nginx-fabric" /></a>

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

## What is NGINX Gateway Fabric

NGINX Gateway Fabric is an [open-source project](https://github.com/nginx/nginx-gateway-fabric) that provides an implementation of the Gateway API using NGINX as the data plane. 

Let's checkout the [Official Documentation](https://docs.nginx.com/nginx-gateway-fabric/overview/gateway-architecture/)

## Install a NGINX Gateway Class

In our other guides, we deployed and managed a Gateway Class outside of the controller installation. </br>
NGINX Fabric is a little different as their `GatewayClass` is managed by helm and there are extra configurations managed by the chart that gets injected into the `GatewayClass` under the `parametersRef` field. </br>

If you've been following this series you will know the `parametersRef` field is used to extend the class. NGINX uses a custom CRD called `NginxProxy` to customize the class, but these settings come from their `helm` chart. </br>

For example:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GatewayClass
metadata:
  name: nginx-gateway
  annotations:
    meta.helm.sh/release-name: ngf
    meta.helm.sh/release-namespace: nginx-gateway
  labels:
    app.kubernetes.io/managed-by: Helm
spec:
  controllerName: gateway.nginx.org/nginx-gateway-controller
  parametersRef:
    group: gateway.nginx.org
    kind: NginxProxy
    name: ngf-proxy-config
    namespace: nginx-gateway
```


## NGINX Fabric: Gateway API controller

NGINX Fabric covers most of the core Gateway API features. </br>
It also supports custom features using policies </br>

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [values.yaml](./values.yaml)

The Helm chart allows us to configure many options. Some I find quite important:

* Control plane deployment & Service
* Data plane deployment & Service modification
* Ports to use for incoming traffic
* Control Gateway API
  * Default CRDs
  * Default GatewayClass
  * Default Gateways

There are [helm examples](https://github.com/nginx/nginx-gateway-fabric/tree/v2.2.2/examples/helm) on Github maintained by NGINX community. </br>

Here is the [helm chart documentation](https://github.com/nginx/nginx-gateway-fabric/blob/v2.2.2/charts/nginx-gateway-fabric/README.md)

It's always good to get a grip on the [default helm values](./default-values.yaml) to see what we can do with the chart. </br>

The values file is used to customise NGINX control & data plane, underlying pods, deployments and services as well as turning features on and off. </br>

We can go ahead and install the controller to start with and use `helm upgrade` to make changes as we need to. </br>

### Installation 
Install:

```shell
CHART_VERSION="2.2.2"

helm show values oci://ghcr.io/nginx/charts/nginx-gateway-fabric > kubernetes/gateway-api/nginx-fabric/default-values.yaml

helm show chart oci://ghcr.io/nginx/charts/nginx-gateway-fabric

helm install ngf oci://ghcr.io/nginx/charts/nginx-gateway-fabric \
  --values kubernetes/gateway-api/nginx-fabric/values.yaml \
  --version $CHART_VERSION \
  --namespace nginx-gateway \
  --create-namespace
```

Upgrade:

```shell
helm upgrade ngf oci://ghcr.io/nginx/charts/nginx-gateway-fabric \
  --version $CHART_VERSION \
  --namespace nginx-gateway
```

Check our installation

```shell
# check the pods
kubectl -n nginx-gateway get pods

# check the logs 
kubectl -n nginx-gateway logs -l app.kubernetes.io/instance=ngf

```

## Install a NGINX Gateway

So we have the `GatewayClass` managed by the `helm` chart, therefore we skip the `01-gatewayclass.yaml` and go straight to the gateway. </br>

This will deploy NGINX dataplane pods to our target `default` namespace:

```shell
kubectl apply -f kubernetes/gateway-api/nginx-fabric/02-gateway.yaml

# check gateway pods
kubectl get pods 

#check logs
kubectl logs -l app.kubernetes.io/instance=ngf

# port forward for access
kubectl port-forward svc/gateway-api-nginx-gateway 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## Check NGINX Dataplane Configuration

When we define `HTTPRoute`'s for our traffic, The NGINX control plane detects those and works with the data plane pods to get an `nginx.conf` for the desired traffic rules. </br>

If you are familiar with NGINX you will know all traffic rules are in the `nginx.conf`. So this solution transforms K8s CRDs to NGINX rules under the hood

```shell
# get a dataplane pod name
kubectl get pods 
pod=gateway-api-nginx-gateway-5899b6cd6f-dpkbh

# check the generated nginx configuration file
kubectl exec -it -n default $pod -- nginx -T
```

## Policy APIs

Kubernetes Gateway API provides [policy attachement](https://gateway-api.sigs.k8s.io/reference/policy-attachment/) that allows us to augment or add configuration to exising resources </br>


### Client Settings Policies
