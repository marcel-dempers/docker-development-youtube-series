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
CHART_VERSION="0.7.2"

helm install llmkube llmkube/llmkube \
  --namespace llmkube-system \
  --version ${CHART_VERSION} \
  --create-namespace
```

Check installation

```shell
kubectl get pods -n llmkube-system
```

Two new CRDs — `Model` and `InferenceService` — are now available in the cluster: 

```shell
kubectl get crds | grep llmkube
inferenceservices.inference.llmkube.dev
models.inference.llmkube.dev
```

## The CRDs

LLMKube introduces two custom resources that work together. </br>
`Model` describes **what** to run — where to get the model file, its format, and the hardware to target. </br>
`InferenceService` describes **how** to run it — which model to reference, replica count, resource limits, and how to expose the endpoint. </br>

### Model

```yaml

# Deploy the Model
kubectl apply -f ai/kubernetes/llm-kube/model.yaml

# Check the Model
kubectl get model
```

Key fields:

* `source` — direct URL to the GGUF file on Hugging Face; the operator downloads and caches it
* `format: gguf` — the GGUF binary format used by llama.cpp (covered in the [Introduction to llama.cpp](../../openai/README.md) guide)
* `quantization` — the compression level applied to the model weights; `Q8_0` is near-lossless quality
* `accelerator` — `cpu`, `cuda` (NVIDIA), or `metal` (Apple Silicon); for this kind demo use `cpu`

### InferenceService

```yaml
# Deploy the InferenceService
kubectl apply -f ai/kubernetes/llm-kube/inference.yaml

# Check the InferenceService
kubectl get inferenceservices

```

Key fields:

* `runtime: llamacpp` — the inference backend; LLMKube also supports `vllm` and `tgi`
* `modelRef` — links this service to the `Model` resource above; the operator watches both and reconciles state
* `replicas` — horizontal scaling is a single field change
* `memory` — rule of thumb: GGUF file size × 1.2 for headroom

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

The response format is identical to the OpenAI API — making this a drop-in replacement for any OpenAI-compatible client or SDK.

<i>Note: CPU inference in kind will produce around 5–15 tokens per second for the E2B model. This is sufficient for a demo. A real node with GPU acceleration delivers 10–17x faster throughput.</i>

## OpenAI-Compatible Integration

Because LLMKube exposes the standard OpenAI `/v1` API, it works with any OpenAI-compatible tooling. </br>
The example below connects [Opencode](https://opencode.ai) to the local service — the same pattern used when connecting directly to a llama-server:

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

The key difference from running llama.cpp manually — previously we managed `llama-server` ourselves inside Docker. </br>
Now Kubernetes manages it: if the pod crashes it restarts automatically, and scaling to more replicas is a single `kubectl patch`. </br>

## My Thoughts

LLMKube is the natural next step after running llama.cpp manually — the same inference engine, now with full Kubernetes lifecycle management. </br>

**Pros:**

* Declarative — the YAML is the source of truth, consistent with everything else in your cluster
* The operator handles restarts, health checks, and scaling with no manual intervention
* Model caching means you download once and redeploy instantly
* The same CRDs work across NVIDIA GPU, Apple Metal, and CPU — just change the `accelerator` field

**Cons:**

* Early-stage project — v0.7.x; expect API changes before a stable 1.0
* CPU inference in kind is slow — this is a learning environment, not a production setup
* GPU support requires a real node with the NVIDIA device plugin configured correctly

**Verdict:** </br>
For teams self-hosting LLMs internally, this is the right abstraction — consistent, observable, and scalable. </br>
For a solo developer on a laptop, llama.cpp or Ollama remain simpler. </br>
The [Metal Agent](https://github.com/defilantech/LLMKube/blob/main/deployment/macos/README.md) — which lets you run inference natively on Apple Silicon while orchestrating it from a Linux Kubernetes cluster — is genuinely novel and worth watching as the project matures. </br>
