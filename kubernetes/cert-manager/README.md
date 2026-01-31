# Introduction to cert-manager for Kubernetes

<a href="https://youtu.be/hoLUigg4V18" title="certmanager"><img src="https://i.ytimg.com/vi/hoLUigg4V18/hqdefault.jpg" width="20%" alt="introduction to certmanager" /></a>

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```shell
kind create cluster --name certmanager --image kindest/node:v1.34.0
```

## Concepts 

It's important to understand the various concepts and new Kubernetes resources that <br/>
`cert-manager` introduces.

* Issuers [docs](https://cert-manager.io/docs/concepts/issuer/)
* Certificate [docs](https://cert-manager.io/docs/concepts/certificate/)
* CertificateRequests [docs](https://cert-manager.io/docs/concepts/certificaterequest/)
* Orders and Challenges [docs](https://cert-manager.io/docs/concepts/acme-orders-challenges/)

## Installation 

You can find the latest release for `cert-manager` on their [GitHub Releases page](https://github.com/cert-manager/cert-manager) <br/>

For this demo, I will use K8s 1.34 and `cert-manager` [v1.19.2](https://github.com/cert-manager/cert-manager/releases/tag/v1.19.2)

```shell
#test cluster access:
kubectl get nodes
NAME                        STATUS   ROLES           AGE     VERSION
certmanager-control-plane   Ready    control-plane   4m25s   v1.34.0

# note: in the legacy video, we used kubectl to install cert-manager.
#       this is now upgraded to use helm instead using the new OCI repo

# install cert-manager 
CHART_VERSION="v1.19.2"

# checkout the values
helm show values oci://quay.io/jetstack/charts/cert-manager > kubernetes/cert-manager/default-values.yaml

helm install \
  cert-manager oci://quay.io/jetstack/charts/cert-manager \
  --version ${CHART_VERSION} \
  --namespace cert-manager \
  --create-namespace \
  --set crds.enabled=true \
  --set config.enableGatewayAPI=true
```

## Cert Manager Resources

We can see our components deployed

```shell
kubectl -n cert-manager get all
NAME                                           READY   STATUS    RESTARTS   AGE
pod/cert-manager-75bb65b7b9-qmjt5              1/1     Running   0          6m57s
pod/cert-manager-cainjector-5cd89979d6-bs6m8   1/1     Running   0          6m57s
pod/cert-manager-webhook-8fc5dcf5f-xn5pq       1/1     Running   0          6m57s

NAME                              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)            AGE
service/cert-manager              ClusterIP   10.96.50.192    <none>        9402/TCP           6m57s
service/cert-manager-cainjector   ClusterIP   10.96.144.242   <none>        9402/TCP           6m57s
service/cert-manager-webhook      ClusterIP   10.96.223.22    <none>        443/TCP,9402/TCP   6m57s

NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/cert-manager              1/1     1            1           6m57s
deployment.apps/cert-manager-cainjector   1/1     1            1           6m57s
deployment.apps/cert-manager-webhook      1/1     1           1           6m57s

NAME                                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/cert-manager-75bb65b7b9              1         1         1       6m57s
replicaset.apps/cert-manager-cainjector-5cd89979d6   1         1         1       6m57s
replicaset.apps/cert-manager-webhook-8fc5dcf5f       1         1         1       6m57s

```

## Test Certificate Issuing 

Let's create some test certificates

```shell
kubectl create ns cert-manager-test

# create a self signed certificate
kubectl apply -f kubernetes/cert-manager/selfsigned/issuer.yaml
kubectl apply -f kubernetes/cert-manager/selfsigned/certificate.yaml

# see progress
kubectl describe certificate -n cert-manager-test
kubectl get secrets -n cert-manager-test

# cleanup
kubectl delete ns cert-manager-test
```

## Configuration 

https://cert-manager.io/docs/configuration/

## Setup my DNS

I can get the public IP address of my computer by running a simple command:

```
curl ifconfig.co
```

I can log into my DNS provider and point my DNS A record to my IP.<br/>
Also setup my router to allow 80 and 443 to come to my PC <br/>

<i>Note: If you are running in the cloud, your Ingress controller or Gateway API and Cloud provider will give you a
public IP and you can point your DNS to that accordingly.
</i>


## Deploy pods that require SSL\TLS

```shell
kubectl apply -f kubernetes/deployments/
kubectl apply -f kubernetes/services/

kubectl get pods
```

## Option 1: Using an Ingress Controller

Let's deploy an Ingress controller: <br/>
<i> Note: In the legacy vide, we used Kubernetes Ingress NGINX which is now deprecated. This guide is updated to use Traefik instead. You should be able to use any Ingress controller. </i>

```shell
CHART_VERSION="37.3.0"
helm repo add traefik https://helm.traefik.io/traefik
helm repo update

helm show values traefik/traefik > kubernetes/cert-manager/traefik-default-values.yaml

helm install traefik traefik/traefik \
  --version $CHART_VERSION \
  --namespace traefik \
  --set service.type=NodePort \
  --set ports.web.exposedPort=30911 \
  --set ports.websecure.exposedPort=30787 \
  --create-namespace

# check install 
kubectl -n traefik get pods

# get our service IP
kubectl -n traefik get svc

# deploy an ingress route that needs TLS for our example app
kubectl apply -f kubernetes/cert-manager/ingress.yaml

# curl over HTTP 
curl http://test.marceldempers.dev
```

### Create an Ingress Let's Encrypt Issuer

We create a `ClusterIssuer` for Ingress that allows us to issue certs in any namespace

```shell
kubectl apply -f kubernetes/cert-manager/issuer-ingress.yaml

# check the issuer
kubectl describe clusterissuer letsencrypt-issuer

```

## Issue Certificate

```shell
kubectl apply -f kubernetes/cert-manager/certificate.yaml

# check the cert issue status
kubectl describe certificate example-app

# you can track and diagnose the ordering request process
kubectl get CertificateRequest
#note:  you can use kubectl describe on the resource too

# TLS created as a secret
kubectl get secrets
NAME                  TYPE                                  DATA   AGE
secret-tls       kubernetes.io/tls                     2      1m

# test TLS 
curl https://test.marceldempers.dev
Hello World!

# cleanup
kubectl delete certificate example-app
kubectl delete clusterissuer letsencrypt-issuer
kubectl delete secret secret-tls
```

## Option 2: Using a Gateway API 

To use this option, we will need a Gateway API enabled cluster. This means you need the Gateway API CRD's installed. </br>

See our [Introduction to Gateway API guide](../gateway-api/README.md)
At the bottom of the guide, there is a matrix of many other Gateway APIs to choose from. </br>

<i> Note: `cert-manager` needs to have `config.enableGatewayAPI=true` enabled.  </i>

### Create a Gateway Class 

```shell
 kubectl apply -f kubernetes/gateway-api/traefik/01-gatewayclass.yaml 
```

### Create a Gateway 

We can use our Gateway from our [Traefik Gateway API Guide](../gateway-api/traefik/README.md)

```shell 
kubectl apply -f kubernetes/gateway-api/traefik/02-gateway.yaml

```

### Create a Gateway API Let's Encrypt Issuer

We create a `ClusterIssuer` for Ingress that allows us to issue certs in any namespace

```shell
kubectl apply -f kubernetes/cert-manager/issuer-gatewayapi.yaml

# check the issuer
kubectl describe clusterissuer letsencrypt-issuer

```

## Issue Certificate

```shell
kubectl apply -f kubernetes/cert-manager/certificate.yaml

# check the cert issue status
kubectl describe certificate example-app

# you can track and diagnose the ordering request process
kubectl get CertificateRequest
#note:  you can use kubectl describe on the resource too

# TLS created as a secret
kubectl get secrets
NAME                  TYPE                                  DATA   AGE
secret-tls       kubernetes.io/tls                     2      1m

# test TLS 
kubectl apply -f kubernetes/cert-manager/httproute.yaml
curl https://test.marceldempers.dev
Hello World!

# cleanup
kubectl delete certificate example-app
kubectl delete clusterissuer letsencrypt-issuer
kubectl delete secret secret-tls
```