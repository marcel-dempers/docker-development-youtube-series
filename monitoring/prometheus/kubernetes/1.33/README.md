# Kubernetes 1.33 Monitoring Guide

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name monitoring --image kindest/node:v1.33.1 --config monitoring/prometheus/kubernetes/1.33/kind.yaml
```

Test our cluster to see all nodes are healthy and ready:

```
kubectl get nodes
NAME                       STATUS   ROLES           AGE    VERSION
monitoring-control-plane   Ready    control-plane   110s   v1.33.1
monitoring-worker          Ready    <none>          98s    v1.33.1
monitoring-worker2         Ready    <none>          98s    v1.33.1
monitoring-worker3         Ready    <none>          98s    v1.33.1
```

# Kube Prometheus

The best method for monitoring, is to use the community project called `kube-prometheus`
which you can find [here](https://github.com/prometheus-operator/kube-prometheus)

This project focusses on compiling manifests for a complete kubernetes monitoring solution, so with it comes the Prometheus Operator, Grafana, Prometheus instances with Service Monitors that scrape all the internal Kubernetes metrics so we can monitor pods, workloads, nodes and applications.

# Prometheus Community Helm Charts

To deploy this, the best method is to use the Prometheus community helm charts which you can find [here](https://github.com/prometheus-community/helm-charts)

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm search repo prometheus-community --versions
```

Check the possible configuration options:

```
helm show values prometheus-community/kube-prometheus-stack > prometheus-values.yaml
```

## Configation

We can take anything from the values file and create our own [values.yaml](./values.yaml) file to configure anything we want.

## Installation

Firstly let's list the available chart versions 

```
helm search repo prometheus-community --versions > versions.log
```

Let's proceed with installation:
```
CHART_VERSION=75.4.0
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --version ${CHART_VERSION} \
  --namespace monitoring \
  --create-namespace \
  --values monitoring/prometheus/kubernetes/1.33/values.yaml
```

## Prometheus Operator


# Check Monitoring

```
kubectl -n monitoring get pods
NAME                                            READY   STATUS    RESTARTS   AGE
grafana-7d484fc668-x4bf9                        2/3     Running   0          58s
kube-state-metrics-7dd4c79774-cdk7k             1/1     Running   0          58s
node-exporter-brkpr                             1/1     Running   0          58s
node-exporter-k9tpq                             1/1     Running   0          58s
node-exporter-m9bfj                             1/1     Running   0          58s
node-exporter-w7ct4                             1/1     Running   0          58s
prometheus-kube-prometheus-stack-prometheus-0   1/2     Running   0          45s
prometheus-operator-c95dfdfb6-jrjwr             1/1     Running   0          58s
```

To see how Prometheus is configured on what to scrape , we list service monitors

```
kubectl -n monitoring get servicemonitors
NAME                                            AGE
grafana                                         92s
kube-prometheus-stack-apiserver                 92s
kube-prometheus-stack-coredns                   92s
kube-prometheus-stack-kube-controller-manager   92s
kube-prometheus-stack-kube-etcd                 92s
kube-prometheus-stack-kube-proxy                92s
kube-prometheus-stack-kube-scheduler            92s
kube-prometheus-stack-kubelet                   92s
kube-prometheus-stack-prometheus                92s
kube-state-metrics                              92s
node-exporter                                   92s
prometheus-operator                             92s
```

Label selectors are used to map service monitor to kubernetes services. </br>
That is how Prometheus is configured on what to scrape.

As an example you can describe a service monitor:

```
kubectl -n monitoring describe servicemonitor node-exporter
```

# View Dashboards 

You can access the dashboards by using `port-forward` to access Grafana.
It does not have a public endpoint for security reasons

```
kubectl -n monitoring port-forward svc/grafana 3000:80
```

Then access Grafana on [localhost:3000](http://localhost:3000/)


## Check Prometheus 

Similar to checking Grafana, we can also check Prometheus:

```
kubectl -n monitoring port-forward svc/prometheus-operated 9090
```





