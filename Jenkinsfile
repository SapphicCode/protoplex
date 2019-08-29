pipeline {
    agent {
        label 'go'
    }
    stages {
        stage('Prepare') {
            steps{
                checkout scm
                sh 'go get -u -v .'
            }
        }
        stage('Build') {
            environment {
                GOOS = 'linux'
                GOARCH = 'amd64'
            }
            steps {
                sh 'go build -o /tmp/protoplex_${GOOS}_${GOARCH} ./cmd/protoplex'
            }
        }
        stage('Cleanup') {
            steps {
                archiveArtifacts '/tmp/protoplex_*'
            }
        }
    }
}
