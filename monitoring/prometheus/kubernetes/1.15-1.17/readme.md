# Kubernetes 1.15-1.17 Monitoring Guide

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
```

# Kubernetes 1.16.9
kind create cluster --name prometheus --image kindest/node:v1.16.9

```

```
kubectl create ns monitoring

# Create the operator to instanciate all CRDs
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.15-1.17/prometheus-operator/

# Deploy monitoring components
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.15-1.17/node-exporter/
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.15-1.17/kube-state-metrics/
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.15-1.17/alertmanager

# Deploy prometheus instance and all the service monitors for targets
kubectl -n monitoring apply -f ./monitoring/prometheus/kubernetes/1.15-1.17/prometheus-cluster-monitoring/

# Dashboarding
kubectl -n monitoring create -f ./monitoring/prometheus/kubernetes/1.15-1.17/grafana/

# Check the pods
kubectl -n monitoring get pods

# Note : Metrics can take couple of minutes to ingest!

# Test target connectivity
kubectl -n monitoring port-forward prometheus-k8s-0 9090

# Dashboards 
kubectl -n monitoring port-forward <grafana-pod-name> 3000
```

# Sources

The source code for monitoring Kubernetes 1.15-1.17 comes from the [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus/tree/v0.3.0/manifests) v0.3.0 tree
