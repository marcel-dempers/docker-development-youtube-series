# Cluster Autoscaling

Scales the number of nodes in our cluster based off usage metrics
[Documentation](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)

## Understanding Resources

In this example, I'll be focusing on CPU for scaling. <br/>
We need to ensure we have an understanding of the compute resources we have. <br/>
1) How many cores do we have <br/>
2) How many cores do our application use <br/>

I go into more details about pod resource utilisation in the Horizontal Pod Autoscaler guide.

# We need a Kubernetes cluster with Cluster Autoscaler

```
# azure example

NAME=aks-getting-started
RESOURCEGROUP=aks-getting-started
SERVICE_PRINCIPAL=
SERVICE_PRINCIPAL_SECRET=

az aks create -n $NAME \
--resource-group $RESOURCEGROUP \
--location australiaeast \
--kubernetes-version 1.16.10 \
--nodepool-name default \
--node-count 1 \
--node-vm-size Standard_F4s_v2  \
--node-osdisk-size 250 \
--service-principal $SERVICE_PRINCIPAL \
--client-secret $SERVICE_PRINCIPAL_SECRET \
--output none \
--enable-cluster-autoscaler \
--min-count 1 \
--max-count 5
```

# Deploy Metric Server

[Metric Server](https://github.com/kubernetes-sigs/metrics-server) provides container resource metrics for use in autoscaling pipelines

We will need to deploy Metric Server [0.3.7](https://github.com/kubernetes-sigs/metrics-server/releases/tag/v0.3.7) <br/>
I used `components.yaml`from the release page link above. <br/>

Note: For Demo clusters (like `kind`), you will need to disable TLS <br/>
You can disable TLS by adding the following to the metrics-server container args

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

cd kubernetes/autoscaling
kubectl apply -f ./autoscaler-cluster/tester.yaml

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

## Deploy an autoscaler

```
# scale the deployment back down to 2
kubectl scale deploy/application-cpu --replicas 2

# deploy the autoscaler
kubectl autoscale deploy/application-cpu --cpu-percent=95 --min=1 --max=10

# pods should scale to roughly 7-8 to match criteria

kubectl describe hpa/application-cpu 
kubectl get hpa/application-cpu  -owide
```
