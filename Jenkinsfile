pipeline {
    agent any

    stages {
        stage('Parallel Checks') {
            stages {
                stage('Static Analysis') {
                    steps {
                        echo 'Running linter...'
                    }
                }
                stage('Security Scan') {
                    steps {
                        echo 'Checking for vulnerabilities...'
                    }
                }
            }
        }

    }
}