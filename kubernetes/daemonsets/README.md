# Kubernetes Daemonsets

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) </br>
Because a Daemonset is all about running pods on every node, lets create a 3 node cluster:

```
cd kubernetes/daemonsets
kind create cluster --name daemonsets --image kindest/node:v1.20.2 --config kind.yaml
```

Test our cluster:

```
kubectl get nodes
NAME                       STATUS     ROLES                  AGE   VERSION
daemonsets-control-plane   Ready      control-plane,master   65s   v1.20.2
daemonsets-worker          Ready      <none>                 31s   v1.20.2
daemonsets-worker2         Ready      <none>                 31s   v1.20.2
daemonsets-worker3         NotReady   <none>                 31s   v1.20.2
```

# Introduction

Kubernetes provide [documentation](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) for what a Daemonset is with examples.


## Basic Daemonset

Let's deploy a daemonset that runs a pod on each node and collects the name of the node

```
kubectl apply -f daemonset.yaml

kubectl get pods -o wide
NAME                      READY   STATUS    RESTARTS   AGE     IP           NODE                       NOMINATED NODE   READINESS GATES
example-daemonset-8lcr5   1/1     Running   0          3m21s   10.244.2.4   daemonsets-worker          <none>           <none>
example-daemonset-9jhgx   1/1     Running   0          81s     10.244.3.4   daemonsets-worker2         <none>           <none>
example-daemonset-lvvsd   1/1     Running   0          2m41s   10.244.1.4   daemonsets-worker3         <none>           <none>
example-daemonset-xxcv9   1/1     Running   0          119s    10.244.0.7   daemonsets-control-plane   <none>           <none>

```

We can see the logs of any pod

```
kubectl logs <example-daemonset-xxxx>
```

Cleanup:

```
kubectl delete ds example-daemonset
```

## Basic Daemonset: Exposing HTTP

Let's deploy a daemonset that runs a pod on each node and exposes an HTTP endpoint on each node. </br>
In this demo we'll use a simple NGINX on port 80

```
kubectl apply -f daemonset-communication.yaml
kubectl get pods
```

## Communicating with Daemonset Pods 

https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/#communicating-with-daemon-pods

Let's deploy a pod that can talk to our daemonset

```
kubectl apply -f pod.yaml

kubectl exec -it pod -- bash
```

### Service Type: ClusterIP or LoadBalancer

Let's deploy a service of type ClusterIP

```
kubectl apply -f ./services/clusterip-service.yaml

while true; do curl http://daemonset-svc-clusterip; sleep 1s; done
Hello from daemonsets-worker2
Hello from daemonsets-control-plane
Hello from daemonsets-worker2
Hello from daemonsets-worker3
Hello from daemonsets-worker
Hello from daemonsets-worker
```

### Node IP and Node Port

We can add the `nodePort` field to the pods port section to expose a port on the node. </br>

Let's expose the node port in the pod spec:

```
ports:
- containerPort: 80
  hostPort: 80
  name: "http"
```
This means we can contact the daemonset pod using the Node IP and port:

```
# get the node ips
kubectl get nodes -owide

NAME                       STATUS   ROLES                  AGE    VERSION   INTERNAL-IP   EXTERNAL-IP       KERNEL-VERSION                CONTAINER-RUNTIME
daemonsets-control-plane   Ready    control-plane,master   112m   v1.20.2   172.18.0.4    <none>
daemonsets-worker          Ready    <none>                 111m   v1.20.2   172.18.0.3    <none>
daemonsets-worker2         Ready    <none>                 111m   v1.20.2   172.18.0.2    <none>
daemonsets-worker3         Ready    <none>                 111m   v1.20.2   172.18.0.6    <none>

#example:

bash-5.1# curl http://172.18.0.4:80
Hello from daemonsets-control-plane
bash-5.1# curl http://172.18.0.2:80
Hello from daemonsets-worker2
```

### Service: Headless service

Let's deploy a headless service where `clusterIP: None`

```
kubectl apply -f ./services/headless-service.yaml
```

There are a few ways to discover our pods:
1) Discover the [DNS](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#services) records via the "Headless" services


```
apk add --no-cache bind-tools
dig daemonset-svc-headless.default.svc.cluster.local
```

Notice it resolves to multiple DNS records for each pod:

```

;; ANSWER SECTION:
daemonset-svc-headless.default.svc.cluster.local. 30 IN A 10.244.0.5
daemonset-svc-headless.default.svc.cluster.local. 30 IN A 10.244.2.2
daemonset-svc-headless.default.svc.cluster.local. 30 IN A 10.244.3.2
daemonset-svc-headless.default.svc.cluster.local. 30 IN A 10.244.1.2
```

2) Discover pods endpoints by retrieving the endpoints for the headless service

```
kubectl describe endpoints daemonset-svc-headless
```

Example:
```
Addresses:          10.244.0.5,10.244.1.2,10.244.2.2,10.244.3.2

```

Get A records for each pod by using the following format: </br>

`<pod-ip-address>.<my-namespace>.pod.<cluster-domain.example>.` </br>

```
#examples:

10-244-0-5.default.pod.cluster.local
10-244-1-2.default.pod.cluster.local
10-244-2-2.default.pod.cluster.local
10-244-3-2.default.pod.cluster.local
```

Communicate with the pods over DNS:

```
curl http://10-244-0-5.default.pod.cluster.local
Hello from daemonsets-control-plane
```

# Real world Examples:

## Monitoring Nodes: Node-Exporter Daemonset

<br/>
We clone the official kube-prometheus repo to get monitoring manifests for Kubernetes.

```
git clone https://github.com/prometheus-operator/kube-prometheus.git
```

Check the compatibility matrix [here](https://github.com/prometheus-operator/kube-prometheus/tree/v0.8.0#kubernetes-compatibility-matrix)

For this demo, we will use the compatible version tag 0.8

```
git checkout v0.8.0
```

Deploy Prometheus Operator and CRDs
```
cd .\manifests\
kubectl create -f .\setup\
```

Deploy remaining resources including node exporter daemonset

```
kubectl create -f .

# wait for pods to be up
kubectl get pods -n monitoring

#access prometheus in the browser
kubectl -n monitoring port-forward svc/prometheus-k8s 9090

```
See the Daemonset communications on the Prometheus [targets](http://localhost:9090/targets) page

Checkout my [monitoring guide for kubernetes](../../monitoring/prometheus/kubernetes/README.md) for more in depth info

## Monitoring: Logging via Fluentd

Take a look at my monitoring guide for [Fluentd](../../monitoring/logging/fluentd/kubernetes/README.md)
