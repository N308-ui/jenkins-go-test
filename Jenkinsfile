pipeline {
    agent any
    tools { go 'go-1.21' } // Ensure this matches your tool name
    stages {
        stage('Build') {
            steps {
                sh 'go build -o app main.go'
            }
        }
        stage('Test') {
            steps {
                echo 'Running Go Tests...'
                sh 'go test ./...'
            }
        }
    }
}
