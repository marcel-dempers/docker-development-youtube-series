# Introduction to llama.cpp

<!-- #TODO: YouTube link -->

In this guide, we will learn how to run large language models locally using llama.cpp </br>

We will cover:
* What is llama.cpp
* What is quantization
* How to find and run models
* How to host a local model server 
* Interact with models using CLI like OpenCode.

## What is llama.cpp

[llama.cpp](https://llama-cpp.com/) is a lightweight, high-performance implementation designed to run large language models locally on your own machine. </br>
It has no external dependencies and no Python — just pure C++. </br>

Key highlights:

* Runs entirely on your own hardware — no cloud, no API keys required
* Supports CPU, NVIDIA GPU (CUDA), AMD GPU (ROCm), Intel GPU (SYCL), and Vulkan
* Natively supports quantized models (Q2 through Q8) for memory efficiency
* Ships with a CLI (`llama-cli`) and a local HTTP server (`llama-server`)

## Quantization

* AI models are massive
  * a 7B parameter model at full precision (F32) needs ~28GB of RAM
  * a 31B parameter model needs a really good top of the range GPU like a 4090 
* Quantization compresses the model's numbers (e.g. 32-bit down to 4-bit), cutting memory by up to 8x with minimal quality loss
* A 7B model can run in as little as 4GB of RAM on a laptop CPU
* llama.cpp natively supports Q2 through Q8, letting you tune the speed, memory, and quality trade-off to your hardware

Quantization levels:

| Quant | Size | Quality | Use when |
|---|---|---|---|
| `Q2_K` | Smallest | Lowest | Very limited RAM |
| `Q3_K_M` | Very small | Low | Tight on RAM |
| `Q4_K_M` | Small | Good | **Best balance — most popular** |
| `Q5_K_M` | Medium | Better | Have spare RAM |
| `Q6_K` | Large | Very good | Plenty of RAM |
| `Q8_0` | Largest | Near lossless | Maximum quality |
| `F16` | Huge | Lossless | Rarely practical |

## Where to Get Models

[Hugging Face](https://huggingface.co/models) is the primary hub for AI models — think of it as the Docker Hub of models. </br>

We can start with [Gemma 4](https://huggingface.co/collections/ggml-org/gemma-4) </br>

## Source Code

llama.cpp is fully open source and available on [GitHub](https://github.com/ggml-org/llama.cpp). </br>
Pre-built binary releases for all platforms are available at [github.com/ggml-org/llama.cpp/releases](https://github.com/ggml-org/llama.cpp/releases).

## Docker Images

The easiest way to get started is via the official Docker images. </br>
[Official Docker Guide](https://github.com/ggml-org/llama.cpp/blob/master/docs/docker.md) </br>

Three image types are available:

* `light`: includes CLI and completion tools
* `full`: includes light + tools for model conversion (safetensors → GGUF)
* `server`: only the `llama-server` executable

### Choosing the right image

| Image | Backend | Vendor |
|---|---|---|
| `ghcr.io/ggml-org/llama.cpp:full` | CPU | Any |
| `ghcr.io/ggml-org/llama.cpp:full-cuda` | CUDA 12 + cuBLAS | **NVIDIA** |
| `ghcr.io/ggml-org/llama.cpp:full-cuda13` | CUDA 13 + cuBLAS | **NVIDIA** |
| `ghcr.io/ggml-org/llama.cpp:full-rocm` | ROCm + hipBLAS | **AMD** |
| `ghcr.io/ggml-org/llama.cpp:full-intel` | SYCL | **Intel** Arc/Xe |
| `ghcr.io/ggml-org/llama.cpp:full-vulkan` | Vulkan | Cross-vendor |
| `ghcr.io/ggml-org/llama.cpp:full-openvino` | OpenVINO | **Intel** CPU/iGPU/NPU |

## Running llama.cpp

Example 1: Linux (CPU or integrated graphics): </br>
Using `--device /dev/dri` for integrated Graphics

```shell
docker run -it \
  -v ~/models:/models \
  --device /dev/dri \
  -p 8080:8080 \
  --entrypoint bash ghcr.io/ggml-org/llama.cpp:full-intel
```

Example 2:

WSL with NVIDIA GPU:

```shell
docker run -it \
  -v ~/models:/models \
  --gpus all \
  -p 8080:8080 \
  --entrypoint bash ghcr.io/ggml-org/llama.cpp:full-cuda
```

Once inside the container, verify the CLI works with `./llama-cli --help`

## Running a Model

### Download and run directly from Hugging Face

The `-hf` flag lets you pull a model directly from Hugging Face by repo and quant:

```shell
./llama-cli -hf ggml-org/gemma-4-E4B-it-GGUF:Q4_K_M
```

<b>Note:</b> This downloads and caches the model under `~/.cache/huggingface/` inside the container, not your `/models` mount. To persist models across container restarts, download them manually to `/models` first. </br>

### Run from a local GGUF file

```shell
./llama-cli -m /models/gemma-4-E4B-it-Q4_K_M.gguf -ngl 99
```

Use `-ngl 0` to disable GPU offload and run on CPU only:

```shell
./llama-cli -m /models/gemma-4-E4B-it-Q4_K_M.gguf -ngl 0
```

## Hosting a Model

To expose your model as an HTTP API, use `llama-server` on port `8080`. </br>
Once running, you can access the built-in web UI and API at [http://localhost:8080](http://localhost:8080). </br>

Linux (CPU):

```shell
export LLAMA_ARG_HOST=0.0.0.0
./llama-server -m /models/gemma-4-E4B-it-Q4_K_M.gguf \
  --port 8080 \
  -ngl 0 \
  --jinja \
  -c 8192 \
  --parallel 1 \
  --temperature 1.0 \
  --top-p 0.95 \
  --top-k 64
```

WSL with NVIDIA GPU:

```shell
export LLAMA_ARG_HOST=0.0.0.0

# gemma-4-31B-it-Q4_K_M.gguf
# gemma-4-26B-A4B-it-Q4_K_M.gguf
# gemma-4-E2B-it-Q8_0.gguf
# gemma-4-E4B-it-Q4_K_M.gguf

./llama-server -m /models/gemma-4-E2B-it-Q8_0.gguf \
  --port 8080 \
  -ngl 99 \
  --jinja \
  -c 131072 \
  --parallel 1 \
  --temperature 1.0 \
  --top-p  0.95 \
  --top-k 64
```

Key flags:

| Flag | Purpose |
|---|---|
| `-ngl 99` | Offload all layers to GPU (llama.cpp uses as many as fit in VRAM) |
| `-c 65536` | Context window size in tokens |
| `--parallel 1` | 1 concurrent request slot — correct for single-user use |
| `--jinja` | Enables Jinja2 chat templates for proper instruction-tuned model formatting |
| `--temperature` | Default sampling temperature (0.0 = deterministic, 1.0 = creative) |

## Model Interaction via OpenCode

You can interact with your locally hosted model using [OpenCode](https://opencode.ai/docs/providers/#llamacpp). </br>
OpenCode connects to the `llama-server` endpoint over its OpenAI-compatible API. </br>

Configure `~/.opencode/opencode.json`:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "llama.cpp": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "llama-server (local)",
      "options": {
        "baseURL": "http://localhost:8080/v1"
      },
      "models": {
        "gemma-4-26B-A4B-it-Q4_K_M.gguf": {
          "name": "gemma-4-26B-A4B-it-Q4_K_M (local)",
          "limit": {
            "context": 65536,
            "output": 16384
          }
        }
      }
    }
  }
}
```