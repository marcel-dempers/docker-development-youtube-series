# Introduction to NGINX Ingress Controller

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name nginx-ingress --image kindest/node:v1.23.5
```

See cluster up and running:

```
kubectl get nodes
NAME                  STATUS   ROLES                  AGE     VERSION
nginx-ingress-control-plane   Ready    control-plane,master   2m12s   v1.23.5
```

## Run a container to work in

### run Alpine Linux: 
```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh
```

### install some tools

```
# install curl 
apk add --no-cache curl

# install kubectl 
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

# install helm 

```

### test cluster access:
```
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
nginx-ingress-control-plane   Ready    control-plane,master   3m26s   v1.23.5
```
