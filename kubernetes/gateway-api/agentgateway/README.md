# Introduction to Agent Gateway

<!-- #TODO: YouTube link LLM Gateway -->

<!-- #TODO: YouTube link MCP Gateway -->

[Official Site](https://agentgateway.dev) | [GitHub](https://github.com/agentgateway/agentgateway) | [Documentation](https://agentgateway.dev/docs/kubernetes/latest/)

## Prerequisites

To get started, you will need to follow the [Introduction to Gateway API](../README.md) first. </br>
You'll need an understanding of the Gateway API. </br>

<b>In the introduction guide, you will:</b>
* Create a local Kubernetes cluster
* Install the Gateway API CRDs
* Example models, either in cluster or external (OpenAI, Anthropic, Gemini)

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
kubectl -n default port-forward svc/agentgateway 8080:80
```

## Features

* Route to [LLM Models](#feature-llm-routing)
* Route to [MCP Servers](#feature-mcp-routing)

### Feature: LLM Routing

AgentGateway introduces CRD `AgentgatewayBackend`. </br>
`AgentgatewayBackend` is an LLM provider as a routable backend.</br>


#### Create an LLM Backend: Gemini

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
kubectl apply -f kubernetes/gateway-api/agentgateway/llm/gemini/02-backend-gemini.yaml
```

Route traffic from the Gateway to the backend using a standard `HTTPRoute`:

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/llm/03-httproute.yaml
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
kubectl apply -f kubernetes/gateway-api/agentgateway/llm/04-httproute-gemini.yaml
```

```shell
curl "localhost:8080/ai/gemini/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d '{
      "messages": [{"role": "user", 
      "content": "What is Kubernetes in one sentence?"}]
  }' | jq
```

#### Create an LLM Backend: Anthropic

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
kubectl apply -f kubernetes/gateway-api/agentgateway/llm/anthropic/05-backend-anthropic.yaml
```

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/llm/06-httproute-anthropic.yaml
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

### Feature: MCP Routing

To route to MCP servers, we will need some example MCP servers setup. </br>
In this example we will use a Kubernetes MCP and setup a read-only `ServiceAccount` so it can read our cluster. </br>

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/kubernetes/02-mcp.yaml

# check mcp server
kubectl get pods
kubectl get sa
kubectl get service
```

This MCP server will serve traffic on `/mcp` </br>
Let's create a backend and an `HTTPRoute` for that. </br>

#### Create an MCP Backend: Kubernetes

Instead of an HTTPRoute routing to an upstream backend `service`, it uses an `AgentgatewayBackend` </br>
Create the MCP backend: </br>
```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/kubernetes/03-mcp-backend.yaml

# check the backend
kubectl get agentgatewaybackend
kubectl describe agentgatewaybackend
```

Route traffic from the Gateway to the backend using a standard `HTTPRoute`:

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/04-httproute.yaml
```

#### Test our MCP Backend: Kubernetes

Initialize session to get a `mcp-session-id` from the response:

```shell
curl -s -i http://localhost:8080/ai/mcps/kubernetes/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'


```

MCP Tool call to list namespaces in our cluster: 
```shell
MCP_SESSION_ID=''

curl -s http://localhost:8080/ai/mcps/kubernetes/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"namespaces_list","arguments":{}},"id":3}'
```

#### Create an MCP Backend: Github

To showcase routing to and federating different MCPs under different or the same endpoints, we will create another MCP server. We have use Github MCP

```shell

kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/github/05-mcp.yaml

# check mcp server
kubectl get pods
kubectl get sa
kubectl get service
```

Then we create a backend to route to:

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/github/06-mcp-backend.yaml

# check the backend
kubectl get agentgatewaybackend
kubectl describe agentgatewaybackend
```

Route traffic from the Gateway to the backend using a standard `HTTPRoute`:

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/07-httproute.yaml
```

#### Test our MCP Backend: Github

Initialize session to get a `mcp-session-id` from the response:

```shell
curl -s -i http://localhost:8080/ai/mcps/github/mcp \
       -H "Content-Type: application/json" \
       -H "Accept: application/json, text/event-stream" \
       -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
       -d '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'
```

MCP Tool call to search repositories: 
```shell
MCP_SESSION_ID=''

```shell
curl -s http://localhost:8080/ai/mcps/github/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"search_repositories","arguments":{"query":"kubernetes gateway-api","sort":"stars","perPage":5}},"id":3}'
```

Now we have implemented MCP routing by path using `HTTPRoute`:

`localhost:8080/ai/mcps/kubernetes` --> `kubernetes MCP service` </br>
`localhost:8080/ai/mcps/github` --> `Github MCP service` </br>

#### Federate our MCPs under one endpoint 

To federate endpoints we have to define an `AgentGatewayBackend` that combines our MCPs 

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/08-federatedbackend.yaml
```

Deploy an `HTTPRoute` to our single federated endpoint: 

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/09-httproute.yaml
```

Get session ID: 
```shell
curl -s -i http://localhost:8080/ai/mcps/mcp \
       -H "Content-Type: application/json" \
       -H "Accept: application/json, text/event-stream" \
       -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
       -d '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'
```

Now we can proceed to call tools under the single endpoint for both Kubernetes and Github MCP. 