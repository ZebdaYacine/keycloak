FROM jenkins/jenkins:lts-jdk21

USER root

# Install prerequisites for adding external repositories
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    git \
    wget \
    unzip \
    nodejs \
    npm

# Add Docker's official GPG key
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://docker.com | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
    chmod a+r /etc/apt/keyrings/docker.gpg

# Add the Docker apt repository
RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://docker.com \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null

# Now install the correct docker-ce-cli and docker-compose-plugin
RUN apt-get update && apt-get install -y \
    docker-ce-cli \
    docker-compose-plugin \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl -fsSL https://go.dev \
    | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:${PATH}"

# Ensure Jenkins user can execute Docker commands
RUN usermod -aG docker jenkins

USER jenkins
