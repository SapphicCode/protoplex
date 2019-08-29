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
            steps {
                sh 'go build protoplex.go'
            }
        }
        stage('Cleanup') {
            steps {
                archiveArtifacts 'protoplex'
            }
        }
    }
}
