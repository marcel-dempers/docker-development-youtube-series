# Introduction to Istio

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name istio --image kindest/node:v1.19.1
```

## Deploy our microservices (Video catalog)

```
# ingress controller
kubectl create ns ingress-nginx
kubectl apply -f kubernetes/servicemesh/applications/ingress-nginx/

# applications
kubectl apply -f kubernetes/servicemesh/applications/playlists-api/
kubectl apply -f kubernetes/servicemesh/applications/playlists-db/
kubectl apply -f kubernetes/servicemesh/applications/videos-api/
kubectl apply -f kubernetes/servicemesh/applications/videos-web/
kubectl apply -f kubernetes/servicemesh/applications/videos-db/
```

## Make sure our applications are running 

```
kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE  
playlists-api-d7f64c9c6-rfhdg   1/1     Running   0          2m19s
playlists-db-67d75dc7f4-p8wk5   1/1     Running   0          2m19s
videos-api-7769dfc56b-fsqsr     1/1     Running   0          2m18s
videos-db-74576d7c7d-5ljdh      1/1     Running   0          2m18s
videos-web-598c76f8f-chhgm      1/1     Running   0          100s 

```

## Make sure our ingress controller is running

```
kubectl -n ingress-nginx get pods
NAME                                        READY   STATUS    RESTARTS   AGE  
nginx-ingress-controller-6fbb446cff-8fwxz   1/1     Running   0          2m38s
nginx-ingress-controller-6fbb446cff-zbw7x   1/1     Running   0          2m38s

```

We'll need a fake DNS name `servicemesh.demo` <br/>
Let's fake one by adding the following entry in our hosts (`C:\Windows\System32\drivers\etc\hosts`) file: <br/>

```
127.0.0.1  servicemesh.demo

```

## Let's access our applications via Ingress 

```
kubectl -n ingress-nginx port-forward deploy/nginx-ingress-controller 80
```

## Access our application in the browser

We should be able to access our site under `http://servicemesh.demo/home/`

<br/>
<hr/>

# Getting Started with Istio

Firstly, I like to do most of my work in containers so everything is reproducable <br/>
and my machine remains clean.

## Get a container to work in
<br/>
Run a small `alpine linux` container where we can install and play with `istio`: <br/>


```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh

# install curl & kubectl
apk add --no-cache curl nano
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
export KUBE_EDITOR="nano"

#test cluster access:
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
istio-control-plane   Ready    master   26m   v1.18.4

```

## Install Istio CLI

```
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.12 TARGET_ARCH=x86_64 sh -

mv istio-1.6.12/bin/istioctl /usr/local/bin/
chmod +x /usr/local/bin/istioctl

mv istio-1.6.12 /tmp/

```

## Pre flight checks

Istio has a great capability to check compatibility with the target cluster <br/>

```
istioctl x precheck

```

## Istio Profiles

https://istio.io/latest/docs/setup/additional-setup/config-profiles/

```
istioctl profile list

istioctl install --set profile=default

kubectl -n istio-system get pods

istioctl verify-install

```

https://istio.io/latest/docs/setup/additional-setup/sidecar-injection/

## Manual sidecar injection
kubectl get deploy playlists-api -o yaml | istioctl kube-inject -f - | kubectl apply -f -


# Mesh our video catalog services

There are 2 ways to mesh:

1) Automated Injection: 

You can set the `istio-injection=enabled` label on a namespace to have the istio side car automatically injected into any pod that gets created in the labelled namespace

This is a more permanent solution:

```
  template:
    metadata:
      labels:
        istio-injection: enabled
```

2) Manual Injection:
This may only be temporary as your CI/CD system may roll out the previous YAML.
You may want to add this command to your CI/CD to keep only certain deployments part of the mesh.

```
kubectl get deploy
NAME            READY   UP-TO-DATE   AVAILABLE   AGE
playlists-api   1/1     1            1           8h 
playlists-db    1/1     1            1           8h 
videos-api      1/1     1            1           8h 
videos-db       1/1     1            1           8h 
videos-web      1/1     1            1           8h 

# You can manually inject istio sidecar to every deployment like this:

kubectl get deploy playlists-api -o yaml | istioctl kube-inject -f - | kubectl apply -f -
kubectl get deploy playlists-db -o yaml | istioctl kube-inject -f - | kubectl apply -f -
kubectl get deploy videos-api -o yaml | istioctl kube-inject -f - | kubectl apply -f -
kubectl get deploy videos-db -o yaml | istioctl kube-inject -f - | kubectl apply -f -
kubectl get deploy videos-web -o yaml | istioctl kube-inject -f - | kubectl apply -f -
kubectl -n ingress-nginx get deploy nginx-ingress-controller  -o yaml | istioctl kube-inject -f - | kubectl apply -f -

```

# TCP \ HTTP traffic

# Observability

## Kiali 

note: rerun for CRDs
```
kubectl apply -f /tmp/istio-1.6.12/samples/addons/kiali.yaml


kubectl -n istio-system port-forward svc/kiali 20001
```