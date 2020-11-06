
# The Basics
```

kubectl apply -f kubernetes/kustomize/application/namespace.yaml
kubectl apply -f kubernetes/kustomize/application/configmap.yaml
kubectl apply -f kubernetes/kustomize/application/deployment.yaml
kubectl apply -f kubernetes/kustomize/application/service.yaml

# OR 

kubectl apply -f kubernetes/kustomize/application/

kubectl delete ns example

```

# Kustomize

## Build
```
kubectl kustomize .\kubernetes\kustomize\ | kubectl apply -f -
# OR
kubectl apply -k .\kubernetes\kustomize\

kubectl delete ns example
```

## Overlays

```
kubectl kustomize .\kubernetes\kustomize\environments\production | kubectl apply -f -
# OR
kubectl apply -k .\kubernetes\kustomize\environments\production

kubectl delete ns example
```



