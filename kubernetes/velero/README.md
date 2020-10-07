# Introduction to Velero

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name velero --image kindest/node:v1.19.1
```

## Get a container to work in
<br/>
Run a small `alpine linux` container where we can install and play with `velero`: <br/>

```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh

# install curl & kubectl
apk add --no-cache curl nano
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
export KUBE_EDITOR="nano"

#test cluster access:
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
velero-control-plane   Ready    master   26m   v1.18.4

```

## Velero CLI

Lets download the `velero` command line tool <br/>
I grabbed the `v1.5.1` release using `curl`

You can go to the [releases](https://github.com/vmware-tanzu/velero/releases/tag/v1.5.1) page to get it

```
curl -L -o /tmp/velero.tar.gz https://github.com/vmware-tanzu/velero/releases/download/v1.5.1/velero-v1.5.1-linux-amd64.tar.gz 
tar -C /tmp -xvf /tmp/velero.tar.gz
mv /tmp/velero-v1.5.1-linux-amd64/velero /usr/local/bin/velero
chmod +x /usr/local/bin/velero

velero --help
```


## Deploy some stuff

```
kubectl apply -f kubernetes/configmaps/configmap.yaml
kubectl apply -f kubernetes/secrets/secret.yaml
kubectl apply -f kubernetes/deployments/deployment.yaml
kubectl apply -f kubernetes/services/service.yaml

kubectl get all 
```

## Create storage in Azure and AWS

In this example, we'll create a storage in AWS and Azure to try both scenarios.</br>
You can follow along in the video </br>

Create a storage account and secret for: [Azure](./azure.md) </br>
Create a storage account and secret for: [AWS](./aws.md) </br>

```


```
## Deploy Velero for Azure

Start [here](./azure.md) </br>

```

# Azure credential file
cat << EOF  > /tmp/credentials-velero
AZURE_STORAGE_ACCOUNT_ACCESS_KEY=${AZURE_STORAGE_ACCOUNT_ACCESS_KEY}
AZURE_CLOUD_NAME=AzurePublicCloud
EOF

velero install \
    --provider azure \
    --plugins velero/velero-plugin-for-microsoft-azure:v1.1.0 \
    --bucket $BLOB_CONTAINER \
    --secret-file /tmp/credentials-velero \
    --backup-location-config resourceGroup=$AZURE_BACKUP_RESOURCE_GROUP,storageAccount=$AZURE_STORAGE_ACCOUNT_NAME,storageAccountKeyEnvVar=AZURE_STORAGE_ACCOUNT_ACCESS_KEY,subscriptionId=$AZURE_BACKUP_SUBSCRIPTION_ID \
    --use-volume-snapshots=false


kubectl -n velero get pods
kubectl logs deployment/velero -n velero

```

## Deploy Velero for AWS

Start [here](./aws.md)

```

cat > /tmp/credentials-velero <<EOF
[default]
aws_access_key_id=$AWS_ACCESS_ID
aws_secret_access_key=$AWS_ACCESS_KEY
EOF

velero install \
    --provider aws \
    --plugins velero/velero-plugin-for-aws:v1.1.0 \
    --bucket $BUCKET \
    --backup-location-config region=$REGION \
    --snapshot-location-config region=$REGION \
    --secret-file /tmp/credentials-velero

kubectl -n velero get pods
kubectl logs deployment/velero -n velero

```

## Create a Backup 

```
velero backup create default-namespace-backup --include-namespaces default

# describe
velero backup describe default-namespace-backup

# logs
velero backup logs default-namespace-backup
```

## Do a Restore

```
# delete all resources

kubectl delete -f kubernetes/configmaps/configmap.yaml
kubectl delete -f kubernetes/secrets/secret.yaml
kubectl delete -f kubernetes/deployments/deployment.yaml
kubectl delete -f kubernetes/services/service.yaml

velero restore create default-namespace-backup --from-backup default-namespace-backup

# describe
velero restore describe default-namespace-backup

#logs 
velero restore logs default-namespace-backup

# see items restored

kubectl get all
```
