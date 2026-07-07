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

# Streamlined: Download the key directly to keyrings without piping to gpg
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://docker.com -o /etc/apt/keyrings/docker.asc && \
    chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository referencing our .asc file
RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://docker.com \
  bookworm stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install the correct docker-ce-cli and docker-compose-plugin
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
