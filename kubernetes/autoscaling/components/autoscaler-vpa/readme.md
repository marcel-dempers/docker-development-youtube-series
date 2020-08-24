# Vertical Pod Autoscaling

Provides recommendations for CPU and Memory request values.

## Understanding Resources

In this example, I'll be focusing on CPU for scaling. <br/>
We need to ensure we have an understanding of the compute resources we have. <br/>
1) How many cores do we have <br/>
2) How many cores do our application use <br/>
3) Observe our applications usage
4) Use the VPA to recommend resource request values for our application

## Create a cluster

My Node has 6 CPU cores for this demo <br/>

```
kind create cluster --name vpa --image kindest/node:v1.18.4
```


# Deploy Metric Server

[Metric Server](https://github.com/kubernetes-sigs/metrics-server) provides container resource metrics for use in autoscaling pipelines

We will need to deploy Metric Server [0.3.7](https://github.com/kubernetes-sigs/metrics-server/releases/tag/v0.3.7) <br/>
I used `components.yaml`from the release page link above. <br/>

Note: For Demo clusters (like `kind`), you will need to disable TLS <br/>
You can disable TLS by adding the following to the metrics-server container args <br/>

For production, make sure you remove the following : <br/>

```
- --kubelet-insecure-tls
- --kubelet-preferred-address-types="InternalIP"

```

Deploy it:

```
cd kubernetes\autoscaling
kubectl -n kube-system apply -f .\metric-server\metricserver-0.3.7.yaml

#test 
kubectl -n kube-system get pods

#wait for metrics to populate
kubectl top nodes

```

## Example App

We have an app that simulates CPU usage

```
# build

cd kubernetes\autoscaling\application-cpu
docker build . -t aimvector/application-cpu:v1.0.0

# push
docker push aimvector/application-cpu:v1.0.0

# resource requirements
resources:
  requests:
    memory: "50Mi"
    cpu: "500m"
  limits:
    memory: "500Mi"
    cpu: "2000m"

# deploy 
kubectl apply -f deployment.yaml

# metrics
kubectl top pods
```

## Generate some CPU load

```
# Deploy a tester to run traffic from

cd kubernetes\autoscaling
kubectl apply -f .\autoscaler-vpa\tester.yaml

# get a terminal
kubectl exec -it tester sh
# install wrk
apk add --no-cache wrk curl

# simulate some load
wrk -c 5 -t 5 -d 99999 -H "Connection: Close" http://application-cpu

# scale and keep checking `kubectl top`
# every time we add a pod, CPU load per pod should drop dramatically.
# roughly 8 pods will have each pod use +- 400m

kubectl scale deploy/application-cpu --replicas 2
```