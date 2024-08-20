
# Whats new 👉🏽 Datree in 2023

<a href="https://youtu.be/iwoIjzS33qE" title="Kubernetes"><img src="https://i.ytimg.com/vi/iwoIjzS33qE/hqdefault.jpg" width="20%" alt="Kubernetes Guide" /></a> 

## Create a Kubernetes cluster

Let's start by creating a local `kind` [cluster](https://kind.sigs.k8s.io/)

Note that we create a Kubernetes 1.23 cluster. </br>
So we want to use `datree` to validate and ensure our manifests comply with that version of Kubernetes. <br/>

```
kind create cluster --name datree --image kindest/node:v1.23.6
```

## Installation 

Best place to start is the [documentation](https://hub.datree.io/)

I like to start all my work inside a docker container. </br>
Let's run a small Alpine linux container

```
docker run -it -v ${PWD}:/work -v ${HOME}/.kube/:/root/.kube/ -w /work --net host alpine sh 
```
### Install Kubectl

Let's install `kubectl` in our container </br>

```
apk add curl jq
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.23.6/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

### Install Helm

Let's install `helm` in our container </br>

```
curl -L https://get.helm.sh/helm-v3.5.4-linux-amd64.tar.gz -o /tmp/helm.tar.gz && \
tar -xzf /tmp/helm.tar.gz -C /tmp && \
chmod +x /tmp/linux-amd64/helm && \
mv /tmp/linux-amd64/helm /usr/local/bin/helm

```

## Install Datree on our cluster

Add the Helm repo:
```
helm repo add datree-webhook https://datreeio.github.io/admission-webhook-datree
helm search repo datree-webhook --versions
```

Install the Helm chart:

```
CHART_VERSION="0.3.22"
DATREE_TOKEN=""

helm install datree-webhook datree-webhook/datree-admission-webhook \
--create-namespace \
--set datree.token=${DATREE_TOKEN} \
--set datree.policy="Default" \
--set datree.clusterName=$(kubectl config current-context) \
--version ${CHART_VERSION} \
--namespace datree 

```

Check the install

```
kubectl -n datree get pods
```

## View our Cluster Score

Now with Datree installed in our cluster, we can review it's current scoring in the Datree [Dashboard](https://app.datree.io/overview) </br>
As we are running a test cluster or if you run in the cloud, there may be some cloud components in namespaces that you may want to ignore. </br>

We can do this by labeling a namespace which is [documented here](https://hub.datree.io/configuration/behavior#ignore-a-namespace) </br>
</p>
OR </br>

We can do this by using the [configuration file](https://hub.datree.io/configuration/behavior#ignore-a-namespace) for datree


```
# skip namespace using label
kubectl label namespaces local-path-storage "admission.datree/validate=skip"
# skip namespace using configmap

kubectl -n datree apply -f kubernetes/datree/configuration/config.yaml
kubectl rollout restart deployment -n datree
```

According to the dashboard, we still have a `D` score, let's rerun the scan:

```
kubectl get job "scan-job" -n datree -o json | jq 'del(.spec.selector)' | jq 'del(.spec.template.metadata.labels)' | kubectl replace --force -f -
```

Now we can see that we have an `A` score. </br>

## Deploy some workloads to our cluster

For most companies and larger teams, it's extremely difficult to fix policy issues. </br>
Let's walk through what this may look like. </br>

Deploy some sample workloads: 

```
kubectl create namespace cms
kubectl -n cms create configmap mysql \
--from-literal MYSQL_RANDOM_ROOT_PASSWORD=1

kubectl -n cms create secret generic wordpress \
--from-literal WORDPRESS_DB_HOST=mysql \
--from-literal WORDPRESS_DB_USER=exampleuser \
--from-literal WORDPRESS_DB_PASSWORD=examplepassword \
--from-literal WORDPRESS_DB_NAME=exampledb

kubectl -n cms create secret generic mysql \
--from-literal MYSQL_USER=exampleuser \
--from-literal MYSQL_PASSWORD=examplepassword \
--from-literal MYSQL_DATABASE=exampledb

kubectl -n cms apply -f kubernetes/datree/example/cms/
```

Check out workloads 

```
kubectl -n cms get all
```

Rerun our scan:

```
kubectl delete jobs/scan-job -n datree; kubectl create job --from=cronjob/scan-cronjob scan-job -n datree
```

Now we can follow the dashboard, to check our `namespace` for policy issues and start fixing them. </br>


Summary of our fixes:

```
spec:
  containers:
    - name: wordpress
      image: wordpress:5.9-apache

kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: wordpress
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
          resources:
            limits:
              memory: "500Mi"
            requests:
              memory: "500Mi"

spec:
  containers:
    - name: wordpress
      livenessProbe:
        httpGet:
          path: /
          port: 80
      readinessProbe:
        httpGet:
          path: /
          port: 80

kind: Deployment
spec:
  template:
    spec:
      containers:
      - name: wordpress
        volumeMounts:
        - mountPath: /tmp
          name: temp
        - mountPath: /var/run/apache2/
          name: apache
      volumes:
      - emptyDir: {}
        name: temp
      - emptyDir: {}
        name: apache

kubectl -n cms apply -f kubernetes/datree/example/cms/
```
## Datree CLI : Testing our YAML locally

We can install the latest version of Datree with the command advertised:

```
apk add unzip
curl https://get.datree.io | /bin/sh
```

### Policy check

Let's test my example manifests under our datree folder `kubernetes\datree\example`

```
datree test ./kubernetes/datree/example/cms/*.yaml
```

# CI/CD examples

The tools as well as the dashboards help us solve these policy issues locally. </br>
Once we have sorted out our policy issues, we can add Datree to our CI/CD pipeline. </br>

Checkout the [CI/CD integrations](https://hub.datree.io/cicd-examples) page. </br>

# Enforcing Policies

Configure Datree to enforce policies. </br>
We can use `helm upgrade` with the `--set` flag and set enforce to true like:

```
--set datree.enforce=true 
```

Let's apply it to a new manifest and deploy it to our cluster:

```
helm upgrade datree-webhook datree-webhook/datree-admission-webhook \
--create-namespace \
--set datree.enforce=true \
--set datree.policy="Default" \
--set datree.token=${DATREE_TOKEN} \
--set datree.clusterName=$(kubectl config current-context) \
--version ${CHART_VERSION} \
--namespace datree 
```

Try to apply our Wordpress MySQL which violates policies :

```
kubectl -n cms apply -f kubernetes/datree/example/cms/statefulset.yaml
```
