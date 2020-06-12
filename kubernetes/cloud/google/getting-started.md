# Getting Started with GKE

## Google Cloud CLI

https://hub.docker.com/r/google/cloud-sdk/

```
# Run Google Cloud CLI
docker run -it --rm -v ${PWD}:/work -w /work --entrypoint /bin/bash google/cloud-sdk:160.0.0

cd ./kubernetes/cloud/google

```

## Login to GCloud

```
#login and follow prompts
gcloud auth login

gcloud projects list

gcloud projects create getting-started-gke
gcloud config set project getting-started-gke

```

## Enable APIs for your Project.

You may be prompted to enable APIs in Google Console for your project in order to proceed.
Follow the  prompts.

## Create our cluster

Machine types : https://cloud.google.com/compute/docs/machine-types

```
# machine types
gcloud compute machine-types list > machine-types.log

# Get k8s versions for your zone
gcloud container get-server-config --zone australia-southeast1-c

# full list of options
gcloud container clusters create --help

gcloud container clusters create gke-getting-started \
--cluster-version 1.16.8-gke.15 \
--disk-size 200 \
--num-nodes 1 \
--machine-type e2-small \
--no-enable-cloud-endpoints \
--no-enable-cloud-logging \
--no-enable-cloud-monitoring  \
--zone australia-southeast1-c

```

## Get a kubeconfig for our cluster

```
gcloud container clusters get-credentials gke-getting-started --zone australia-southeast1-c

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
gcloud container clusters delete gke-getting-started --zone australia-southeast1-c
```