# Introduction to MCP

<!-- #TODO: YouTube link MCP -->

## Running an MCP

Start our simple portable sandbox in a linux container. </br>
We'll expose a port so we can showcase `HTTP` calls as well. </br>

```shell
docker run -it --rm \
  -p 8080:8080 \
  ubuntu:latest bash

# install some tools we'll use:
apt-get update && apt-get install -y curl tree git jq tar

#install nodejs and python as a requirement for many MCP servers
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.4/install.sh | bash
# in lieu of restarting the shell
\. "$HOME/.nvm/nvm.sh"
nvm install 24
node -v

apt-get install -y python3 python3-pip python3-venv# Create the environment
python3 -m venv /opt/venv
source /opt/venv/bin/activate
```

### Example 1: STDIO (File & GIT MCP Demo)

Firstly, let's look at a `git` MCP server. </br>
To do this we need a basic `git` repo, lets create one:

```shell
mkdir -p /workspace && cd /workspace
echo "# Hello MCP" > README.md
echo "console.log('hello')" > app.js
mkdir src && echo "export default {}" > src/index.js
tree

git init
git config user.email "demo@mcp.dev"
git config user.name "MCP Demo"
git add .
git commit -m "initial commit"

# leave an uncommitted change for the diff demo
echo "## Updated" >> README.md
```


#### File MCP Example:

Use the MCP in three steps:

```shell
(
  echo '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'
  echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"list_directory","arguments":{"path":"/workspace"}},"id":2}'
  echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"read_text_file","arguments":{"path":"/workspace/README.md"}},"id":3}'
  sleep 2
) | npx -y @modelcontextprotocol/server-filesystem /workspace 2>/dev/null | jq -c .
```

#### GIT MCP Example:

Setup `git` MCP:

```shell
pip3 install mcp-server-git
```

Use the MCP in three steps:

```shell
(
  echo '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'
  echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"git_log","arguments":{"repo_path":"/workspace","max_count":3}},"id":2}'
  echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"git_diff_unstaged","arguments":{"repo_path":"/workspace"}},"id":3}'
  sleep 2
) | python3 -u -m mcp_server_git --repository /workspace 2>/dev/null | jq -c .
```

### Example 2: HTTP (GitHub & K8s MCP Demo)

In this example, we move from stdio to HTTP which is basically just MCP hosted over a Web Endpoint. </br>

#### Github MCP Example:

There are official MCP servers available online. The [Official GitHub MCP server](https://github.com/github/github-mcp-server) is what we will be using </br>

Let's install it:

```shell
curl -fsSL https://github.com/github/github-mcp-server/releases/download/v1.0.5/github-mcp-server_Linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin github-mcp-server
github-mcp-server --help
```

Start it:

```shell 
github-mcp-server http --port 8080
```

We need to provide authentication with Github:

```shell
export GITHUB_PERSONAL_ACCESS_TOKEN=''
```
Call it in three steps:

```shell
# step-1: initialise
curl -s -i http://localhost:8080/mcp \
       -H "Content-Type: application/json" \
       -H "Accept: application/json, text/event-stream" \
       -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
       -d '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'

MCP_SESSION_ID=''

# step-2: list tools
curl -s http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/list","id":2}'

# step-3: tool call
curl -s http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Authorization: Bearer ${GITHUB_PERSONAL_ACCESS_TOKEN}" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"search_repositories","arguments":{"query":"kubernetes gateway-api","sort":"stars","perPage":5}},"id":3}'
```

#### Kubernetes MCP Example:

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```shell
kind create cluster --name mcp --image kindest/node:v1.34.0
```

In this example we will use a Kubernetes MCP and setup a read-only `ServiceAccount` so it can read our cluster. </br>

```shell
kubectl apply -f kubernetes/gateway-api/agentgateway/mcp/kubernetes/02-mcp.yaml

# check mcp server
kubectl get pods
kubectl get sa
kubectl get service

# forward ports
kubectl port-forward svc/mcp-k8s 8080
```

Call it in three steps:

```shell 
# step-1: initialise
curl -s -i http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}},"id":1}'

MCP_SESSION_ID=''

# step-2: list tools
curl -s http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/list","id":2}'

# step-3: tool call
curl -s http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -H "Mcp-Session-Id: ${MCP_SESSION_ID}" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"namespaces_list","arguments":{}},"id":3}'
```