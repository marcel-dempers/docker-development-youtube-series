# Vertical Pod Autoscaling

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name vpa --image kindest/node:v1.19.1
```
<hr/>

## Metric Server

<br/>

* For `Cluster Autoscaler` - On cloud-based clusters, Metric server may already be installed. <br/>
* For `HPA` - We're using kind

[Metric Server](https://github.com/kubernetes-sigs/metrics-server) provides container resource metrics for use in autoscaling pipelines <br/>

Because I run K8s `1.19` in `kind`, the Metric Server version i need is `0.3.7` <br/>
We will need to deploy Metric Server [0.3.7](https://github.com/kubernetes-sigs/metrics-server/releases/tag/v0.3.7) <br/>
I used `components.yaml`from the release page link above. <br/>

<b>Important Note</b> : For Demo clusters (like `kind`), you will need to disable TLS <br/>
You can disable TLS by adding the following to the metrics-server container args <br/>

<b>For production, make sure you remove the following :</b> <br/>

```
- --kubelet-insecure-tls
- --kubelet-preferred-address-types="InternalIP"

```

Deployment: <br/>

```
cd kubernetes\autoscaling
kubectl -n kube-system apply -f .\components\metric-server\metricserver-0.3.7.yaml

#test 
kubectl -n kube-system get pods

#note: wait for metrics to populate!
kubectl top nodes

```

## VPA

VPA docs [here]("https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#install-command") <br/>
Let's install the VPA from a container that can access our cluster

```
cd kubernetes/autoscaling/vertical-pod-autoscaling
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host debian:buster bash

# install git
apt-get update && apt-get install -y git curl nano

# install kubectl 
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl


cd /tmp
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler/

./hack/vpa-up.sh

# after few seconds, we can see the VPA components in:

kubectl -n kube-system get pods
```

## Build and deploy example app

```
# build

cd kubernetes\autoscaling\components\application
docker build . -t aimvector/application-cpu:v1.0.0

# push
docker push aimvector/application-cpu:v1.0.0

# deploy 
kubectl apply -f deployment.yaml

# metrics
kubectl top pods

```

## Generate some traffic

Let's deploy a simple traffic generator pod

```
cd kubernetes\autoscaling\components\application
kubectl apply -f .\traffic-generator.yaml

# get a terminal to the traffic-generator
kubectl exec -it traffic-generator sh

# install wrk
apk add --no-cache wrk

# simulate some load
wrk -c 5 -t 5 -d 99999 -H "Connection: Close" http://application-cpu

```

# Deploy an example VPA

```

kubectl apply -f .\vertical-pod-autoscaling\vpa.yaml

kubectl describe vpa application-cpu

```

# Deploy Goldilocks

```
cd /tmp
git clone https://github.com/FairwindsOps/goldilocks.git
cd goldilocks/hack/manifests/

kubectl create namespace goldilocks
kubectl -n goldilocks apply -f ./controller
kubectl -n goldilocks apply -f ./dashboard


kubectl label ns default goldilocks.fairwinds.com/enabled=true
kubectl label ns default goldilocks.fairwinds.com/vpa-update-mode="off"

kubectl -n goldilocks port-forward svc/goldilocks-dashboard 80

```