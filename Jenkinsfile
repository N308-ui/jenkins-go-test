pipeline {
    agent any
    tools { go 'go-1.21' } // Use your exact tool name
    stages {
        stage('Build') {
            steps { sh 'go build -o app main.go' }
        }
    }
}
