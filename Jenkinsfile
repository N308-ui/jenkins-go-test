pipeline {
    agent any

    tools {go 'go-1.21'}

    stages {
        stage('Build') {
            steps {
                echo 'Compiling code and preparing artifacts...'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    export PATH=$PATH:$(go env GOPATH)/bin
                    go install github.com/jstemmer/go-junit-report@latest
                    go test -v ./... 2>&1 | go-junit-report > report.xml
                '''
            }
            post {
                always {
                    junit allowEmptyResults: true, testResults: '**/*.xml'
                }
            }
        }

        stage('Deploy Gate') {
            options {
                // Requirement: Abort if no action in 60 seconds
                timeout(time: 10, unit: 'SECONDS')
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

        stage ('Notify') {
            when {
                not {
                    branch 'master'
                }
            }
            steps {
                echo 'not master branch'
            }
        }

        stage('Deploy') {
            when { branch 'master' }
            steps {
                echo 'Deploying to production environment...'
            }
        }
    }

}
