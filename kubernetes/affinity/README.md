# Kubernetes Concept: Affinity \ Anti-Affinity 

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
cd kubernetes/affinity
kind create cluster --name demo --image kindest/node:v1.28.0 --config kind.yaml
```

Test the cluster:
```
kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
demo-control-plane   Ready    control-plane   59s   v1.28.0
demo-worker          Ready    <none>          36s   v1.28.0
demo-worker2         Ready    <none>          35s   v1.28.0
demo-worker3         Ready    <none>          35s   v1.28.0

```

## Node Affinity 

[Node Affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity) is similar to `nodeSelector` however you can define more complex expressions. "Like my pods must run on SSD nodes or preffer SSD nodes"

For example: 
* Node selector is a hard and fast rule meaning a pod will not be scheduled if the selection is not satisfied
* For example, when using `os` selector as `linux` , a pod can only be scheduled if there is a node available where `os` label is `linux` 

Node Affinity allows an expression.

```
kubectl apply -f node-affinity.yaml
```

We can see our pods are prefering SSD and are always going to `us-east`

```
kubectl get pods -owide 

#introduce more pods
kubectl scale deploy app-disk --replicas 10

#observe all pods on demo-worker
```

If there is some trouble with our `ssd` disk, `kubectl taint nodes demo-worker type=ssd:NoSchedule`, we can see pods going to the non-ssd disk nodes in `us-east` </br>

This is because our pods prefer SSD, however there is no SSD available, so would still go to non-SSD nodes as long as there are nodes available in `us-east` </br>

If something goes wrong in our last `us-east` node: `kubectl taint nodes demo-worker3 type=ssd:NoSchedule` and we roll out more pods `kubectl scale deploy app-disk --replicas 20`,
notice that our new pods are now in `Pending` status because no nodes satisfy our node affinity rules </br>


Fix our nodes.
```
kubectl taint nodes demo-worker type=ssd:NoSchedule-
kubectl taint nodes demo-worker3 type=ssd:NoSchedule-
```
Scale back down to 0
```
kubectl scale deploy app-disk --replicas 0
kubectl scale deploy app-disk --replicas 1

# pod should go back to demo-worker , node 1
kubectl get pods -owide
```

## Pod Affinity 

Now [Pod Affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity) is an expression to allow us to state that pods should gravitate towards other pods

```
kubectl apply -f pod-affinity.yaml

# observe where pods get deployed
kubectl get pods -owide

kubectl scale deploy app-disk --replicas 3
kubectl scale deploy web-disk --replicas 3
```

## Pod Anti-Affinity

Let's say we observe our `app-disk` application disk usage is quite intense, and we would like to prevent `app-disk` pods from running together. </br>
This is where anti-affinity comes in:

```
podAntiAffinity:
  requiredDuringSchedulingIgnoredDuringExecution:
  - labelSelector:
      matchExpressions:
      - key: app
        operator: In
        values:
        - app-disk
    topologyKey: "kubernetes.io/hostname"
```

After applying the above, we can roll it out and observe scheduling:

```
kubectl scale deploy app-disk --replicas 0
kubectl scale deploy web-disk --replicas 0
kubectl apply -f node-affinity.yaml
kubectl get pods -owide

kubectl scale deploy app-disk --replicas 2 #notice pending pods when scaling to 3
kubectl get pods -owide
kubectl scale deploy web-disk --replicas 2
kubectl get pods -owide

```

