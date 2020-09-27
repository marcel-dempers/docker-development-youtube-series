# Introduction to Linkerd

## Kubernetes

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name linkerd --image kindest/node:v1.18.4
```

## Get a container to work in

```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh

# install curl & kubectl
apk add --no-cache curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

#test cluster access:
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
linkerd-control-plane   Ready    master   26m   v1.18.4

```

## Linkerd CLI

Lets download the `linkerd` command line tool <br/>
I grabbed the `edge-20.9.4` release using `curl`

You can go to the [releases](https://github.com/linkerd/linkerd2/releases/tag/edge-20.9.4) page to get it

```
curl -L -o linkerd https://github.com/linkerd/linkerd2/releases/download/edge-20.9.4/linkerd2-cli-edge-20.9.4-linux-amd64 
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
linkerd install > ./kubernetes/servicemesh/linkerd/manifest/linkerd-edge-20.9.4.yaml
```

## Install Linkerd

```
kubectl apply -f ./kubernetes/servicemesh/linkerd/manifest/linkerd-edge-20.9.4.yaml
```

Let's wait until all components are running

```
kubectl -n linkerd get deploy
kubectl -n linkerd get svc
```

# Grafana

```
kubectl -n linkerd port-forward svc/linkerd-grafana 3000

```

## Deploy example microservices (Video catalogue)

```
kubectl apply -f kubernetes/servicemesh/applications/playlists-api/
kubectl apply -f kubernetes/servicemesh/applications/videos-api/
kubectl apply -f kubernetes/servicemesh/applications/videos-web/
```

```
kubectl port-forward svc/playlists-api 81:80
kubectl port-forward svc/videos-web 80
kubectl -n ingress-nginx port-forward deploy/nginx-ingress-controller 80
```

# Mesh our video catalog services

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


linkerd -n default check --proxy

linkerd -n default stat deploy

```


