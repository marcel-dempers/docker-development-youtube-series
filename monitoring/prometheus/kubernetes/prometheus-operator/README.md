# Prometheus Operator

This walkthrough demonstrates the usage of the Prometheus Operator. </br>
To follow this walkthrough , you will need [Kubernetes Monitoring](../1.33/README.md) in your cluster.

## Deploy applications to monitor

You may need to build the applications first using docker

```
docker compose -f monitoring\prometheus\docker-compose.yaml build 
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


