# Introduction to Cluster API `k8s-clusterapi-intro`

Documentation: <br/>
The [concepts](https://cluster-api.sigs.k8s.io/user/concepts)

## The Management Cluster

The cluster used to manage all other clusters. <br/>
We'll use a `kind` cluster, but you can use any cluster <br/>

Management clusters can create other clusters. These "other" clusters are what we call workload clusters <br/>

For `kind`, to create clusters, it needs access to docker. <br/>
For cloud-based clusters like AKS, GKE, EKS, you would need some kind of service accounts for the management cluster to be able to provision workload clusters <br/>

To create workload clusters, the management cluster needs "providers" configured. <br/>

Create a management cluster:

```bash
kind create cluster --name cluster-api --config kubernetes/cluster-api/introduction/kind.yaml
```

Install clusterctl:

```bash
curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/v1.14.3/clusterctl-linux-amd64 -o clusterctl

chmod +x clusterctl 
mv clusterctl /usr/local/bin/

clusterctl --help
```
The `clusterctl` commands will run against the cluster our `KUBECONFIG` is pointing to, in our case its the `kind` cluster <br/>

Initialise:

```bash
# Enable the experimental Cluster topology feature.
export CLUSTER_TOPOLOGY=true

# Initialize the management cluster
clusterctl init --infrastructure docker

# check the cluster 
kubectl get pods -A
```

## The Workload Cluster

```bash
clusterctl generate cluster capi-quickstart --flavor development \
  --kubernetes-version v1.37.0 \
  --control-plane-machine-count=3 \
  --worker-machine-count=3 \
  > kubernetes/cluster-api/introduction/workload-1.yaml
```

Deploy the cluster:

```bash
kubectl apply -f kubernetes/cluster-api/introduction/workload-1.yaml
```

Check the cluster:

```
kubectl get cluster
kubectl describe cluster
kubectl get kubeadmcontrolplane
```

Access the cluster:

```bash
clusterctl get kubeconfig capi-quickstart > capi-quickstart.kubeconfig 

export KUBECONFIG=./capi-quickstart.kubeconfig
kubectl cluster-info 
kubectl get nodes
```

## Exploring the cluster objects

- `ClusterClass`, `Cluster`, `KubeadmControlPlaneTemplate`, `KubeadmConfigTemplate`
- `MachineDeployment`, `MachineSet`, `Machine`
- `DevClusterTemplate`, `DevMachineTemplate`

```bash
kubectl get clusterclass,cluster
kubectl get kubeadmcontrolplanes
kubectl get machinedeployments,machinesets,machines
kubectl get devclusters,devmachines
```

## Cleanup

Switch back to the management cluster, then delete the workload cluster before removing kind:

```bash
unset KUBECONFIG

kubectl delete cluster capi-quickstart
kubectl wait --for=delete cluster/capi-quickstart --timeout=10m

kind delete cluster --name cluster-api
```

