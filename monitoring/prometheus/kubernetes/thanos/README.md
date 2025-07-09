# Thanos in Kubernetes

This walkthrough further demonstrates the features of Thanos, by using the Prometheus Operator. </br>
To follow this walkthrough , you will need [Kubernetes Monitoring](../1.33/README.md) in your cluster as well as an understanding of Prometheus Sharding.

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

## Deploy Prometheus Instances:


We will need a service account with RBAC permissions for our new Prometheus instances to access service monitors and allow scraping of service endpoints
In this guide I will use a single service account in two Prometheus instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/serviceaccount.yaml
```

To showcase Thanos, we want to have more than one instance of `Prometheus` running, so that we can demonstrate the sidecar functionality & the global query view that Thanos provides. </br>

Apply Prometheus instances. </br>
We can either apply both manual shards, or use a single defined automated shard which will split into two `StatefulSet` objects. In this guide I will keep it simple and use two `Prometheus` instances. 

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/prometheus-00.yaml
kubectl apply -f monitoring/prometheus/kubernetes/thanos/prometheus-00.yaml

```

Apply Service Monitors 

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/servicemonitors.yaml
```

Each service monitor has an extra label which called `prometheus-shard` which indicates which Prometheus instance should scrape it. The manual shards will use this label to manually select services to scrape. 
Two service monitors will be scraped by instance `prometheus-00` and the other two will be scraped by `prometheus-01`.

For automated sharding, the automated shard will select `ServiceMonitor`'s based on label `monitoring` and `Prometheus` will automatically split each of the selected `ServiceMonitor` items into different `StatefulSets` . </br>

We can now see our Prometheus instances in the `default` namespace:

```
NAME                                  READY   STATUS    RESTARTS   AGE
dotnet-application-74dbc8b5d9-fq2p5   1/1     Running   0          24h
go-application-65bbc698f-jxnls        1/1     Running   0          24h
nodejs-application-c47c5f4c8-cj7jl    1/1     Running   0          24h
prometheus-prometheus-0               2/2     Running   0          17h
prometheus-prometheus-00-0            2/2     Running   0          23h
prometheus-prometheus-01-0            2/2     Running   0          23h
prometheus-prometheus-shard-1-0       2/2     Running   0          17h
python-application-759b44fff7-s276f   1/1     Running   0          24h
```

Checkout each of the automated Prometheus shards

```
kubectl port-forward prometheus-prometheus-0 9090
kubectl port-forward prometheus-prometheus-shard-1-0 9091:9090
```

Now that we have applications to monitor and we have `Prometheus` instances scraping these applications in different instances, we can finally implement Thanos into our solution


## Create an S3 storage for Thanos

Thanos needs an S3 compatible storage for storing its data and long term retention.

In this guide I have created a very basic simple S3 storage using [Minio](https://github.com/minio/minio). Please note my example is not a highly available minio instance and not production ready either.

Deploy our test Minio instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/minio.yaml
```

This will create an S3 storage for testing purpose and a `Job` that will create a bucket for our Thanos data. </br>
