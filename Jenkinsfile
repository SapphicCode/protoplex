pipeline {
    agent {
        label 'go'
    }
    stages {
        stage('Prepare') {
            checkout scm
            sh 'go get -u -v .'
        }
        stage('Build') {
            sh 'go build protoplex.go'
        }
        stage('Cleanup') {
            archiveArtifacts 'protoplex'
        }
    }
}
