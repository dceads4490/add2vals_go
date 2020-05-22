pipeline {
    agent none
    options {
        skipStagesAfterUnstable()
    }
    stages {
        stage('Build') {
            agent any
            steps {
                sh '$(pwd)/build.sh'
                stash(name: 'compiled-results', includes: 'add2vals')
            }
        }
        stage('Test') {
            agent any
            steps {
                sh '/usr/local/go/bin/go test'
            }
        }
        stage('Deliver') {
            agent any
            environment {
                VOLUME = '$(pwd)/sources:/src'
            }
            steps {
                dir(path: env.BUILD_ID) {
                    unstash(name: 'compiled-results')
                }
            }
            post {
                success {
                    archiveArtifacts "${env.BUILD_ID}/add2vals"
                }
            }
        }
    }
}
