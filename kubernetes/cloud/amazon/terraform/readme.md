# Getting Started with Amazon EKS using Terraform

More resources:

Terraform provider for AWS [here](https://www.terraform.io/docs/providers/aws/index.html) <br/>

## Amazon CLI

You can get the Amazon CLI on [Docker-Hub](https://hub.docker.com/r/amazon/aws-cli) <br/>
We'll need the Amazon CLI to gather information so we can build our Terraform file.

```
# Run Amazon CLI
docker run -it --rm -v ${PWD}:/work -w /work --entrypoint /bin/sh amazon/aws-cli:2.0.17

# some handy tools :)
yum install jq gzip nano tar git unzip wget

```

## Login to Amazon

```
# Access your "My Security Credentials" section in your profile. 
# Create an access key

aws configure

```

# Terraform CLI 

```
# Get Terraform

curl -o /tmp/terraform.zip -LO https://releases.hashicorp.com/terraform/0.12.28/terraform_0.12.28_linux_amd64.zip

unzip /tmp/terraform.zip
chmod +x terraform && mv terraform /usr/local/bin/

cd kubernetes/cloud/amazon/terraform/

```

# Generate SSH key

```
ssh-keygen -t rsa -b 4096 -N "VeryStrongSecret123!" -C "your_email@example.com" -q -f  ~/.ssh/id_rsa
SSH_KEY=$(cat ~/.ssh/id_rsa.pub)
```

## Terraform Amazon Kubernetes Provider 

Documentation on all the Kubernetes fields for terraform [here](https://www.terraform.io/docs/providers/aws/r/eks_cluster.html)

```
terraform init

terraform plan -var access_key=$access_key -var secret_key=$secret_key

terraform apply -var access_key=$access_key -var secret_key=$secret_key

```

# Lets see what we deployed

```
# grab our EKS config
aws eks update-kubeconfig --name eks-getting-started --region ap-southeast-2

# Get kubectl

curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

kubectl get svc

```

# Clean up 

```
terraform destroy -var access_key=$access_key -var secret_key=$secret_key
```