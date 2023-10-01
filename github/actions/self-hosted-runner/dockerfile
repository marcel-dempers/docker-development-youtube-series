FROM debian:bookworm-slim

ARG RUNNER_VERSION="2.302.1"

ENV GITHUB_PERSONAL_TOKEN ""
ENV GITHUB_OWNER ""
ENV GITHUB_REPOSITORY ""

# Install Docker -> https://docs.docker.com/engine/install/debian/

# Add Docker's official GPG key:
RUN apt-get update && \
    apt-get install -y ca-certificates curl gnupg
RUN install -m 0755 -d /etc/apt/keyrings
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
RUN chmod a+r /etc/apt/keyrings/docker.gpg

# Add the repository to Apt sources:
RUN echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt-get update

# I only install the CLI, we will run docker in another container!
RUN apt-get install -y docker-ce-cli

# Install the GitHub Actions Runner 
RUN apt-get update && apt-get install -y sudo jq

RUN useradd -m github && \
  usermod -aG sudo github && \
  echo "%sudo ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER github
WORKDIR /actions-runner
RUN curl -Ls https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz | tar xz \
  && sudo ./bin/installdependencies.sh

COPY --chown=github:github entrypoint.sh  /actions-runner/entrypoint.sh
RUN sudo chmod u+x /actions-runner/entrypoint.sh

#working folder for the runner 
RUN sudo mkdir /work 

ENTRYPOINT ["/actions-runner/entrypoint.sh"]