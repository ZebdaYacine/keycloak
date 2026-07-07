FROM jenkins/jenkins:lts-jdk21

USER root

RUN apt-get update && apt-get install -y \
    docker.io \
    docker-compose-v2 \
    git \
    curl \
    wget \
    unzip \
    nodejs \
    npm

RUN curl -fsSL https://go.dev/dl/go1.25.0.linux-amd64.tar.gz \
    | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:${PATH}"

USER jenkins