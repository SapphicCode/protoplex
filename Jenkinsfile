pipeline {
    options {
        buildDiscarder logRotator(artifactNumToKeepStr: '10')
    }
    agent {
        label 'go'
    }
    stages {
        stage('Prepare') {
            steps {
                checkout scm
                sh 'go get github.com/mitchellh/gox'
                sh 'go get -d -u -v ./cmd/protoplex'
                sh 'mkdir -p builds'
            }
        }
        stage('Build') {
            environment {
                CGO_ENABLED = '0'
            }
            steps {
                sh '''
                    version="$(git describe --tags --abbrev=0 HEAD || true)"
                    if [ -z "${version}" ]; then
                        version="v0.0.0"
                    fi
                    build="$(git rev-parse --short HEAD)"
                    fullver="${version}+${build}"
                    gox -parallel=2 -ldflags="-s -w -X main.version=${fullver}" -output="builds/{{.Dir}}_{{.OS}}_{{.Arch}}" ./cmd/protoplex
                '''
            }
        }
        stage('Cleanup') {
            steps {
                archiveArtifacts 'builds/*'
            }
        }
    }
}
