/*
Required Jenkins setup:
- Run this pipeline on a Linux agent with ssh, scp, ssh-keyscan, tar, node, npm, and go installed.
- Create an SSH private key credential and set its ID in SSH_CREDENTIALS_ID.
- The VPS deploy user must be allowed to restart the target systemd services with sudo.

Remote assumptions:
- The app is deployed under DEPLOY_BASE_DIR/current.
- The frontend env file already exists on the VPS at REMOTE_FRONTEND_ENV.
- The backend env file already exists on the VPS at REMOTE_BACKEND_ENV.
- The frontend systemd unit starts Next.js from DEPLOY_BASE_DIR/current/frontend.
- The backend systemd unit starts DEPLOY_BASE_DIR/current/backend/bin/echifa.
- Example systemd units are included under echifa/deploy/systemd/.
*/
pipeline {
  agent any

  options {
    disableConcurrentBuilds()
    timestamps()
    timeout(time: 45, unit: 'MINUTES')
  }

  parameters {
    string(name: 'SSH_CREDENTIALS_ID', defaultValue: 'vps-ssh', description: 'Jenkins SSH private key credential ID')
    string(name: 'VPS_HOST', defaultValue: 'your-vps-host', description: 'Target VPS hostname or IP')
    string(name: 'VPS_PORT', defaultValue: '22', description: 'SSH port on the VPS')
    string(name: 'DEPLOY_BASE_DIR', defaultValue: '/opt/echifa', description: 'Base deployment directory on the VPS')
    string(name: 'REMOTE_FRONTEND_ENV', defaultValue: '/etc/echifa/frontend.env', description: 'Frontend env file already present on the VPS')
    string(name: 'REMOTE_BACKEND_ENV', defaultValue: '/etc/echifa/backend.env', description: 'Backend env file already present on the VPS')
    string(name: 'FRONTEND_SERVICE', defaultValue: 'echifa-frontend', description: 'systemd service name for the Next.js frontend')
    string(name: 'BACKEND_SERVICE', defaultValue: 'echifa-backend', description: 'systemd service name for the Go backend')
    string(name: 'FRONTEND_HEALTH_URL', defaultValue: 'http://127.0.0.1:3000/', description: 'Frontend health check URL on the VPS')
    string(name: 'BACKEND_HEALTH_URL', defaultValue: 'http://127.0.0.1:8900/', description: 'Backend health check URL on the VPS')
  }

  environment {
    APP_DIR = 'echifa'
    FRONTEND_DIR = 'echifa/frontend'
    BACKEND_DIR = 'echifa/backend'
    ARTIFACT_FILE = 'echifa-release.tar.gz'
  }

  stages {
    stage('Validate Agent') {
      steps {
        script {
          if (!isUnix()) {
            error('This Jenkinsfile expects a Linux Jenkins agent because deployment is done over SSH to a Linux VPS.')
          }
        }

        sh '''
          set -eu
          command -v ssh
          command -v scp
          command -v ssh-keyscan
          command -v tar
          command -v node
          command -v npm
          command -v go
          node --version
          npm --version
          go version
        '''
      }
    }

    stage('Build Frontend') {
      steps {
        dir(env.FRONTEND_DIR) {
          sh '''
            set -eu
            npm ci --no-audit --no-fund
            npm run build
          '''
        }
      }
    }

    stage('Test And Build Backend') {
      steps {
        dir(env.BACKEND_DIR) {
          sh '''
            set -eu
            export GOCACHE="${WORKSPACE}/.cache/go-build"
            export GOMODCACHE="${WORKSPACE}/.cache/go-mod"
            mkdir -p "$GOCACHE" "$GOMODCACHE" bin
            go mod download
            go test ./...
            CGO_ENABLED=0 go build -buildvcs=false -o bin/echifa ./cmd/server
          '''
        }
      }
    }

    stage('Package Release') {
      steps {
        sh '''
          set -eu
          rm -f "$ARTIFACT_FILE"
          tar \
            --exclude='echifa/frontend/node_modules' \
            --exclude='echifa/frontend/.next' \
            --exclude='echifa/frontend/.env*' \
            --exclude='echifa/backend/bin' \
            --exclude='echifa/backend/.env*' \
            -czf "$ARTIFACT_FILE" \
            "$APP_DIR"
        '''

        archiveArtifacts artifacts: "${ARTIFACT_FILE}", fingerprint: true
      }
    }

    stage('Deploy To VPS') {
      when {
        expression {
          return env.CHANGE_ID == null && (env.BRANCH_NAME == null || env.BRANCH_NAME == 'main' || env.BRANCH_NAME == 'master')
        }
      }

      steps {
        withCredentials([
          sshUserPrivateKey(
            credentialsId: params.SSH_CREDENTIALS_ID,
            keyFileVariable: 'SSH_KEY',
            usernameVariable: 'SSH_USER'
          )
        ]) {
          sh '''
            set -euo pipefail

            REMOTE="${SSH_USER}@${VPS_HOST}"
            REMOTE_ARTIFACT="/tmp/${ARTIFACT_FILE%.tar.gz}-${BUILD_NUMBER}.tar.gz"

            mkdir -p "$HOME/.ssh"
            touch "$HOME/.ssh/known_hosts"
            ssh-keyscan -p "$VPS_PORT" -H "$VPS_HOST" >> "$HOME/.ssh/known_hosts"

            scp -i "$SSH_KEY" -P "$VPS_PORT" "$ARTIFACT_FILE" "$REMOTE:$REMOTE_ARTIFACT"

            ssh -i "$SSH_KEY" -p "$VPS_PORT" "$REMOTE" \
              DEPLOY_BASE_DIR="$DEPLOY_BASE_DIR" \
              RELEASE_NAME="$BUILD_NUMBER" \
              REMOTE_ARTIFACT="$REMOTE_ARTIFACT" \
              REMOTE_FRONTEND_ENV="$REMOTE_FRONTEND_ENV" \
              REMOTE_BACKEND_ENV="$REMOTE_BACKEND_ENV" \
              FRONTEND_SERVICE="$FRONTEND_SERVICE" \
              BACKEND_SERVICE="$BACKEND_SERVICE" \
              FRONTEND_HEALTH_URL="$FRONTEND_HEALTH_URL" \
              BACKEND_HEALTH_URL="$BACKEND_HEALTH_URL" \
              'bash -s' <<'REMOTE_SCRIPT'
set -euo pipefail

RELEASE_DIR="${DEPLOY_BASE_DIR}/releases/${RELEASE_NAME}"
CURRENT_LINK="${DEPLOY_BASE_DIR}/current"

command -v node >/dev/null
command -v npm >/dev/null
command -v go >/dev/null
command -v tar >/dev/null
command -v curl >/dev/null
command -v systemctl >/dev/null

if [ ! -f "$REMOTE_FRONTEND_ENV" ]; then
  echo "Missing frontend env file: $REMOTE_FRONTEND_ENV" >&2
  exit 1
fi

if [ ! -f "$REMOTE_BACKEND_ENV" ]; then
  echo "Missing backend env file: $REMOTE_BACKEND_ENV" >&2
  exit 1
fi

export GOCACHE="${DEPLOY_BASE_DIR}/.cache/go-build"
export GOMODCACHE="${DEPLOY_BASE_DIR}/.cache/go-mod"
mkdir -p "${DEPLOY_BASE_DIR}/releases" "$GOCACHE" "$GOMODCACHE"

rm -rf "${RELEASE_DIR}"
mkdir -p "${RELEASE_DIR}"

tar -xzf "$REMOTE_ARTIFACT" -C "${RELEASE_DIR}"
rm -f "$REMOTE_ARTIFACT"

cp "$REMOTE_FRONTEND_ENV" "${RELEASE_DIR}/echifa/frontend/.env.local"

cd "${RELEASE_DIR}/echifa/frontend"
npm ci --no-audit --no-fund
npm run build

cd "${RELEASE_DIR}/echifa/backend"
mkdir -p bin
CGO_ENABLED=0 go build -buildvcs=false -o bin/echifa ./cmd/server

ln -sfn "${RELEASE_DIR}/echifa" "${CURRENT_LINK}"

sudo systemctl restart "$BACKEND_SERVICE"
sudo systemctl restart "$FRONTEND_SERVICE"

sleep 5
curl -fsS "$BACKEND_HEALTH_URL" >/dev/null
curl -fsS "$FRONTEND_HEALTH_URL" >/dev/null

ls -1dt "${DEPLOY_BASE_DIR}/releases"/* 2>/dev/null | tail -n +6 | xargs -r rm -rf || true
REMOTE_SCRIPT
          '''
        }
      }
    }
  }

  post {
    success {
      echo 'Pipeline completed successfully.'
    }
    failure {
      echo 'Pipeline failed. Check the stage logs above for the failing command.'
    }
  }
}
