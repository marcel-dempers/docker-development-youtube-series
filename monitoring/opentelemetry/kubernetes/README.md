# Introduction to OpenTelemetry Operator

OpenTelemetry Operator [documentation](https://opentelemetry.io/docs/platforms/kubernetes/operator/)

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name otel --image kindest/node:v1.34.0
```

Test our cluster:

```
kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
otel-control-plane   Ready    control-plane   40s   v1.34.0
```

## Helm charts

### OpenTelemetry Operator chart
```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update

helm search repo open-telemetry --versions
```

We'll install version `0.93.1` at the time of this guide.
I would suggest to make sure the version you pick is compatible with the Kubernetes version you are running. </br> 

```
OTEL_VERSION=0.93.1
```

### Cert-manager chart

We will need cert-manager deployed which Otel uses for local TLS certificate management

```
helm repo add jetstack https://charts.jetstack.io

helm search repo jetstack --versions
```

We'll install version `v1.18.2` of cert-manager at the time of this guide which is compatible with our Kubernetes version. </br>

```
CERTMANAGER_VERSION=v1.18.2
```

### Install helm charts 

Install cert-manager:

```
helm install \
cert-manager jetstack/cert-manager \
--namespace cert-manager \
--create-namespace \
--version $CERTMANAGER_VERSION \
--set crds.enabled=true \
--set startupapicheck.timeout="5m"
```

Install OpenTelemetry Operator:

```
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
--namespace opentelemetry-operator-system \
--create-namespace \
--version $OTEL_VERSION \
--values=monitoring/opentelemetry/kubernetes/values.yaml
```

View our install: 

```
kubectl get pods -n cert-manager
kubectl get pods -n opentelemetry-operator-system
```

## Create a Collector

In this guide I copied a starter collector from the [official documentation](https://opentelemetry.io/docs/platforms/kubernetes/operator/#getting-started) and changed some of the values. 

To split out the guides, we'll have separate collectors for logs, metrics and traces.

Let's deploy our collectors in a central `monitoring` namespace:

```

kubectl create namespace monitoring

kubectl apply -n monitoring -f monitoring/opentelemetry/kubernetes/collector-tracing.yaml
```

View our collector:

```
kubectl -n monitoring get pods
kubectl -n monitoring get svc
```

## Tracing

### Instrumentation 

To start receiving traces, we'll be using the OpenTelemetry [auto instrumentation](https://opentelemetry.io/docs/platforms/kubernetes/operator/automatic/) to inject the instrumentation into our Kubernetes pods. 

To make use of the auto injection for Opentelemetry, we use the `Instrumentation` resources. 

Note that I am creating this in the default namespace next to where my applications are. 

```
kubectl apply -f monitoring/opentelemetry/kubernetes/instrumentation.yaml
```

### Deploy Microservices

Build applications: 

```
docker compose --file monitoring/opentelemetry/docker-compose.yaml build
```

Load images into `kind`:

```
kind load docker-image aimvector/service-mesh:videos-web-1.0.0 --name otel
kind load docker-image aimvector/service-mesh:playlists-api-1.0.0 --name otel
kind load docker-image aimvector/service-mesh:videos-api-1.0.0 --name otel
```

```
kubectl apply -f monitoring/opentelemetry/applications/playlists-api/
kubectl apply -f monitoring/opentelemetry/applications/playlists-db/
kubectl apply -f monitoring/opentelemetry/applications/videos-web/
kubectl apply -f monitoring/opentelemetry/applications/videos-db/
kubectl apply -f monitoring/opentelemetry/applications/videos-api/
```

Generate some traffic with `port-forward` 

```
kubectl port-forward svc/videos-web 80:80
kubectl port-forward svc/playlists-api 81:80
```

### Tracing Data Store

In this guide, we'll use a Tempo database for tracing data

```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm search repo  grafana/tempo

TEMPO_VERSION=1.23.3

helm install tempo grafana/tempo \
    --create-namespace \
    --namespace grafana \
    --version $TEMPO_VERSION \
    --values monitoring/opentelemetry/kubernetes/tempo.yaml
```

### Metrics Data Store

In this guide, we'll use a Prometheus for our metrics data

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm search repo prometheus-community --versions

PROMETHEUS_STACK_VERSION=77.5.0

helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --version ${PROMETHEUS_STACK_VERSION} \
  --namespace prometheus-operator-system \
  --create-namespace \
  --set prometheusOperator.enabled=true \
  --set prometheusOperator.nodeSelector."kubernetes\.io/os"=linux \
  --set prometheusOperator.fullnameOverride="prometheus-operator" \
  --set prometheusOperator.manageCrds=true \
  --set alertmanager.enabled=false \
  --set grafana.enabled=false \
  --set prometheus-node-exporter.enabled=false \
  --set nodeExporter.enabled=false \
  --set kubeStateMetrics.enabled=false \
  --set prometheus.enabled=false

kubectl -n prometheus-operator-system get pods 

# Deploy our dedicated Prometheus for metrics storage

kubectl apply -n monitoring -f monitoring/opentelemetry/kubernetes/prometheus.yaml
```

### Dashboards

In this guide I use a simple Grafana for dashboards

```
helm search repo grafana/grafana

GRAFANA_VERSION=9.4.4

helm install grafana grafana/grafana \
  --namespace grafana \
  --version $GRAFANA_VERSION \
  --values monitoring/opentelemetry/kubernetes/grafana.yaml

```

### Access Grafana 

We can `port-forward` to Grafana

```
kubectl -n grafana port-forward svc/grafana 3000:80
```