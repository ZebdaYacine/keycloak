FROM jenkins/jenkins:lts-jdk21

USER root

# Fixed: Changed docker-compose-v2 to docker-compose-plugin
RUN apt-get update && apt-get install -y \
    docker.io \
    docker-compose-plugin \
    git \
    curl \
    wget \
    unzip \
    nodejs \
    npm && \
    rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl -fsSL https://go.dev/dl/go1.25.0.linux-amd64.tar.gz \
    | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:${PATH}"

# Essential: Grant the jenkins user permission to run docker commands
RUN usermod -aG docker jenkins

USER jenkins
