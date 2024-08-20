# Introduction to Service Monitors

<a href="https://youtu.be/_NtRkBipepg" title="k8s-servicemonitors"><img src="https://i.ytimg.com/vi/_NtRkBipepg/hqdefault.jpg" width="20%" alt="k8s-servicemonitors" /></a> 

In order to understand service monitors, we will need to understand how to monitor 
kubernetes environment. </br>
You will need a base understanding of Kubernetes and have a basic understanding of the `kube-prometheus` monitoring stack. </br>

Checkout the video [How to monitor Kubernetes in 2022](https://youtu.be/YDtuwlNTzRc): 

<a href="https://youtu.be/YDtuwlNTzRc" title="Monitoring Kubernetes"><img src="https://i.ytimg.com/vi/YDtuwlNTzRc/hqdefault.jpg" width="50%" alt="Monitoring Kubernetes" /></a>


## Create a kubernetes cluster

```
# create cluster
kind create cluster --name monitoring --image kindest/node:v1.23.5

# see cluster up and running
kubectl get nodes
NAME                  STATUS   ROLES                  AGE     VERSION
monitoring-control-plane   Ready    control-plane,master   2m12s   v1.23.5
```

## Deploy kube-prometheus

Installation: 

```
kubectl create -f ./monitoring/prometheus/kubernetes/1.23/manifests/setup/
kubectl create -f ./monitoring/prometheus/kubernetes/1.23/manifests/
```

Check the install:

```
kubectl -n monitoring get pods
```

After a few minutes, everything should be up and running:

```
kubectl -n monitoring get pods
NAME                                   READY   STATUS    RESTARTS   AGE
alertmanager-main-0                    2/2     Running   0          3m10s
alertmanager-main-1                    2/2     Running   0          3m10s
alertmanager-main-2                    2/2     Running   0          3m10s
blackbox-exporter-6b79c4588b-t4czf     3/3     Running   0          4m7s
grafana-7fd69887fb-zm2d2               1/1     Running   0          4m7s
kube-state-metrics-55f67795cd-f7frb    3/3     Running   0          4m6s
node-exporter-xjdtn                    2/2     Running   0          4m6s
prometheus-adapter-85664b6b74-bvmnj    1/1     Running   0          4m6s
prometheus-adapter-85664b6b74-mcgbz    1/1     Running   0          4m6s
prometheus-k8s-0                       2/2     Running   0          3m9s
prometheus-k8s-1                       2/2     Running   0          3m9s
prometheus-operator-6dc9f66cb7-z98nj   2/2     Running   0          4m6s
```

## View dashboards

```
kubectl -n monitoring port-forward svc/grafana 3000
```

Then access Grafana on [localhost:3000](http://localhost:3000)

## Access Prometheus 

```
kubectl -n monitoring port-forward svc/prometheus-operated 9090
```

Then access Prometheus on [localhost:9090](http://localhost:9090).

## Create our own Prometheus 


```
kubectl apply -n monitoring -f ./kubernetes/servicemonitors/prometheus.yaml

```

View our prometheus `prometheus-applications-0` instance:

```
kubectl -n monitoring get pods
```

Checkout our prometheus UI

```
kubectl -n monitoring port-forward prometheus-applications-0 9090
```

## Deploy a service monitor for example app

```
kubectl -n default apply -f ./kubernetes/servicemonitors/servicemonitor.yaml
```

After applying the service monitor, if Prometheus is correctly selecting it, we should see the item appear under the [Service Discovery](http://localhost:9090/service-discovery) page in Prometheus. </br>
Double check with with `port-forward` before proceeding. </br>
If it does not appear, that means your Prometheus instance is not selecting the service monitor accordingly. Either a label mismatch on the namespace or the service monitor. </br>

## Deploy our example app

```
kubectl -n default apply -f ./kubernetes/servicemonitors/example-app/
```

Now we should see a target in the Prometheus [Targets](http://localhost:9090/targets) page. </br>
