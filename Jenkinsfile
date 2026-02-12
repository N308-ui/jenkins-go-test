pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Compiling code and preparing artifacts...'
            }
        }

        stage('Test') {
            steps {
                echo 'Running unit tests...'
            }
        }

        stage('Deploy Gate') {
            options {
                // Requirement: Abort if no action in 60 seconds
                timeout(time: 60, unit: 'SECONDS')
            }
            input {
                // Requirement: Custom button message
                message "Proceed with Deployment?"
                ok "Approve"
            }
            steps {
                echo "Gate cleared!"
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying to production environment...'
            }
        }
    }
}
