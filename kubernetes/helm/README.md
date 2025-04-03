# Introduction to Helm

<a href="https://youtu.be/5_J7RWLLVeQ" title="k8s-helm"><img src="https://i.ytimg.com/vi/5_J7RWLLVeQ/hqdefault.jpg" width="20%" alt="k8s-helm" /></a> 

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name helm --image kindest/node:v1.31.1
```

# Getting Started with Helm

Firstly, I like to do most of my work in containers so everything is reproducible  <br/>
and my machine remains clean.

## Get a container to work in
<br/>
Run a small `alpine linux` container where we can install and play with `helm`: <br/>

```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh

# install curl & kubectl
apk add --no-cache curl nano
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
export KUBE_EDITOR="nano"

# test cluster access:
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
helm-control-plane   Ready    control-plane   26m   v1.31.1

```

## Install Helm CLI

```
curl -LO https://get.helm.sh/helm-v3.17.2-linux-amd64.tar.gz
tar -C /tmp/ -zxvf helm-v3.17.2-linux-amd64.tar.gz
rm helm-v3.17.2-linux-amd64.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm

```

## Create our first Chart

For reference in the rest of the guide, I have left my full templates in: <br/>
`<GitRepo>/kubernetes/helm/example-app`
```
cd kubernetes/helm

mkdir temp && cd temp
helm create example-app
```

## Cleanup the template 

We can delete unwanted files:

* delete everything under /templates, keeping only `_helpers.tpl`
* delete `tests` folder under `templates`

## Add Kubernetes files to our new Chart

Copy the following files into our `example-app/templates/` folder

* `<GitRepo>/kubernetes/deployments/deployment.yaml`
* `<GitRepo>/kubernetes/services/service.yaml`
* `<GitRepo>/kubernetes/configmaps/configmap.yaml`
* `<GitRepo>/kubernetes/secrets/secret.yaml`

## Test the rendering of our template

```
helm template example-app example-app
```

## Install our app using our Chart

```
helm install example-app example-app

# list our releases

helm list

# see our deployed components

kubectl get all
kubectl get cm
kubectl get secret
```
 
## Value injections for our Chart

For CI systems, we may want to inject an image tag as a build number <br/>

Basic parameter injection: <br/>

```
# values.yaml

deployment:
  image: "aimvector/python"
  tag: "1.0.4"

# deployment.yaml

image: {{ .Values.deployment.image }}:{{ .Values.deployment.tag }}

# upgrade our release

helm upgrade example-app example-app --values ./example-app/values.yaml

# see revision increased

helm list
```

## Make our Chart more generic

Let's make our chart generic so it can be reused: <br/>
For the following objects, replace `example-deploy` and `example-app` to inject: `"{{ .Values.name }}"`

* deployment.yaml
* services.yaml
* secret.yaml
* configmap.yaml

Now that our application is generic <br/>
We can deploy another copy of it.<br/>
<br/>

Rename `values.yaml` to `example-app.values.yaml`
Create our second app values file `example-app-02.values.yaml`

```
helm install example-app-02 example-app --values ./example-app/example-app-02.values.yaml
```

## Trigger deployment change when config changes

By default, a deployment will not rollout new pods when a configmap changes. <br/>
Some application read configuration at start up and deployment may need to roll out </br>
new pods when a configmap changes. Let's do that:

```
# deployment.yaml

kind: Deployment
spec:
  template:
    metadata:
      annotations:
        checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}

# rollout the change

helm upgrade example-app example-app --values ./example-app/example-app-01.values.yaml
```

## If\Else and Default values

You can also set default values in case they are not supplied by the `values.yaml` file. <br/>
This may help you keep the `values.yaml` file small <br/>

```
{{- if .Values.deployment.resources }}
  resources:
    {{- if .Values.deployment.resources.requests }}
    requests:
      memory: {{ .Values.deployment.resources.requests.memory | default "50Mi" | quote }}
      cpu: {{ .Values.deployment.resources.requests.cpu | default "10m" | quote }}
    {{- else}}
    requests:
      memory: "50Mi"
      cpu: "10m"
    {{- end}}
    {{- if .Values.deployment.resources.limits }}
    limits:
      memory: {{ .Values.deployment.resources.limits.memory | default "1024Mi" | quote }}
      cpu: {{ .Values.deployment.resources.limits.cpu | default "1" | quote }}
    {{- else}}  
    limits:
      memory: "1024Mi"
      cpu: "1"
    {{- end }}
  {{- else }}
  resources:
    requests:
      memory: "50Mi"
      cpu: "10m"
    limits:
      memory: "1024Mi"
      cpu: "1"
  {{- end}} 


# rollout the change

helm upgrade example-app example-app --values ./example-app/example-app-01.values.yaml
```
