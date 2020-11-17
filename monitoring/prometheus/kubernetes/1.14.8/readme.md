# Kubernetes 1.14.8 Monitoring Guide

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
```
kind create cluster --name prometheus --image kindest/node:v1.14.9
```

```
kubectl create ns monitoring

# Create the operator to instanciate all CRDs
# Note: You will see error: no matches for kind "ServiceMonitor" ...
#       Wait till the operator is running, then rerun the command
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.14.8/prometheus-operator/

# Deploy monitoring components
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.14.8/node-exporter/
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.14.8/kube-state-metrics/
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.14.8/alertmanager

# Deploy prometheus instance and all the service monitors for targets
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.14.8/prometheus-cluster-monitoring/

# Dashboarding
kubectl -n monitoring create -f ./monitoring/prometheus/kubernetes/1.14.8/grafana/

```

# Sources

The source code for monitoring Kubernetes 1.14 comes from the [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus/tree/v0.3.0/manifests) v0.3.0 tree
