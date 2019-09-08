
VIDEO : https://youtu.be/feLpGydQVio

## Configs

```
kubectl config 

#${HOME}/.kube/config
#kubectl config --kubeconfig="C:\someotherfolder\config"
#$KUBECONFIG

```

### contexts

```
#get the current context
kubectl config current-context

#get and set contexts
kubectl config get-contexts
kubectl config use-context

```

## GET commands
```
kubectl get <resource>

#examples
kubectl get pods
kubectl get deployments
kubectl get services
kubectl get configmaps
kubectl get secrets
kubectl get ingress

```

## Namespaces

```
kubectl get namespaces
kubectl create namespace test
kubectl get pods -n test

```

## Describe command

Used to troubleshoot states and statuses of objects

```
kubectl describe <resource> <name>
```

## Version

```
kubectl version
```