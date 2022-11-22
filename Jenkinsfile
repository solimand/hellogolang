/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.19.1-alpine' } }
    stages {
        stage('build') {
            when {
                expression {
                    env.BRANCH_NAME == 'release'
                }
            }
            steps {
                sh 'go version'
            }
        }
    }
}