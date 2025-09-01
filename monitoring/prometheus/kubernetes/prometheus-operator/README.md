# Prometheus Operator

This walkthrough demonstrates the usage of the Prometheus Operator. </br>
To follow this walkthrough , you will need [Kubernetes Monitoring](../1.33/README.md) in your cluster.

## Deploy applications to monitor

You may need to build the applications first using docker

```
docker compose -f monitoring/prometheus/docker-compose.yaml build 
```

Load the images into our cluster created in the monitoring guide:

```
kind load docker-image docker.io/library/go-application:latest --name monitoring
kind load docker-image docker.io/library/dotnet-application:latest --name monitoring
kind load docker-image docker.io/library/python-application:latest --name monitoring
kind load docker-image docker.io/library/nodejs-application:latest --name monitoring
```

Deploy our microservices: 

```
kubectl apply -f monitoring/prometheus/go-application/deployment.yaml
kubectl apply -f monitoring/prometheus/dotnet-application/deployment.yaml
kubectl apply -f monitoring/prometheus/python-application/deployment.yaml
kubectl apply -f monitoring/prometheus/nodejs-application/deployment.yaml
```

## Deploy a Prometheus Instance:

We will need a service account with RBAC permissions for our new Prometheus instances to access service monitors and allow scraping of service endpoints

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/serviceaccount.yaml
```

Apply Prometheus instance

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/prometheus.yaml
```

Apply Service Monitors 

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/servicemonitors.yaml
```

We can now see our Prometheus instance in the `default` namespace:

```
kubectl get pods
NAME                                   READY   STATUS    RESTARTS   AGE
dotnet-application-74dbc8b5d9-gmr7h    1/1     Running   0          14m
go-application-65bbc698f-fjmqr         1/1     Running   0          14m
nodejs-application-c47c5f4c8-b9hls     1/1     Running   0          14m
prometheus-prometheus-applications-0   2/2     Running   0          33m
python-application-759b44fff7-9q7ws    1/1     Running   0          14m
```

Checkout the Prometheus instance 

```
kubectl port-forward prometheus-prometheus-applications-0 9090
```

Checkout Grafana

```
kubectl -n monitoring port-forward svc/grafana 3000:80
```

Then access Grafana on [localhost:3000](http://localhost:3000/)



