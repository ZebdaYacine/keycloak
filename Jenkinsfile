def buildBackend(String service) {
    dir("${service}/backend") {
        sh """
            go mod tidy
            go test ./...
            go build -o app .
            docker build -t ${service}-backend:latest .
        """
    }
}

def buildFrontend(String service) {
    dir("${service}/frontend") {
        sh """
            npm install
            npm run build
            docker build -t ${service}-frontend:latest .
        """
    }
}

pipeline {
    agent any

    options {
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build & Test') {
            parallel {
                stage('SEGMA Backend') {
                    steps { script { buildBackend("segma") } }
                }

                stage('SEGMA Frontend') {
                    steps { script { buildFrontend("segma") } }
                }

                stage('ECHIFA Backend') {
                    steps { script { buildBackend("echifa") } }
                }

                stage('ECHIFA Frontend') {
                    steps { script { buildFrontend("echifa") } }
                }
            }
        }

        stage('Deploy') {
            steps {
                sh """
                    docker compose -f docker-compose.prod.yml down || true
                    docker compose -f docker-compose.prod.yml up -d --remove-orphans
                    docker image prune -f
                    docker ps
                """
            }
        }
    }

    post {
        success {
            echo "Deployment completed successfully."
            echo "SEGMA Frontend : http://167.86.79.16:3001"
            echo "SEGMA Backend  : http://167.86.79.16:8901"
            echo "ECHIFA Frontend: http://167.86.79.16:3002"
            echo "ECHIFA Backend : http://167.86.79.16:8902"
        }

        failure {
            echo "Pipeline failed."
        }

        always {
            cleanWs()
        }
    }
}