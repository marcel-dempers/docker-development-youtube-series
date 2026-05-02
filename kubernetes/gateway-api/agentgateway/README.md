# Introduction to Agent Gateway: LLM Gateway

<!-- #TODO: YouTube link -->

[Official Site](https://agentgateway.dev) | [GitHub](https://github.com/agentgateway/agentgateway) | [Documentation](https://agentgateway.dev/docs/kubernetes/latest/)

## Prerequisites

To get started, you will need to follow the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs
* Example models, either in cluster or external (OpenAI, Anthropic, Gemini)
* Have Domains for our traffic
* Have TLS certificates

This will allow us access to the Gateway API so we can go ahead and deploy a Gateway API controller to use. </br

Let's start with [Official Site](https://agentgateway.dev) | [GitHub](https://github.com/agentgateway/agentgateway) | [Documentation](https://agentgateway.dev/docs/kubernetes/latest/)

## Install AgentGateway

Agent Gateway is distributed as two Helm charts: one for its CRDs, one for the control plane. </br>


To get the version you want, you can use the agent gateway [Github Release page](https://github.com/agentgateway/agentgateway/releases). </br>
For this guide, we will use 1.1.0. </br>

You can list helm chart versions too:

```shell
helm show chart oci://cr.agentgateway.dev/charts/agentgateway-crds
helm show chart oci://cr.agentgateway.dev/charts/agentgateway
```

```shell
CHART_VERSION="v1.1.0"

# Install CRDs

helm upgrade -i agentgateway-crds oci://cr.agentgateway.dev/charts/agentgateway-crds \
  --create-namespace \
  --namespace agentgateway-system \
  --version $CHART_VERSION

# Install ControlPlane
helm upgrade -i agentgateway oci://cr.agentgateway.dev/charts/agentgateway \
  --namespace agentgateway-system \
  --version $CHART_VERSION \
  --set controller.image.pullPolicy=Always \
  --wait
```

Check our installation

```shell
# check pods are running
kubectl -n agentgateway-system get pods

# check logs
kubectl -n agentgateway-system logs -l app.kubernetes.io/name=agentgateway
```


## Check the AgentGateway Class

Verify the `GatewayClass` was registered:

```shell
kubectl get gatewayclass
```

## Install an AgentGateway Gateway

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/01-gateway.yaml

# we can see our gateway dataplane pods
kubectl get pods
kubectl get svc

# check gateway 
kubectl get gateway

# port forward for access
kubectl -n default port-forward svc/agentgateway 80
```

## Configuring an LLM Backend

AgentGateway introduces CRD `AgentgatewayBackend`. </br>
`AgentgatewayBackend` is an LLM provider as a routable backend.</br>


### LLM Backend

If your LLM Provider needs an API key, we need a Kubernetes secret for it:


```shell

export GEMINI_API_KEY=""

kubectl apply -f- <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: gemini-secret
type: Opaque
stringData:
  Authorization: $GEMINI_API_KEY
EOF
```

Create the LLM backend: </br>

Instead of an HTTPRoute routing to an upstream backend `service`, it uses an `AgentgatewayBackend`

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/02-backend-gemini.yaml
```

Route traffic from the Gateway to the backend using a standard `HTTPRoute`:

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/03-httproute.yaml
```

Test our HTTPRoute:

```shell
curl "localhost:8080/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d '{
      "messages": [{"role": "user", 
      "content": "What is Kubernetes in one sentence?"}]
  }' | jq
```

We can route by path as well. This means one domain and route to provider by path

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/04-httproute-gemini.yaml
```

```shell
curl "localhost:8080/ai/gemini/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d '{
      "messages": [{"role": "user", 
      "content": "What is Kubernetes in one sentence?"}]
  }' | jq
```

### Anthropic Backend

The same pattern works for any supported provider.</br>

```shell
export ANTHROPIC_API_KEY=""

kubectl apply -f- <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: anthropic-secret
type: Opaque
stringData:
  Authorization: $ANTHROPIC_API_KEY
EOF
```

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/05-backend-anthropic.yaml
```

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/06-httproute-anthropic.yaml
```

Route to Anthropic Claude:

```shell
curl "localhost:8080/ai/claude/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d '{
      "messages": [{"role": "user", 
      "content": "What is Kubernetes in one sentence?"}]
  }' | jq
```

Now we have implemented LLM routing by path using `HTTPRoute`:

`localhost:8080/ai/gemini` --> `gemini flash` </br>
`localhost:8080/ai/claude` --> `claude sonnet` </br>

