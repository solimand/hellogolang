/* Requires the Docker Pipeline plugin */
pipeline {
    //agent { docker { image 'golang:1.19.1-alpine' } }
    agent any
    /*ISSUE https://issues.jenkins.io/browse/JENKINS-60473
        no container lnx under windows */
    stages {
        stage('build') {
            when {
                expression {
                    env.BRANCH_NAME == 'release'
                }
            }
            steps {
                powershell 'go version'
                //sh 'go version'
                script{ /*build the Docker image*/
                    docker.build("myfibo", "Functions/")
                }                
            }
        }
    }
}