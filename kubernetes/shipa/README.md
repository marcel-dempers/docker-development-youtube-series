# Introduction to Shipa

## We need a Kubernetes cluster

To get the most out of Shipa, I'll be using real Cloud Provider Kubernetes as well as a local <br/>
`minikube` cluster. <br/>
To create a Kubernetes cluster, you can follow my guides on each Cloud provider below: <br/>

## Kubernetes in the Cloud

|Cloud   | Kubernetes   | Video  | Source Code  |   |
|---|---|---|---|---|
|Azure          | AKS  |[Source Code](../cloud/azure/getting-started.md) | <a href="https://youtu.be/eyvLwK5C2dw" title="AKS"><img src="https://i.ytimg.com/vi/eyvLwK5C2dw/hqdefault.jpg" width="25%" height="25%" alt="AKS Guide" /></a>  |   
|Amazon         | EKS  |[Source Code](../cloud/amazon/getting-started.md)  |  <a href="https://youtu.be/QThadS3Soig" title="EKS"><img src="https://i.ytimg.com/vi/QThadS3Soig/hqdefault.jpg" width="25%" height="25%" alt="EKS Guide" /></a> |   
|Google         | GKE  |[Source Code](../cloud/google/getting-started.md)  | <a href="https://youtu.be/-fbH5Qs3QXU" title="GKE"><img src="https://i.ytimg.com/vi/-fbH5Qs3QXU/hqdefault.jpg" width="25%" height="25%" alt="GKE Guide" /></a>  |   
|Digital Ocean  | DO  |[Source Code](../cloud/digitalocean/getting-started.md)   | <a href="https://youtu.be/PvfBCE-xgBY" title="DO"><img src="https://i.ytimg.com/vi/PvfBCE-xgBY/hqdefault.jpg" width="25%" height="25%" alt="DO Guide" /></a>  |   
|Linode         | LKE |[Source Code](../cloud/linode/getting-started.md)  | <a href="https://youtu.be/VSPUWEtqtnY" title="LKE"><img src="https://i.ytimg.com/vi/VSPUWEtqtnY/hqdefault.jpg" width="25%" height="25%" alt="LKE Guide" /></a>  |   


## Minikube 

I will start with a local minikube cluster to get Shipa running: <br/>

```
# start up a cluster

minikube start --kubernetes-version='v1.18.2' --memory='5gb' --disk-size='20gb' --driver=hyperv

# check our cluster

kubectl get nodes
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   45s   v1.18.2

```

# Getting Started with Shipa

## Install Dependencies

```
docker run -it --rm  -v ${PWD}:/work -w /work alpine sh
apk add --no-cache curl unzip

cd kubernetes/shipa/
mkdir installs && cd installs

```

## Install Kubectl

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.19.4/bin/windows/amd64/kubectl.exe

```

## Install Helm CLI

```
curl -LO https://get.helm.sh/helm-v3.4.1-windows-amd64.zip && \
unzip helm-v3.4.1-windows-amd64.zip && \
mv windows-amd64/helm.exe . && \
rm -rf windows-amd64 && \
rm helm-v3.4.1-windows-amd64.zip

```

## Install Shipa CLI

```
curl -LO  https://storage.googleapis.com/shipa-client/v1.1/shipa_windows_amd64.exe
mv shipa_windows_amd64.exe shipa.exe
```


## Add all the CLI to our $PATH

We can now add the CLI executables to a folder somewhere on our machine <br/>
Then we add it to our `$PATH`

* installs/helm.exe
* installs/kubectl.exe
* installs/shipa.exe

## Download Shipa Helm Chart

Official Docs [here](https://learn.shipa.io/docs/installing-shipa)
We can find all the releases of Shipa [here](https://github.com/shipa-corp/helm-chart/releases)
In this demo, I will use version `1.1.1`

Let's download Shipa: <br/>

```
curl  -L -s -o shipa1.1.1.zip https://github.com/shipa-corp/helm-chart/archive/v1.1.1.zip && \
unzip shipa1.1.1.zip && rm shipa1.1.1.zip && \
mv helm-chart-1.1.1 shipa-helm-chart-1.1.1 && \

# we can abandon this container
exit

```

## Install Shipa

Let's add Shipa to our `minikube` cluster: <br/>

```
cd .\kubernetes\shipa\installs\shipa-helm-chart-1.1.1\

kubectl apply -f limits.yaml

# deploy shipa dependencies

helm dep up

# install

helm install shipa . `
--timeout=15m `
--set=metrics.image=gcr.io/shipa-1000/metrics:30m `
--set=auth.adminUser=admin@shipa.io `
--set=auth.adminPassword=shipa2020 `
--set=shipaCore.serviceType=ClusterIP `
--set=shipaCore.ip=10.100.10.20 `
--set=service.nginx.serviceType=ClusterIP `
--set=service.nginx.clusterIP=10.100.10.10

# ensure everything is up and running

