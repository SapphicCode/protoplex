String buildCommand = '''
export GOOS=$(echo ${STAGE_NAME} | cut -d '/' -f 1)
export GOARCH=$(echo ${STAGE_NAME} | cut -d '/' -f 2)
go build -o builds/protoplex_${GOOS}_${GOARCH} ./cmd/protoplex
'''

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
                stage('linux/386') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/amd64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/arm') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/arm64') {
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
