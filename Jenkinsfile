pipeline {
    agent any

    stages {
        stage ('build') {
            steps {
                sh 'touch temp.txt'
            }
        }
    }
    post {
        always {
            sh 'rm -rf temp.txt'
        }
        failure {
            echo 'Oh no! Something went wrong'
        }
    }

}