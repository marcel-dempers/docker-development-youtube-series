# Flux Getting Started Guide

# 1 - Kubernetes

For this tutorial, I use Kuberentes 1.17
To get 1.17 for Linux\Windows, just use `kind`

```
#Windows
kind create cluster --name flux --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62

#Linux
kind create cluster --name flux --kubeconfig ~/.kube/kind-flux --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62
```

You can use Kubernetes from Docker for Windows\Mac too! :)

# 2 - Flux CLI Container

Let's run a docker container that can access our Kind cluster and run all the dependencies
for Flux. 
If you installed Flux CTL & KubeCTL on your machine, go to Part 4
```
# Note: make sure we mount the correct network, check the network used (should be bridge)
#grab the ip address
docker inspect flux-control-plane

docker run -it --rm --net bridge -v ${home}/.kube/:/root/.kube/ ubuntu:19.10 bash

apt-get update && apt-get install -y nano curl
# edit kubeconfig, set api address so we can access it from our container

nano ~/.kube/config
# set it to https://flux-control-plane:6443

```

# 3 - Get kubectl and Flux CTL

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

# confirm we can access kind cluster:
kubectl get nodes

curl -o fluxctl -L https://github.com/fluxcd/flux/releases/download/1.19.0/fluxctl_linux_amd64
chmod +x ./fluxctl
mv ./fluxctl /usr/local/bin/fluxctl
```

# 4 - Installing Flux

```
kubectl create ns flux

export GHUSER="marcel-dempers"
fluxctl install \
--git-user=${GHUSER} \
--git-email=${GHUSER}@users.noreply.github.com \
--git-url=git@github.com:${GHUSER}/docker-development-youtube-series \
--git-path=kubernetes/configmaps,kubernetes/secrets,kubernetes/deployments \
--namespace=flux | kubectl apply -f -

kubectl -n flux rollout status deployment/flux

export FLUX_FORWARD_NAMESPACE=flux
fluxctl list-workloads
fluxctl identity


https://github.com/marcel-dempers/docker-development-youtube-series/settings/keys/new

fluxctl sync
```