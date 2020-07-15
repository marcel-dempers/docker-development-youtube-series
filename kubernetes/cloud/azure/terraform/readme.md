# Getting Started with AKS using Terraform

More resources:

Terraform provider for Azure [here](https://github.com/terraform-providers/terraform-provider-azurerm) <br/>

## Azure CLI

We'll need the Azure CLI to gather information so we can build our Terraform file.

```
# Run Azure CLI
docker run -it --rm -v ${PWD}:/work -w /work --entrypoint /bin/sh mcr.microsoft.com/azure-cli:2.6.0

# Get Terraform

curl -o /tmp/terraform.zip -LO https://releases.hashicorp.com/terraform/0.12.28/terraform_0.12.28_linux_amd64.zip

unzip /tmp/terraform.zip
chmod +x terraform && mv terraform /usr/local/bin/

cd kubernetes/cloud/azure/terraform/
terraform init

```

## Login to Azure

```
#login and follow prompts
az login 

# view and select your subscription account

az account list -o table
SUBSCRIPTION=<id>
az account set --subscription $SUBSCRIPTION

```

## Create our Resource Group

```
RESOURCEGROUP=aks-getting-started
RESOURCEGROUP_ID=$(az group create -n $RESOURCEGROUP -l australiaeast | jq -r '.id')

```

## Create Service Principal

Kubernetes needs a service account to manage our Kubernetes cluster </br>
Lets create one! </br>

```

SERVICE_PRINCIPAL_JSON=$(az ad sp create-for-rbac --skip-assignment --name aks-getting-started-sp -o json)

#Keep the `appId` and `password` for later use!

SERVICE_PRINCIPAL=$(echo $SERVICE_PRINCIPAL_JSON | jq -r '.appId')
SERVICE_PRINCIPAL_SECRET=$(echo $SERVICE_PRINCIPAL_JSON | jq -r '.password')

#grant contributor role over the resource group to our service principal

az role assignment create --assignee $SERVICE_PRINCIPAL \
--scope "/subscriptions/$SUBSCRIPTION/resourceGroups/$RESOURCEGROUP" \
--role Contributor

```
For extra reference you can also take a look at the Microsoft Docs: [here](https://github.com/MicrosoftDocs/azure-docs/blob/master/articles/aks/kubernetes-service-principal.md) </br>

# Generate SSH key

```
ssh-keygen -t rsa -b 4096 -N "VeryStrongSecret123!" -C "your_email@example.com" -q -f  ~/.ssh/id_rsa
SSH_KEY=$(cat ~/.ssh/id_rsa.pub)
```

## Terraform Azure Kubernetes Provider 

Documentation on all the Kubernetes fields for terraform [here](https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster.html)

```
terraform plan -var serviceprinciple_id=$SERVICE_PRINCIPAL -var serviceprinciple_key="$SERVICE_PRINCIPAL_SECRET" -var ssh_key="$SSH_KEY"

# Import existing resource group 
terraform import -var serviceprinciple_id=$SERVICE_PRINCIPAL -var serviceprinciple_key="$SERVICE_PRINCIPAL_SECRET" -var ssh_key="$SSH_KEY"  module.cluster.azurerm_resource_group.aks-getting-started $RESOURCEGROUP_ID
terraform apply -var serviceprinciple_id=$SERVICE_PRINCIPAL -var serviceprinciple_key="$SERVICE_PRINCIPAL_SECRET" -var ssh_key="$SSH_KEY"
```