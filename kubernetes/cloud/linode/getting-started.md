# Getting Started with Linode

<a href="https://youtu.be/VSPUWEtqtnY" title="k8s-linode"><img src="https://i.ytimg.com/vi/VSPUWEtqtnY/hqdefault.jpg" width="20%" alt="k8s-linode" /></a> 

## Trial Account

Promo Link to get $20 credit to try out Linode: <br/>
https://login.linode.com/signup?promo=DOCS20AA00X1

## Linode CLI

At the time of this video there is not docker image for Linode CLI, so lets make our own :) <br/>
Take a look at the dockerfile in this folder.

```
# Linode CLI

# Run this from the root of the repo!

docker run -it --rm -v ${PWD}:/work -w /work --entrypoint /bin/bash aimvector/linode:2.15.0
cd ./kubernetes/cloud/linode

```

## Login to Linode

```
#login and follow prompts
linode-cli

```

## Gather our options

https://www.linode.com/docs/platform/api/linode-cli/

```
linode-cli lke --help

linode-cli regions list --text
linode-cli lke versions-list 
linode-cli linodes list --region ap-south

```

## Create our cluster

https://www.linode.com/docs/platform/api/linode-cli/#linode-kubernetes-engine-lke

```

# full list of options
linode-cli lke cluster-create --help

linode-cli lke cluster-create \
  --label getting-started-lke \
  --region ap-south \
  --k8s_version 1.16 \
  --node_pools.type g6-standard-2 --node_pools.count 1 \
  --tags marcel

```

## Get a kubeconfig for our cluster

```
linode-cli lke clusters-list

linode-cli lke kubeconfig-view <id>
```

Download a kubeconfig from the [dashboard](https://cloud.linode.com/kubernetes/clusters) <br/>
Rename and drop it into `./kubernetes/cloud/linode/config.yaml`

## Get kubectl

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

export KUBECONFIG=/work/kubernetes/cloud/linode/config.yaml

kubectl create ns example-app

# lets create some resources.
kubectl apply -n example-app -f /work/kubernetes/secrets/secret.yaml
kubectl apply -n example-app -f /work/kubernetes/configmaps/configmap.yaml
kubectl apply -n example-app -f /work/kubernetes/deployments/deployment.yaml

# remember to change the `type: LoadBalancer`
kubectl apply -n example-app -f /work/kubernetes/services/service.yaml

```

## Clean up 

```
linode-cli lke clusters-list
linode-cli lke cluster-delete <id>

# remember to delete the load balancer manually!
```