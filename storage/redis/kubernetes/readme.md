# Redis on Kubernetes

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name redis --image kindest/node:v1.18.4
```

## Namespace

```
kubectl create ns redis
```

## Storage Class

```
kubectl get storageclass
NAME                 PROVISIONER             RECLAIMPOLICY   VOLUMEBINDINGMODE      ALLOWVOLUMEEXPANSION   AGE
standard (default)   rancher.io/local-path   Delete          WaitForFirstConsumer   false                  84s
```

## Deployment: Redis nodes

```
cd storage/redis/kubernetes/
kubectl apply -n redis -f ./redis/redis-configmap.yaml
kubectl apply -n redis -f ./redis/redis-statefulset.yaml

kubectl -n redis get pods
kubectl -n redis get pv

kubectl -n redis logs redis-0
kubectl -n redis logs redis-1
kubectl -n redis logs redis-2
```

## Test replication status

```
kubectl -n redis exec -it redis-0 sh
redis-cli 
auth a-very-complex-password-here
info replication
```

## Deployment: Redis Sentinel (3 instances)

```
cd storage/redis/kubernetes/
kubectl apply -n redis -f ./sentinel/sentinel-statefulset.yaml

kubectl -n redis get pods
kubectl -n redis get pv
kubectl -n redis logs sentinel-0
```