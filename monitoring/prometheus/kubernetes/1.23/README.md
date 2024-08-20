# Kubernetes 1.23 Monitoring Guide

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
cd monitoring\prometheus\kubernetes\1.23

kind create cluster --name monitoring --image kindest/node:v1.23.6 --config kind.yaml
```

Test our cluster to see all nodes are healthy and ready:

```
kubectl get nodes
NAME                       STATUS   ROLES                  AGE   VERSION
monitoring-control-plane   Ready    control-plane,master   66s   v1.23.1
monitoring-worker          Ready    <none>                 25s   v1.23.1
monitoring-worker2         Ready    <none>                 38s   v1.23.1
monitoring-worker3         Ready    <none>                 38s   v1.23.1
```

# Kube Prometheus

The best method for monitoring, is to use the community manifests on the `kube-prometheus`
repository [here](https://github.com/prometheus-operator/kube-prometheus)

Now according to the compatibility matrix, we will need `release-0.10` to be compatible with
Kubernetes 1.23. </br>

Let's use docker to grab it! 

```
docker run -it -v ${PWD}:/work -w /work alpine sh
apk add git
```

Shallow clone the release branch into a temporary folder:

```
# clone
git clone --depth 1 https://github.com/prometheus-operator/kube-prometheus.git -b release-0.10 /tmp/

# view the files
ls /tmp/ -l

# we are interested in the "manifests" folder
ls /tmp/manifests -l

# let's grab it by coping it out the container
cp -R /tmp/manifests .
```

## Prometheus Operator

To deploy all these manifests, we will need to setup the prometheus operator and custom resource definitions required.

This is all in the `setup` directory:

```
ls /tmp/manifests/setup -l
```

Now that we have the source code manifests, we can exit our temporary container

```
exit
```

# Setup CRDs

Let's create the CRD's and prometheus operator

```
kubectl create -f ./manifests/setup/
```

# Setup Manifests

Apply rest of manifests

```
kubectl create -f ./manifests/
```

# Check Monitoring

```
kubectl -n monitoring get pods

NAME                                   READY   STATUS    RESTARTS   AGE
alertmanager-main-0                    2/2     Running   0          26m
alertmanager-main-1                    2/2     Running   0          26m
alertmanager-main-2                    2/2     Running   0          26m
blackbox-exporter-6b79c4588b-rvd2n     3/3     Running   0          27m
grafana-7fd69887fb-vzshr               1/1     Running   0          27m
kube-state-metrics-55f67795cd-gmwlk    3/3     Running   0          27m
node-exporter-77d29                    2/2     Running   0          27m
node-exporter-7ndbl                    2/2     Running   0          27m
node-exporter-pgzq7                    2/2     Running   0          27m
node-exporter-vbxrt                    2/2     Running   0          27m
prometheus-adapter-85664b6b74-nhjw8    1/1     Running   0          27m
prometheus-adapter-85664b6b74-t5zfj    1/1     Running   0          27m
prometheus-k8s-0                       2/2     Running   0          26m
prometheus-k8s-1                       2/2     Running   0          26m
prometheus-operator-6dc9f66cb7-8bg77   2/2     Running   0          27m
```

# View Dashboards 

You can access the dashboards by using `port-forward` to access Grafana.
It does not have a public endpoint for security reasons

```
kubectl -n monitoring port-forward svc/grafana 3000
```

Then access Grafana on [localhost:3000](http://localhost:3000/)

# Fix Grafana Datasource

Now for some reason, the Prometheus data source in Grafana does not work out the box.
To fix it, we need to change the service endpoint of the data source. </br>

To do this, edit `manifests/grafana-dashboardDatasources.yaml` and replace the datasource url endpoint with `http://prometheus-operated.monitoring.svc:9090` </br>

We'll need to patch that and restart Grafana

```
kubectl apply -f ./manifests/grafana-dashboardDatasources.yaml
kubectl -n monitoring delete po <grafana-pod>
kubectl -n monitoring port-forward svc/grafana 3000
```

Now our datasource should be healthy.

## Check Prometheus 

Similar to checking Grafana, we can also check Prometheus:

```
kubectl -n monitoring port-forward svc/prometheus-operated 9090
```

## Check Service Monitors 

To see how Prometheus is configured on what to scrape , we list service monitors

```
kubectl -n monitoring get servicemonitors
NAME                      AGE
alertmanager-main         8m58s
blackbox-exporter         8m57s
coredns                   8m57s
grafana                   8m57s
kube-apiserver            8m57s
kube-controller-manager   8m57s
kube-scheduler            8m57s
kube-state-metrics        8m57s
kubelet                   8m57s
node-exporter             8m57s
prometheus-adapter        8m56s
prometheus-k8s            8m56s
prometheus-operator       8m56s

kubectl -n monitoring describe servicemonitor node-exporter
```

Label selectors are used to map service monitor to kubernetes services. </br>

That is how Prometheus is configured on what to scrape.