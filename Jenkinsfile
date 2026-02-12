pipeline {
    agent any

    stages {
        stage ('Parallel Checks') {
            parallel {
                stage ('Static Analysis') {
                    echo 'Running linter...'
                }
                stage ('Security Scan') {
                    echo 'Checking for vulnerabilities...'
                }
            }
        }
    }
}