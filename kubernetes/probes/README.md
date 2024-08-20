# Introduction to Kubernetes Probes


## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
cd kubernetes/probes
kind create cluster --name demo --image kindest/node:v1.28.0 
```

Test the cluster:
```
kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
demo-control-plane   Ready    control-plane   59s   v1.28.0

```

## Applications

Client app is used to act as a client that sends web requests :

```
kubectl apply -f client.yaml
```

The server app is the app that will receive web requests:

```
kubectl apply -f server.yaml
```

Test making web requests constantly:

```
while true; do curl http://server; sleep 1s; done
```

Bump the server `version` label up and apply to force a new deployment </br>
Notice the client throws an error, so traffic is interupted, not good! </br>

This is because our new pod during deployment is not ready to take traffic!

## Readiness Probes

Let's add a readiness probe that tells Kubernetes when we are ready:

```
readinessProbe:
  httpGet:
    path: /
    port: 5000
  initialDelaySeconds: 3
  periodSeconds: 3
  failureThreshold: 3
```

### Automatic failover with Readiness probes

Let's pretend our application starts hanging and not longer returns responses </br>
This is common with some web servers and may need to be manually restarted 

```
kubectl exec -it podname -- sh -c "rm /data.txt"   
```

Now we will notice our client app starts getting errors. </br>
Few things to notice:

* Our readiness probe detected an issue and removed traffic from the faulty pod.
* We should be running more than one application so we would be highly available

```
kubectl scale deploy server --replicas 2
```

* Notice traffic comes back as its routed to the healthy pod

Fix our old pod:  `kubectl exec -it podname -- sh -c "echo 'ok' > /data.txt"` </br>

* If we do this again with 2 pods, notice we still get an interuption but our app automaticall stabalises after some time
* This is because readinessProbe has `failureThreshold` and some failure will be expected before recovery
* Do not set this `failureThreshold` too low as you may remove traffic frequently. Tune accordingly!

Readiness probes help us automatically remove traffic when there are intermittent network issues </br>

## Liveness Probes

Liveness probe helps us when we cannot automatically recover. </br>
Let's use the same mechanism to create a vaulty pod:

```
kubectl exec -it podname -- sh -c "rm /data.txt" 
```

Our readiness probe has saved us from traffic issues. </br>
But we want the pod to recover automatically, so let's create livenessProbe:

```
livenessProbe:
  httpGet:
    path: /
    port: 5000
  initialDelaySeconds: 3
  periodSeconds: 4
  failureThreshold: 8
```

Scale back up: `kubectl scale deploy server --replicas 2`
Create a vaulty pod: `kubectl exec -it podname -- sh -c "rm /data.txt" `

If we observe we will notice the readinessProbe saves our traffic, and livenessProbe will eventually replace the bad pod </br>

## Startup Probes

The [startup probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes) is for slow starting applications </br>
It's important to understand difference between start up and readiness probes. </br>
In our examples here, readiness probe acts as a startup probe too, since our app is fairly slow starting! </br>
This difference is explained in the video. </br>