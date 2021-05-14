# Introduction to Linkerd : Distributed Tracing with Jaeger

Get the jaeger extension from the linkerd CLI. See [Official Docs](https://linkerd.io/2.10/tasks/distributed-tracing/)

```
linkerd jaeger install > ./kubernetes/servicemesh/linkerd/tracing/manifests/linkerd-jaeger-21.4.3.yaml
```

Deploy the manifests:

```
kubectl apply -f ./kubernetes/servicemesh/linkerd/tracing/manifests/linkerd-jaeger-21.4.3.yaml
```

See components

```
watch kubectl -n linkerd-jaeger get pods

#do a check
linkerd jaeger check
```


kubectl -n default set env --all deploy OC_AGENT_HOST=collector.linkerd-jaeger:55678

# Jaeger Dashboard

See service mesh traces in the dashboard

```
kubectl -n linkerd-jaeger port-forward svc/jaeger 16686
```

Deploy the Jaeger All-in-One image for demo purposes

 kubectl apply -f .\kubernetes\servicemesh\linkerd\tracing\jaeger-all-in-one.yaml

kubectl port-forward svc/jaeger-query 16686:80

 Enable tracing :

 linkerd upgrade config --addon-config kubernetes/servicemesh/linkerd/tracing/config.yaml | kubectl apply -f -
