# Introduction to Linkerd

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name linkerd --image kindest/node:v1.19.1
```

## Deploy our microservices (Video catalog)

```
# ingress controller
kubectl create ns ingress-nginx
kubectl apply -f kubernetes/servicemesh/applications/ingress-nginx/

# applications
kubectl apply -f kubernetes/servicemesh/applications/playlists-api/
kubectl apply -f kubernetes/servicemesh/applications/playlists-db/
kubectl apply -f kubernetes/servicemesh/applications/videos-web/
kubectl apply -f kubernetes/servicemesh/applications/videos-api/
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

# Getting Started with Linkerd

Firstly, I like to do most of my work in containers so everything is reproducible  <br/>
and my machine remains clean.

## Get a container to work in
<br/>
Run a small `alpine linux` container where we can install and play with `linkerd`: <br/>

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
linkerd-control-plane   Ready    master   26m   v1.19.1

```

## Linkerd CLI

Lets download the `linkerd` command line tool <br/>
I grabbed the `edge-20.10.1` release using `curl`

You can go to the [releases](https://github.com/linkerd/linkerd2/releases/tag/edge-20.10.1) page to get it

```
curl -L -o linkerd https://github.com/linkerd/linkerd2/releases/download/edge-20.10.1/linkerd2-cli-edge-20.10.1-linux-amd64 
chmod +x linkerd && mv ./linkerd /usr/local/bin/

linkerd --help
```

## Pre flight checks

Linkerd has a great capability to check compatibility with the target cluster <br/>

```
linkerd check --pre

```

## Get the YAML

```
linkerd install > ./kubernetes/servicemesh/linkerd/manifest/linkerd-edge-20.10.1.yaml
```

## Install Linkerd

```
kubectl apply -f ./kubernetes/servicemesh/linkerd/manifest/linkerd-edge-20.10.1.yaml
```

Let's wait until all components are running

```
watch kubectl -n linkerd get pods
kubectl -n linkerd get svc
```

## Do a final check

```
linkerd check
```

## The dashboard

Let's access the `linkerd` dashboard via `port-forward`

```
kubectl -n linkerd port-forward svc/linkerd-web 8084
```

# Mesh our video catalog services

There are 2 ways to mesh:

1) We can add an annotation to your deployment to persist the mesh if our YAML is part of a GitOps flow:
This is a more permanent solution:

```
  template:
    metadata:
      annotations:
        linkerd.io/inject: enabled
```

2) Or inject `linkerd` on the fly:
This may only be temporary as your CI/CD system may roll out the previous YAML

```
kubectl get deploy
NAME            READY   UP-TO-DATE   AVAILABLE   AGE
playlists-api   1/1     1            1           8h 
playlists-db    1/1     1            1           8h 
videos-api      1/1     1            1           8h 
videos-db       1/1     1            1           8h 
videos-web      1/1     1            1           8h 

kubectl get deploy playlists-api -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy playlists-db -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-api -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-db -o yaml | linkerd inject - | kubectl apply -f -
kubectl get deploy videos-web -o yaml | linkerd inject - | kubectl apply -f -
kubectl -n ingress-nginx get deploy nginx-ingress-controller  -o yaml | linkerd inject - | kubectl apply -f -

```

# Generate some traffic

Let's run a `curl` loop to generate some traffic to our site </br>
We'll make a call to `/home/` and to simulate the browser making a call to get the playlists, <br/>
we'll make a follow up call to `/api/playlists`

```
While ($true) { curl -UseBasicParsing http://servicemesh.demo/home/;curl -UseBasicParsing http://servicemesh.demo/api/playlists; Start-Sleep -Seconds 1;}

linkerd -n default check --proxy

linkerd -n default stat deploy

```

# Add Faulty behaviour in videos API

```
kubectl edit deploy videos-api

#set environment FLAKY=true
```

# Service Profile 

```
linkerd profile -n default videos-api --tap deploy/videos-api --tap-duration 10s
```

After crafting the `serviceprofile`, we can apply it using `kubectl`

```
 kubectl apply -f kubernetes/servicemesh/linkerd/serviceprofiles/videos-api.yaml
```

We can see that service profile helps us add retry policies in place: <br/>

```
linkerd routes -n default deploy/playlists-api --to svc/videos-api -o wide
linkerd top deploy/videos-api
```

# Mutual TLS

We can validate if mTLS is working 

```
/work # linkerd -n default edges deployment
SRC                  DST             SRC_NS    DST_NS    SECURED       
playlists-api        videos-api      default   default   √
linkerd-prometheus   playlists-api   linkerd   default   √
linkerd-prometheus   playlists-db    linkerd   default   √
linkerd-prometheus   videos-api      linkerd   default   √
linkerd-prometheus   videos-db       linkerd   default   √
linkerd-prometheus   videos-web      linkerd   default   √
linkerd-tap          playlists-api   linkerd   default   √
linkerd-tap          playlists-db    linkerd   default   √
linkerd-tap          videos-api      linkerd   default   √
linkerd-tap          videos-db       linkerd   default   √
linkerd-tap          videos-web      linkerd   default   √

linkerd -n default tap deploy

```