kubectl get pods
NAME                                       READY   STATUS      RESTARTS   AGE
dashboard-web-1-6f8b58fb89-bjf7c           1/1     Running     0          12m
node-container-busybody-theonepool-zz7sw   1/1     Running     0          15m
node-container-netdata-theonepool-kqb5z    1/1     Running     0          15m
shipa-api-57b69645d9-rd2bz                 1/1     Running     0          21m
shipa-clair-d7554fc6f-8nqgz                1/1     Running     1          21m
shipa-docker-registry-5885d6f467-dvkjb     1/1     Running     0          18m
shipa-etcd-85cc6c6458-6cgx6                1/1     Running     1          21m
shipa-guardian-5466f58668-25zkf            1/1     Running     0          16m
shipa-init-job-1-9xdgw                     0/1     Completed   0          21m
shipa-metrics-786468c5cc-h7zfb             1/1     Running     0          21m
shipa-mongodb-replicaset-0                 1/1     Running     0          21m
shipa-nginx-ingress-75dccdb4fb-nq7xq       1/1     Running     0          21m
shipa-postgres-7c55df4758-7s64w            1/1     Running     0          21m

kubectl get svc 
NAME                              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                                      AGE
dashboard-web-1                   ClusterIP   10.96.3.55       <none>        8888/TCP                                     12m
dashboard-web-1-units             ClusterIP   None             <none>        8888/TCP                                     12m
kubernetes                        ClusterIP   10.96.0.1        <none>        443/TCP                                      6m
shipa-api                         ClusterIP   10.100.120.116   <none>        8080/TCP,8081/TCP                            21m
shipa-clair                       ClusterIP   10.111.91.38     <none>        6060/TCP,6061/TCP                            21m
shipa-docker-registry             ClusterIP   10.104.180.204   <none>        5000/TCP                                     21m
shipa-etcd                        ClusterIP   10.110.16.90     <none>        2379/TCP                                     21m
shipa-guardian                    ClusterIP   10.97.114.38     <none>        8000/TCP,22/TCP                              21m
shipa-ingress-nginx               ClusterIP   10.100.10.10     <none>        22/TCP,5000/TCP,8081/TCP,8080/TCP,2379/TCP   21m
shipa-metrics                     ClusterIP   10.108.223.229   <none>        9090/TCP,9091/TCP                            21m
shipa-mongodb-replicaset          ClusterIP   None             <none>        27017/TCP                                    21m
shipa-mongodb-replicaset-client   ClusterIP   None             <none>        27017/TCP                                    21m
shipa-postgres                    ClusterIP   10.104.18.123    <none>        5432/TCP                                     21m
```

## Targets

```
# add a route for accessing Shipa API
route add 10.100.10.10/32 MASK 255.255.255.255  $(minikube ip)

# add a route for accessing our Applications
route add 10.100.10.20/32 MASK 255.255.255.255  $(minikube ip)

shipa target-add dev 10.100.10.10
shipa target-list
shipa target-set dev
```

## Pools 

https://learn.shipa.io/docs/pool-management

```
shipa pool-add prod --public --kube-namespace blue-team --provisioner kubernetes
shipa pool-list

```

## Clusters

```
shipa cluster-list
+------------+-------------+---------------+-------------+---------+------------+-------+-------+
| Name       | Provisioner | Addresses     | Custom Data | Default | Pools      | Teams | Error |
+------------+-------------+---------------+-------------+---------+------------+-------+-------+
| shipa-core | kubernetes  | 10.96.0.1:443 |             | false   | theonepool |       |       |
+------------+-------------+---------------+-------------+---------+------------+-------+-------+

```

## Applications

```
shipa login
shipa app-create go-helloworld static -t admin -o theonepool


cd kubernetes\shipa\developers

docker build .-t aimvector/shipa-golang:v1
docker push aimvector/shipa-golang:v1

shipa app-deploy -i aimvector/shipa-golang:v1 -a go-helloworld


cd .\kubernetes\shipa\developers\docker\python\

docker build . -t aimvector/shipa-python:v1
docker push aimvector/shipa-python:v1

shipa app-create python-helloworld static -t admin -o theonepool
shipa env set FLASK_APP=/app/server.py -a python-helloworld
shipa app-deploy -i aimvector/shipa-python:v1 -a python-helloworld



# deploy to prod

shipa app-create python-helloworld-prod static -t admin -o prod
shipa env set FLASK_APP=/app/server.py -a python-helloworld-prod
shipa app-deploy -i aimvector/shipa-python:v1 -a python-helloworld-prod


shipa app-create go-helloworld-prod static -t admin -o prod
shipa app-deploy -i aimvector/shipa-golang:v1 -a go-helloworld-prod

```


```
kubectl apply -f shipa-admin-service-account.yaml

# get the sa token
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep shipa-admin | awk '{print $1}')
# get the k8s CA

kubectl get secret $(kubectl get secret | grep default-token | awk '{print $1}') -o jsonpath='{.data.ca\.crt}' | base64 -d

```

























```

https://collabnix.com/the-rise-of-shipa-a-continuous-operation-platform/

1. Configured default user:

Username: admin@admin.com
Password: adminadmin123

2. If this is a production cluster, please configure persistent volumes.
   The default reclaimPolicy for dynamically provisioned persistent volumes is "Delete" and
   users are advised to change it for production

   The code snippet below can be used to set reclaimPolicy to "Retain" for all volumes:

PVCs=$(kubectl --namespace=shipa-system get pvc -l release=shipa -o name)

for pvc in $PVCs; do
    volumeName=$(kubectl -n shipa-system get $pvc -o template --template=\{\{.spec.volumeName\}\})
    kubectl -n shipa-system patch pv $volumeName -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'
done

3. Set default target for shipa-client:
export SHIPA_HOST=$(kubectl --namespace=shipa-system get svc shipa-ingress-nginx  -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

shipa target-add shipa $SHIPA_HOST -s

shipa login admin@admin.com
shipa node-list
shipa app-list

```