# Github Actions Runner Controller on Kubernetes

Let's start with the github [actions-runner-controller](https://github.com/actions/actions-runner-controller) documentation 

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
cd github/actions/kubernetes

kind create cluster --name demo --image kindest/node:v1.28.0
```

Next up, I will be running a small container where I will be doing all the work from:
You can skip this part if you already have `kubectl` and `helm` on your machine.

```
docker run -it --rm --net host -v ${HOME}/.kube/:/root/.kube/ -v ${PWD}:/work -w /work alpine sh
```

Install `kubectl`

```
apk add --no-cache curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

Install `helm`

```
HELM_VERSION=3.14.1
curl -LO https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz
tar -C /tmp/ -zxvf helm-v${HELM_VERSION}-linux-amd64.tar.gz
rm helm-v${HELM_VERSION}-linux-amd64.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm
```

Now we have `helm` and `kubectl` and we can test our cluster access:

```
kubectl get nodes

NAME                 STATUS   ROLES           AGE     VERSION
demo-control-plane   Ready    control-plane   4m12s   v1.28.0
```

## Deploy the Github Action Runner Controller

The runner controller is the core controller that manages the entire runner ecosystem. </br>
Its uses the Kubernetes Operator pattern so you can deploy runners. </br>

Let's go to the [Quickstart Guide](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/quickstart-for-actions-runner-controller) </br>

We can deploy the controller using `helm`
```
VERSION=0.8.2
NAMESPACE="github"
helm install arc \
    --namespace "${NAMESPACE}" \
    --create-namespace \
    --version ${VERSION} \
    oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set-controller

```

Once chart is deployed, we can see the controller running in the `github` namespace

```
kubectl -n github get pods
NAME                                     READY   STATUS    RESTARTS   AGE
arc-gha-rs-controller-7bf474df55-v7vm9   1/1     Running   0          16s
```

Now this pod will not do anything, other than allowing us to deploy runner scale sets. </br>
We can check its logs: 
```
kubectl -n github logs -l app.kubernetes.io/name=gha-rs-controller
```

## Deploy Github Action Runner scale sets

To make it useful, we need to add it to our repo or organisation by [Authenticating to Github](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/authenticating-to-the-github-api)

Create a new [Github App](https://docs.github.com/en/apps/using-github-apps/installing-your-own-github-app) for your account or organization.  

Once we have created our Github App, we need to configure its authentication by creating a kubernetes secret with the authentication details of the app. </br>

```
kubectl create secret generic github-app-secret \
   --namespace=github \
   --from-literal=github_app_id=xxxxx \
   --from-literal=github_app_installation_id=xxxxx \
   --from-file=github_app_private_key='github.pem'

```

The [documentation](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/deploying-runner-scale-sets-with-actions-runner-controller#configuring-the-runner-image) showcases how to customise the runner with the helm values file 

This allows us to customer the pod template for our runner. <br/>
Let's create a basic one:

```
template:
  spec:
    containers:
      - name: runner
        image: "ghcr.io/actions/actions-runner:latest"
        imagePullPolicy: Always
        command: ["/home/runner/run.sh"]
```

We'll need to use the helm upgrade to apply the changes again. </br>

We can also set our runner name using the values file using the [advanced configuration options](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/deploying-runner-scale-sets-with-actions-runner-controller#using-advanced-configuration-options)

```
runnerScaleSetName: "marcels-runner"
```

Deploy the Github actions runner scaleset 

```
INSTALLATION_NAME="arc-runner-set"
NAMESPACE="github"
GITHUB_CONFIG_URL="https://github.com/marcel-dempers/docker-development-youtube-series"

helm install "${INSTALLATION_NAME}" \
    --namespace "${NAMESPACE}" \
    --create-namespace \
    --values scaleset-values.yaml \
    --set githubConfigUrl="${GITHUB_CONFIG_URL}" \
    oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set
```

Now this will deploy a listener which will listen to jobs for the organisation or github repo \ account :

```
kubectl -n github get pods
NAME                                     READY   STATUS    RESTARTS   AGE
arc-gha-rs-controller-7bf474df55-v7vm9   1/1     Running   0          14m
marcels-runner-9d8dc86f-listener         1/1     Running   0          10s
```

If we look at the logs again we can see the controller created a listener

```
kubectl -n github logs -l app.kubernetes.io/name=gha-rs-controller
kubectl -n github logs -l app.kubernetes.io/component=runner-scale-set-listener
```

It's important to check both logs to ensure all your authentication is setup correctly and there are no problems. </br>

### Customise the runner image

Now when we run our GitHub action, it fails because there is no `docker` command  and no `docker` daemon running in the runner. </br>
To add docker there are a number of things we can do, we can setup docker in the container or make use of docker on the host or run docker as a sidecar container and leverage the `docker-in-docker` image. </br>

This approach is also explained in the [documentation](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/deploying-runner-scale-sets-with-actions-runner-controller#using-docker-in-docker-mode)

```
template:
```

To update, we simply use the `helm upgrade` command

```
INSTALLATION_NAME="arc-runner-set"
NAMESPACE="github"
GITHUB_CONFIG_URL="https://github.com/marcel-dempers/docker-development-youtube-series"

helm upgrade "${INSTALLATION_NAME}" \
    --namespace "${NAMESPACE}" \
    --values scaleset-values.yaml \
    --set githubConfigUrl="${GITHUB_CONFIG_URL}" \
    oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set
```

