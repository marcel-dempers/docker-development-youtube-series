# Github Actions Runner Controller on Kubernetes

Let's start with the github [actions-runner-controller](https://github.com/actions/actions-runner-controller) documentation 

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
cd github/actions/kubernetes

kind create cluster --name demo --image kindest/node:v1.28.0
```

Test the cluster:
```
kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
demo-control-plane   Ready    control-plane   59s   v1.28.0

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
```

Now this pod will not do anything, since its not authenticated with any Github organisation or repository </br>
We can check the logs to see what it's doing: 

```
kubectl -n github logs -l app.kubernetes.io/name=gha-rs-controller
```

To make it useful, we need to add it to our repo or organisation by [Authenticating to Github](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/authenticating-to-the-github-api)


Create a new [Github App](https://docs.github.com/en/apps/using-github-apps/installing-your-own-github-app) for your account or organization.  


Once we have created our Github App, we need to configure its authentication by creating a kubernetes secret with the authentication details of the app. </br>
When we have a `.pem` file , we can convert it to a single line: 

`awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' yourfile.pem`

```
kubectl create secret generic github-app-secret \
   --namespace=github \
   --from-literal=github_app_id=xxxxx \
   --from-literal=github_app_installation_id=xxxxx \
   --from-literal=github_app_private_key='-----BEGIN RSA PRIVATE KEY-----xxxxxx-----END RSA PRIVATE KEY-----'

```
Now we can upgrade our chart to apply the authentication changes

```
VERSION=0.8.2
NAMESPACE="github"
helm upgrade arc \
    --namespace "${NAMESPACE}" \
    --create-namespace \
    --version ${VERSION} \
    --values values.yaml \
    oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set-controller
```

## Create your own runner

The [documentation](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/deploying-runner-scale-sets-with-actions-runner-controller#configuring-the-runner-image) showcases how to customise the runner with the helm values file 

This allows us to customer the pod template for our runner. <br/>
Let's create a basic one:

```
template:
  spec:
    containers:
      - name: runner
        image: "custom-registry.io/actions-runner:latest"
        imagePullPolicy: Always
        command: ["/home/runner/run.sh"]
```

We'll need to use the helm upgrade to apply the changes again. </br>

We can also set our runner name using the values file using the [advanced configuration options](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners-with-actions-runner-controller/deploying-runner-scale-sets-with-actions-runner-controller#using-advanced-configuration-options)

```
runnerScaleSetName: "marcels-runner"
```