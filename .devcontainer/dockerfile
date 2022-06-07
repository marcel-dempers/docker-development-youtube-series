FROM alpine:latest 

ARG KIND_VERSION=0.14.0

RUN apk add --no-cache curl docker-cli

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin/kubectl

RUN curl -L https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-linux-amd64 -o /usr/local/bin/kind && \
    chmod +x /usr/local/bin/kind