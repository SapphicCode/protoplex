pipeline {
    agent {
        label 'go'
    }
    stages {
        stage('Prepare') {
            steps {
                checkout scm
                sh 'go get github.com/mitchellh/gox'
                sh 'go get -u -v .'
                sh 'mkdir -p builds'
            }
        }
        stage('Build') {
            steps {
                sh 'gox -ldflags="-s -w" -output="builds/{{.Dir}}_{{.OS}}_{{.Arch}}" ./cmd/protoplex'
            }
        }
        stage('Cleanup') {
            steps {
                archiveArtifacts 'builds/*'
            }
        }
    }
}
