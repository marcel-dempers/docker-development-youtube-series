# Getting Started with DGO

<a href="https://youtu.be/PvfBCE-xgBY" title="k8s-do"><img src="https://i.ytimg.com/vi/PvfBCE-xgBY/hqdefault.jpg" width="20%" alt="k8s-do" /></a> 

## Trial Account

Coupon Link to get $100 credit for 60 days: <br/>
https://m.do.co/c/74a1c5d63dac

## Digital Ocean CLI

https://hub.docker.com/r/digitalocean/doctl

```
# Digital Ocean CLI
docker run -it --rm -v ${PWD}:/work -w /work --entrypoint /bin/bash digitalocean/doctl:1.45.0
mv /app/doctl /usr/local/bin/
cd ./kubernetes/cloud/digitalocean

```

## Login to DGO

```
#login and follow prompts
doctl auth init
doctl auth list

```

## Create a new project

```
doctl projects create --name getting-started-dgo --purpose testing
doctl projects list
# grab the project ID
```

## Gather our options

https://www.digitalocean.com/docs/kubernetes/

```
doctl kubernetes options
doctl kubernetes options regions
doctl kubernetes options versions

```

## Create our cluster

```




# full list of options
doctl kubernetes cluster create --help

doctl kubernetes cluster create dgo-getting-started \
--version 1.17.5-do.0 \
--count 1 \
--size s-1vcpu-2gb \
--region sgp1

```

## Get a kubeconfig for our cluster

```
doctl kubernetes cluster kubeconfig save dgo-getting-started

#grab the config if you want it
cp ~/.kube/config .

```

## Get kubectl

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

cd ../..


kubectl create ns example-app

# lets create some resources.
kubectl apply -n example-app -f secrets/secret.yaml
kubectl apply -n example-app -f configmaps/configmap.yaml
kubectl apply -n example-app -f deployments/deployment.yaml

# remember to change the `type: LoadBalancer`
kubectl apply -n example-app -f services/service.yaml

```

## Clean up 

```
doctl kubernetes cluster delete dgo-getting-started

# remember to delete the load balancer manually!
```