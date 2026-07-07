FROM jenkins/jenkins:lts-jdk21

USER root

# Install prerequisites
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    git \
    wget \
    unzip \
    nodejs \
    npm

# Download the key directly to keyrings
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://docker.com -o /etc/apt/keyrings/docker.asc && \
    chmod a+r /etc/apt/keyrings/docker.asc

# Force cache destruction by changing the format slightly to download.docker.com
RUN rm -f /etc/apt/sources.list.d/docker.list && \
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://docker.com bookworm stable" \
  > /etc/apt/sources.list.d/docker.list

# Clean, update, and install
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get update && \
    apt-get install -y docker-ce-cli docker-compose-plugin

# Install Go
RUN curl -fsSL https://go.dev \
    | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:${PATH}"

# Ensure Jenkins user can execute Docker commands
RUN usermod -aG docker jenkins

USER jenkins
