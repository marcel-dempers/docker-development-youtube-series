# Thanos in Kubernetes

This walkthrough further demonstrates the features of Thanos, by using the Prometheus Operator. </br>
To follow this walkthrough , you will need [Kubernetes Monitoring](../1.33/README.md) in your cluster as well as an understanding of Prometheus Sharding.

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

# see all our apps 
kubectl get pods
NAME                                  READY   STATUS    RESTARTS   AGE
dotnet-application-74dbc8b5d9-tjtxk   1/1     Running   0          3s
go-application-65bbc698f-92jdr        1/1     Running   0          2m7s
nodejs-application-c47c5f4c8-b9jhg    1/1     Running   0          2m7s
python-application-759b44fff7-5fn8r   1/1     Running   0          2m7s
```

## Create Service Monitors 

We'll need `ServiceMonitor` objects to point to our applications. In my guide I generally have one `ServiceMonitor` for each application.

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/servicemonitors.yaml
```

Each service monitor has an extra label which called `prometheus-shard` which indicates which Prometheus instance should scrape it. The manual shards will use this label to manually select services to scrape. 
Two service monitors will be scraped by instance `prometheus-00` and the other two will be scraped by `prometheus-01`.

For automated sharding, the automated shard will select `ServiceMonitor`'s based on label `monitoring` and `Prometheus` will automatically split each of the selected `ServiceMonitor` items into different `StatefulSets` . </br>
See my Prometheus Sharding video in Kubernetes for more on that. </br>

## Create an S3 storage for Thanos

Thanos needs an S3 compatible storage for storing its data and long term retention.

In this guide I have created a very basic simple S3 storage using [Minio](https://github.com/minio/minio). Please note my example is not a highly available minio instance and not production ready either.

Deploy our test Minio instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/minio.yaml
```

This will create an S3 storage for testing purpose and a `Job` that will create a bucket for our Thanos data. </br>

## Create Thanos storage secret 

Before we apply Prometheus instances, we will need a secret that Thanos sidecars in each Prometheus instance will use to connect to store data in S3. 

Let's create a secret for minio:

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/thanos-secret.yaml
```

## Deploy Prometheus Instances:

We will need a service account with RBAC permissions for our new Prometheus instances to access service monitors and allow scraping of service endpoints
In this guide I will use a single service account in two Prometheus instance:

```
kubectl apply -f monitoring/prometheus/kubernetes/prometheus-operator/serviceaccount.yaml
```

To showcase Thanos, we want to have more than one instance of `Prometheus` running, so that we can demonstrate the sidecar functionality & the global query view that Thanos provides. </br>

We can either apply both manual shards, or use a single defined automated shard which will split into two `StatefulSet` objects. In this guide I will keep it simple and use two `Prometheus` instances.

To understand Thanos, we'll take a closer look at the `Prometheus` spec in our instances we are going to deploy: 

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/prometheus-00.yaml
kubectl apply -f monitoring/prometheus/kubernetes/thanos/prometheus-01.yaml

```

We can now see our Prometheus instances in the `default` namespace:

```
kubectl get pods
NAME                                  READY   STATUS            RESTARTS   AGE
dotnet-application-74dbc8b5d9-tjtxk   1/1     Running           0          60m
go-application-65bbc698f-92jdr        1/1     Running           0          62m
nodejs-application-c47c5f4c8-b9jhg    1/1     Running           0          62m
prometheus-prometheus-00-0            3/3     Running           0          21s
prometheus-prometheus-01-0            3/3     Running           0          17s
python-application-759b44fff7-5fn8r   1/1     Running           0          62m
```

Checkout each of the `Prometheus` instances and see our 4 applications sharded to 2 instances evenly. </br>

```
kubectl port-forward prometheus-prometheus-00-0  9090
kubectl port-forward prometheus-prometheus-01-0  9091:9090
```

## Deploying the Thanos Query Service

Thanos is now enabled on our `Prometheus` instances, which means the sidecars are shipping metrics data from each instance to our S3. </br>
For us to create a global query service we need to deploy the Thanos Query component which will link up to each Thanos sidecar, using the service created by the Prometheus operator

```
kubectl apply -f monitoring/prometheus/kubernetes/thanos/thanos-query.yaml
```

## Test in Grafana 

We can now verify that we can query Thanos in Grafana by using `port-forward` to access Grafana.

```
kubectl -n monitoring port-forward svc/grafana 3000:80
```

Then access Grafana on [localhost:3000](http://localhost:3000/)