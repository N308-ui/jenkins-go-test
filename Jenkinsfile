pipeline {
    agent any

    // Step 1: Add the environment block at the top level
    environment {
        // Step 2: Define your variable
        APP_NAME = "my-go-app"
    }

    tools {
        go 'go-1.21' // Ensure this matches your Go tool name
    }

    stages {
        stage('Build') {
            steps {
                echo "Starting build for ${env.APP_NAME}..."

                // Step 3: Use double quotes to expand the variable
                // This becomes: go build -o my-go-app main.go
                sh "go build -o ${APP_NAME} main.go"
            }
        }

        stage('Verify') {
            steps {
                // You can reuse the variable in any stage
                sh "ls -lh ${APP_NAME}"
            }
        }
    }

    post {
        success {
            archiveArtifacts artifacts: "${APP_NAME}", fingerprint: true
        }
        cleanup {
            cleanWs()
        }
    }
}