@Library('my-shared-library') _ // Import the library

pipeline {
    agent any
    tools { go 'go-1.21' }
    stages {
        stage('Build') {
            steps {
                // Call your custom library function!
                standardGoBuild('my-awesome-app') 
            }
        }
    }
}
