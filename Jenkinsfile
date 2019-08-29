String buildCommand = 'go build -o builds/protoplex_${GOOS}_${GOARCH} ./cmd/protoplex'

pipeline {
    agent {
        label 'go'
    }
    stages {
        stage('Prepare') {
            steps{
                checkout scm
                sh 'go get -u -v .'
                sh 'mkdir -p builds'
            }
        }
        stage('Build') {
            parallel {
                stage('linux/amd64') {
                    String go_os = env.STAGE_NAME.split('/')[0]
                    String go_arch = env.STAGE_NAME.split('/')[1]
                    environment {
                        GOOS = go_os
                        GOARCH = go_arch
                    }
                    steps {
                        sh buildCommand
                    }
                }
            }
        }
        stage('Cleanup') {
            steps {
                archiveArtifacts 'builds/*'
            }
        }
    }
}
