# Introduction to OpenClaw

<!-- #TODO 
#<a href="https://youtu.be/xxxxxx" title="openclaw"><img src="https://i.ytimg.com/vi/xxxxxxx/hqdefault.jpg" width="40%" alt="openclaw" /></a>
-->

In this video, we will dive into the features of OpenClaw. </br>

[Documentation](https://openclaw.ai/)

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name openclaw --image kindest/node:v1.36.1
```

Test our cluster and makes sure `kubectl` is configured for it:

```
kubectl get nodes
NAME                     STATUS   ROLES           AGE     VERSION
openclaw-control-plane   Ready    control-plane   2m30s   v1.36.1
```

## Source Code

```shell
git clone --depth 1 https://github.com/openclaw/openclaw.git /tmp/openclaw
```

View kubernetes deployment scripts:

```shell
tree /tmp/openclaw/scripts/k8s/

/tmp/openclaw/scripts/k8s/
├── create-kind.sh
├── deploy.sh
└── manifests
    ├── configmap.yaml
    ├── deployment.yaml
    ├── kustomization.yaml
    ├── pvc.yaml
    └── service.yaml

2 directories, 7 files
```

I've taken a copy of these manifests by using `cp`. </br>

```shell
cp /tmp/openclaw/scripts/k8s/manifests/* ai/agents/openclaw/kubernetes/
```

You can now use Helm, Kuztomize, ArgoCD, Flux or any mechanism of your choice to manage this deployment. </br>

## Deploy Secret

We can refer to the `deployment.yaml` and we see the secret references under `secretKeyRef` </br>

We will need a secret for our Gateway UI, a secret for our model provider as well as our channel. </br>

```shell
export OPENCLAW_GATEWAY_TOKEN="$(openssl rand -hex 32)"
export GEMINI_API_KEY="xxxxxxx"
export DISCORD_BOT_TOKEN="xxxxxxx"
```

I'll be modifying the `deployment.yaml` to add my channel token `DISCORD_BOT_TOKEN` 

```yaml
- name: DISCORD_BOT_TOKEN
  valueFrom:
    secretKeyRef:
      name: openclaw-secrets
      key: DISCORD_BOT_TOKEN
      optional: false
```

My `openclaw.json` has a reference to that token for my discord channel

```json
"channels": {
    "discord": {
      "enabled": true,
      "token": {
        "source": "env",
        "provider": "default",
        "id": "DISCORD_BOT_TOKEN"
      },
      "groupPolicy": "allowlist",
      "guilds": {}
    }
  }
```

Create the secret:

```shell
kubectl create ns openclaw
kubectl create secret generic openclaw-secrets -n openclaw \
  --from-literal OPENCLAW_GATEWAY_TOKEN=${OPENCLAW_GATEWAY_TOKEN} \
  --from-literal GEMINI_API_KEY=${GEMINI_API_KEY} \
  --from-literal DISCORD_BOT_TOKEN=${DISCORD_BOT_TOKEN}

```

## Deploy Openclaw

Apply the manifests

```shell
kubectl -n openclaw apply -f ai/agents/openclaw/kubernetes/
```

## Install Plugins 

```shell
export NPM_CONFIG_CACHE=~/.openclaw/.npm
openclaw plugins install @openclaw/discord
kubectl -n openclaw delete pods --all
```