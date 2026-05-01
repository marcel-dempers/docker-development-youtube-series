# Introduction to LLMKube

<!-- #TODO: YouTube link -->

Website: [LLMKube](https://llmkube.com) </br>

## What is LLMKube?

LLMKube is a Kubernetes operator for self-hosted LLM inference. </br>
Source Code: [defilantech/LLMKube](https://github.com/defilantech/LLMKube) </br>
Built on [llama.cpp](https://github.com/ggml-org/llama.cpp). </br>

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name llmkube --image kindest/node:v1.34.3
```

Test our cluster and makes sure `kubectl` is configured for it:

```
kubectl get nodes
NAME                       STATUS   ROLES           AGE   VERSION
llmkube-control-plane   Ready    control-plane   40s   v1.34.3
```

## Installing LLMKube with Helm

```shell
helm repo add llmkube https://defilantech.github.io/LLMKube
helm repo update

helm search repo llmkube --versions
CHART_VERSION="0.7.5"

helm install llmkube llmkube/llmkube \
  --namespace llmkube-system \
  --version ${CHART_VERSION} \
  --create-namespace
```

Check installation:

```shell
kubectl get pods -n llmkube-system
```

Two new CRDs `Model` and `InferenceService` are now available in the cluster: 

```shell
kubectl get crds | grep llmkube
inferenceservices.inference.llmkube.dev
models.inference.llmkube.dev
```

## The CRDs

### Model

```yaml

# Deploy the Model
kubectl apply -f ai/kubernetes/llm-kube/model.yaml

# Check the Model
kubectl get model
kubectl describe model

```

### InferenceService

```yaml
# Deploy the InferenceService
kubectl apply -f ai/kubernetes/llm-kube/inference.yaml

# Check the InferenceService
kubectl get inferenceservices
kubectl describe inferenceservices
```

## Testing the Inference API

Port-forward the service to reach it from your machine:

```shell
kubectl port-forward svc/gemma-e2b-service 8080:8080
```

Send a test request to the OpenAI-compatible endpoint:

```shell
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemma-e2b",
    "messages": [
      {"role": "user", "content": "What is Kubernetes in one sentence?"}
    ],
    "max_tokens": 80
  }'
```

## OpenAI-Compatible Integration

The example below connects [Opencode](https://opencode.ai) to the local service

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "llmkube": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "LLMKube (local)",
      "options": {
        "baseURL": "http://localhost:8080/v1"
      },
      "models": {
        "gemma-e2b": {
          "name": "Gemma 4 E2B (kind)",
          "limit": {
            "context": 32768,
            "output": 4096
          }
        }
      }
    }
  }
}
```

