# Prometheus Sharding in Kubernetes using Prometheus Operator

This walkthrough further demonstrates the features of Thanos, by using the Prometheus Operator. </br>
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

## Sharding Strategies

There are two main types of strategies to achieve sharding, Manual and Automated. </br>

**Manual Sharding**

Manual sharding is more robust, and you are in control. </br>
It involves manually creating a `Prometheus` instance or "shard" and then setting `ServiceMonitors` per instance to scrape.  

**Automated Sharding**

Automated sharding is an automated sharding strategy where Prometheus will use a hasmhmod technique to automatically select `ServiceMonitors` per shard to scrape. 

--- 

### Setup a service account

We will need a service account with RBAC permissions for our new Prometheus instances to access service monitors and allow scraping of service endpoints
In this guide I will use a single service account in two Prometheus instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/serviceaccount.yaml
```

### Manual Sharding 

Apply two Prometheus instances that are manually setup to target their own sets of `ServiceMonitor`'s to scrape:

```
kubectl apply -f monitoring/prometheus/kubernetes/sharding/manual/prometheus-00.yaml
kubectl apply -f monitoring/prometheus/kubernetes/sharding/manual/prometheus-01.yaml
```

Each service monitor has an extra label which called `prometheus-shard` which indicates which Prometheus instance should scrape it. 
Two service monitors will be scraped by instance `prometheus-00` and the other two will be scraped by `prometheus-01`. </br>


Apply Service Monitors : 

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/servicemonitors.yaml
```


We can now see our Prometheus instance in the `default` namespace:

```
NAME                                  READY   STATUS    RESTARTS   AGE
dotnet-application-74dbc8b5d9-fq2p5   1/1     Running   0          46m
go-application-65bbc698f-jxnls        1/1     Running   0          49m
nodejs-application-c47c5f4c8-cj7jl    1/1     Running   0          49m
prometheus-prometheus-00-0            2/2     Running   0          3m15s
prometheus-prometheus-01-0            2/2     Running   0          3m15s
python-application-759b44fff7-s276f   1/1     Running   0          49m
```

Checkout each of the Prometheus instances

```
kubectl port-forward prometheus-prometheus-00-0 9090
kubectl port-forward prometheus-prometheus-01-0 9091:9090
```

### Automated Sharding 

In larger clusters, it can become toil to manually set sharding labels on `ServiceMonitor` resources and manually balancing `Prometheus` instances. 

The Prometheus-Operator supports automated sharding, which follows the technique outlined in our Sharding Introduction video. </br>

The Prometheus-Operator will automate everything we learned (and performed in docker). </br>
It will create a separate `StatefulSet` for each shard and use the `hashmod` technique, typically on the address field of the service endpoints. 

All we need to do is use the `spec.shards` field. </br>
The operator will also set external labels for each replica or shard as `prometheus_replica` which you can view in the Promerheus configuration in the UI. </br>

Let's apply our sharded `Prometheus` instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/sharding/automated/prometheus.yaml
```

Checkout each of the automated Prometheus shards

```
kubectl port-forward prometheus-prometheus-0 9090
kubectl port-forward prometheus-prometheus-shard-1-0 9091:9090
```