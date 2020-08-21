# Horizontal Pod Autoscaling

Scales the number of pods in a deployment based off  metrics.
Kubernetes [documentation](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

## Understanding Resources

In this example, I'll be focusing on CPU for scaling. <br/>
We need to ensure we have an understanding of the compute resources we have. <br/>
1) How many cores do we have <br/>
2) How many cores do our application use <br



# scale and keep checking `kubectl top`
# every time we add a pod, CPU load per pod should drop dramatically.
# roughly 8 pods will have each pod use +- 400m

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
