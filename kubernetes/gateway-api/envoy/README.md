# Introduction to Envoy: Gateway API

<a href="https://youtu.be/me_5W_Q4ZWg" title="envoy"><img src="https://i.ytimg.com/vi/me_5W_Q4ZWg/hqdefault.jpg" width="40%" alt="envoy" /></a>

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
  --values kubernetes/gateway-api/envoy/values.yaml \
  --namespace envoy-gateway-system \
  --create-namespace
```

Upgrade:

```shell
helm upgrade envoy-gateway oci://docker.io/envoyproxy/gateway-helm \
  --values kubernetes/gateway-api/envoy/values.yaml \
  --version $CHART_VERSION \
  --namespace envoy-gateway-system
```

### Configuration

Most of the Gateway API controllers are installed using `helm`. </br>
Before we install it, let's take a look at the [default-values.yaml](./default-values.yaml)

### Check Installation

Now we should have the Envoy gateway API controller up and running. </br>
This is not the gateway itself, but the controller that will manage the CRDs we get access to and implements some gateway API CRD's. </br>

```shell
# check the controller pods
kubectl -n envoy-gateway-system get pods

# check the controller pod logs 
kubectl -n envoy-gateway-system logs -l app.kubernetes.io/instance=envoy-gateway

```

## Install an Envoy Gateway Class

```shell
kubectl apply -f kubernetes/gateway-api/envoy/01-gatewayclass.yaml
```

## Install an Envoy Gateway

```shell
kubectl apply -f kubernetes/gateway-api/envoy/02-gateway.yaml
```

When we apply the gateway, we get a new gateway api pod. 

```shell
# check the new gateway-api pod
kubectl -n envoy-gateway-system get pods

# we also have a new service
kubectl -n envoy-gateway-system get svc
```

### Gateway Configuration

`EnvoyProxy` is the CRD that allows us to configure each gateway API instance. </br>
For example, we can change the `deployment` or `service` of the instance like so:

```shell
kubectl apply -f kubernetes/gateway-api/envoy/02.1-gateway-config.yaml

# we should see 2 replicas 
kubectl -n envoy-gateway-system get deploy
kubectl -n envoy-gateway-system get pods

# port forward for access (get the correct service)
kubectl -n envoy-gateway-system get svc
kubectl -n envoy-gateway-system port-forward svc/envoy-gateway-default 80
```

## HTTP Traffic management

Feel free to quickly run through the basic [traffic management table](../README.md#traffic-management-features--http-routes) for using `HTTPRoute` routing for traffic. </br>

<i>Note: HTTPRoute features are not specific to this controller and should be available to any other gateway API controller that you choose.</i>

## Gateway API Extensions

Envoy Gateway has it's own extensions on top of the native Gateway API features we've seen so far, using custom CRD's.

### Client Traffic Policies

#### Connection Limit

Envoy allows us to limit connections to gateways & listeners. </br>
We can use tools like `hey` to demonstrate connection limit.

```shell

hey -c 10 -q 1 -z 10s -host "example-app.com" http://localhost:80/api/go/status

kubectl apply -f kubernetes/gateway-api/envoy/08-clientpolicy-connectionlimit.yaml

hey -c 10 -q 1 -z 10s -host "example-app.com" http://localhost:80/api/go/status
```

#### Mutual TLS

Currently we have a listener for TLS (HTTPs). For normal TLS, the client does not have to pass a private key. </br>
Envoy supports mutual TLS

```shell

# tell Envoy about the CA cert 
kubectl create secret generic secret-tls-ca --from-file=rootCA.pem=kubernetes/gateway-api/tls/rootCA.pem

# generate client certs for curl
mkcert -client client-user.com

kubectl apply -f kubernetes/gateway-api/envoy/09-clientpolicy-mtls.yaml

# browser TLS should now be broken "Access to example-app.com was denied"
# This is because a client needs to pass certificates to trust 

curl -v -H "Host: example-app.com" \
--cert kubernetes/gateway-api/tls/client-user.com-client.pem \
--key kubernetes/gateway-api/tls/client-user.com-client-key.pem \
--cacert kubernetes/gateway-api/tls/rootCA.pem \
https://example-app.com/api/go/status

```

### Backend Traffic Policies

#### Local and Global Rate limits

We can rate limit each HTTP route using Local Rate Limit:

```shell

kubectl apply -f kubernetes/gateway-api/envoy/10-backendpolicy-ratelimit.yaml

for i in {1..4}; do
curl -I --header "x-user-id: one" http://example-app.com/api/go/status ; sleep 1;
done

```

### Security Policies

### CORS 

```shell

# port-forward to set pre-cors
kubectl port-forward svc/web-app 8000:80

kubectl apply -f kubernetes/gateway-api/envoy/11-securitypolicy-cors.yaml

```

### Basic Auth

```shell

#cleanup client policies to avoid conflict
kubectl delete clienttrafficpolicies mtls-policy
kubectl delete SecurityPolicy go-route-cors

# port-forward to 443 
kubectl -n envoy-gateway-system port-forward svc/envoy-gateway-default 443

#install htpasswd (alpine)
apk add apache2-utils
htpasswd -cbs .htpasswd bob password123

kubectl create secret generic basic-auth --from-file=.htpasswd

kubectl apply -f kubernetes/gateway-api/envoy/12-securitypolicy-basicauth.yaml

```

### API Auth 

```shell

#cleanup previous basic auth policy
kubectl delete SecurityPolicy go-route-basicauth

kubectl apply -f kubernetes/gateway-api/envoy/13-securitypolicy-apiauth.yaml

# port-forward
kubectl -n envoy-gateway-system port-forward svc/envoy-gateway-default 443

```

You can see `Client authentication failed.` in the browser. 
We need to pass the header 

```
curl -v -H "Host: example-app.com" \
-H 'x-api-key: secret123' \
--cert kubernetes/gateway-api/tls/client-user.com-client.pem \
--key kubernetes/gateway-api/tls/client-user.com-client-key.pem \
--cacert kubernetes/gateway-api/tls/rootCA.pem \
https://example-app.com/api/go/status
```

Logs:

```shell
kubectl -n envoy-gateway-system logs -l app.kubernetes.io/managed-by=envoy-gateway
```