# Introduction to Gateway API

In this video, we will dive into the features of Gateway API. </br>
Once you finish this video, click [here]() to deep dive each of the Gateway API controllers. </br>

[Documentation](https://gateway-api.sigs.k8s.io/)

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name gatewayapi --image kindest/node:v1.34.0
```

Test our cluster and makes sure `kubectl` is configured for it:

```
kubectl get nodes
NAME                       STATUS   ROLES           AGE   VERSION
gatewayapi-control-plane   Ready    control-plane   40s   v1.34.0
```

## Gateway API CRDs

The Kubernetes Gateway API is not installed on kubernetes by default. My guess is it may at some point. </br>
For now, we can grab it from the [Gateway API SIGS Guide](https://gateway-api.sigs.k8s.io/guides/#installing-gateway-api)

<i><b>Important Note: </b>At the time of recording this guide, we'd like to look at as many features as possible, hence the `experimental` install is used.</i> </br>

```shell
kubectl apply --server-side -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.4.0/experimental-install.yaml
```

This will install (Stable Channel):

* Gateway Classes: `kubectl get gatewayclass`
* Gateways: `kubectl get gateway`
* HTTP Routes:  `kubectl get httproute`

These APIs are part of the Experimental Channel:
* TLS Routes: `kubectl get tlsroute`
* TCP Routes: `kubectl get tcproute`
* UDP Routes: `kubectl get udproute`

<b>Note: Gateway API is very new, and all of the above is subject to change quite rapidly </br></b>

## Setup some example applications

The following will deploy a `deployment`, `service` and require `configMap` and `secret` for the applications to work.

```shell
# deploy example apps
kubectl apply -f python/deployment.yaml 
kubectl apply -f golang/deployment.yaml
kubectl apply -f kubernetes/gateway-api/web-app.yaml

# check example apps
kubectl get pods
NAME                             READY   STATUS    RESTARTS
go-deploy-b9c69978d-529qb        1/1     Running   0
go-deploy-b9c69978d-jfb2b        1/1     Running   0
python-deploy-54cbfc948b-5ch7w   1/1     Running   0
python-deploy-54cbfc948b-qh22f   1/1     Running   0
web-app-67fbb5d844-68wq4         1/1     Running   0

kubectl get service
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
go-svc       ClusterIP   10.96.173.22   <none>        5000/TCP
python-svc   ClusterIP   10.96.113.49   <none>        5000/TCP
web-app      ClusterIP   10.96.44.18    <none>        80/TCP
```

### Create test Domains 

We also need to imagine we have a domain called `example-app.com` , so let's set that up on our hosts file

```log
127.0.0.1  example-app.com
127.0.0.1  example-app-go.com
127.0.0.1  example-app-python.com
```


## Install a Gateway API controller

To use the Gateway API features in Kubernetes, you need a controller that implements the above CRDs. </br>
In this introduction guide I will use an existing Gateway API controller called Traefik. </br>

<i><b>Note:</b> Keep in mind that this introduction has no dependency on Traefik specifically, therefore any controller that supports Gateway API can be used.
At the bottom of this guide, I will provide guides on each of the popular Gateway API implementations.
</i>

A Basic Gateway API controller install:

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

As we don't have a LoadBalancer service in `kind`, let's `port-forward` so we can pretend we have one 

```shell

# check the pods
kubectl -n traefik get pods 

# check the logs 
kubectl -n traefik logs -l app.kubernetes.io/instance=traefik-traefik

# port forward for access
kubectl -n traefik port-forward svc/traefik 80
```


## Install a Gateway Class

To start enabling traffic to our newly created apps, we will start with installing a Gateway Class. </br>

Note that we use a Traefik Class in our example. </br>

[Documentation](https://gateway-api.sigs.k8s.io/api-types/gatewayclass/)

`GatewayClass` is a cluster-scoped resource defined by the infrastructure provider. This resource represents a class of Gateways that can be instantiated. </br>

```shell
kubectl apply -f kubernetes/gateway-api/traefik/01-gatewayclass.yaml

# check
kubectl get gatewayclass

# describe
kubectl describe gatewayclass
```

## Install a Gateway 

Next we need to install a Gateway that implements our Gateway Class </br>
Note that we use a Traefik Gateway in our example. </br>

[Documentation](https://gateway-api.sigs.k8s.io/api-types/gateway)

This gateway lives in the same namespace as the routes and applications

```shell
kubectl apply -f kubernetes/gateway-api/traefik/02-gateway.yaml
```

## Traffic Management Features : HTTP Routes

[Documentation](https://gateway-api.sigs.k8s.io/api-types/httproute/)

The important fields on HTTP Route we will cover:

* `parentRefs`
* `sectionName` 
* `hostnames`
* `rules` 
* `matches`
* `filters`

Feature Table: 

| Feature | Example |
| ---|---|
| Route by Hostname | [example](#route-by-hostname)|
| Route by Path | [example](#route-by-path) | 
| Route using URL Rewrite | [example](#route-using-url-rewrite)|
| Header Modification | [example](#requestresponse-header-manipulation)|
| HTTPS & TLS | [example](#https-and-tls) |

For traffic management, we can take a look at some basic HTTP routes.</br>

### Route by Hostname

We can route by host. </br>
This will route all traffic that matches the `Host` header with the `hostnames` field: </br>


http://example-app-python.com/ 👉🏽 http://python-svc:5000 </br>
http://example-app-go.com/ 👉🏽 http://go-svc:5000 </br>

```shell
kubectl apply -f kubernetes/gateway-api/03-httproute-by-hostname.yaml
```

### Route by Path

We can also route by host and path with different matching strategies. </br>

Exact: </br>
http://example-app-python.com/ 👉🏽 http://python-svc:5000/ </br>
http://example-app-go.com/ 👉🏽 http://go-svc:5000/ </br>

PathPrefix: </br>
http://example-app-python.com/* 👉🏽 http://python-svc:5000/* </br>
http://example-app-go.com/* 👉🏽 http://go-svc:5000/* </br>


```shell
kubectl apply -f kubernetes/gateway-api/04-httproute-by-path-exact.yaml
```

### Route using URL Rewrite

We can rewrite the hostname or URL using URL rewrite. </br>
This way, we can combine our services under one domain and our controller can act as a true API gateway:

http://example-app.com/api/python 👉🏽 http://python-svc:5000/ </br>
http://example-app.com/api/go 👉🏽 http://go-svc:5000/ </br>
As well as: </br>
http://example-app.com/api/go/status 👉🏽 http://go-svc:5000/status </br>

```shell
kubectl apply -f kubernetes/gateway-api/05-httproute-by-path-rewrite.yaml
```

### Request\Response Header Manipulation

With Gateway API, you can modify request and response headers. </br>
This is possible with the `ResponseHeaderModifier` filter </br>

At the time of this recording, Gateway API does not natively support CORS. </br>
Even with it in the Experimental channel, many controllers do not support it yet. </br>

Let's do a basic CORS header modification for our Go HTTPRoute </br>

```shell
kubectl port-forward svc/web-app 8000:80
```

The Header modification:

```shell
kubectl apply -f kubernetes/gateway-api/06-httproute-header-modify.yaml
```

### HTTPS and TLS

In my video, I generate a test TLS cert using [mkcert](https://github.com/FiloSottile/mkcert)

```shell
curl -L https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64 -o mkcert && chmod +x mkcert && mv mkcert /usr/local/bin/

#linux
export CAROOT=${PWD}/kubernetes/gateway-api/tls
#windows
$env:CAROOT = "${PWD}" + "\kubernetes\gateway-api\tls"

mkcert -key-file kubernetes/gateway-api/tls/key.pem -cert-file kubernetes/gateway-api/tls/cert.pem example-app.com

mkcert -install
```

Now that we have a TLS cert, we can create a Kubernetes secret to store it:

```
kubectl create secret tls secret-tls -n default --cert kubernetes/gateway-api/tls/cert.pem --key kubernetes/gateway-api/tls/key.pem
```

We need to 

* Adjust our Gateway, to enable the TLS Listener first!
* Then apply the TLS listener in our HTTP Route to enable TLS, using `sectionName`

```shell
kubectl apply -f kubernetes/gateway-api/07-httproute-tls.yaml
```

Let's `port-forward` to 443 since that is where TLS is exposed:

```shell
kubectl -n traefik port-forward svc/traefik 443
```

Result: </br>
https://example-app.com/api/go 👉🏽 http://go-svc:5000/ </br>

Checkout [More Official Guides](https://gateway-api.sigs.k8s.io/guides/) on the Kubernetes Gateway API SIGs page.</br>

## Infrastructure Labels

A useful feature is the ability to customize infrastructure under the hood for Gateways. </br>
For example, cloud load balancers etc. 

We can use [Infrastructure Labels](https://kubernetes.io/blog/2023/11/28/gateway-api-ga/#gateway-infrastructure-labels) to do so. This will set annotations or labels on any infrastructure that gets created. </br>

## Gateway API Controllers Guides

| Name | Guide | Video |
|------|-------|-------|
| Traefik | [Introduction to Traefik Gateway API](./traefik/README.md)   | #TODO      |
| Envoy | [Introduction to Envoy Gateway API](./envoy/README.md)   | #TODO      |
| Istio | #TODO   | #TODO      |
