String buildCommand = '''
export CGO_ENABLED=0

export GOOS=$(echo ${STAGE_NAME} | cut -d '/' -f 1)
export GOARCH=$(echo ${STAGE_NAME} | cut -d '/' -f 2)

go build -o builds/protoplex_${GOOS}_${GOARCH} -ldflags "-s -w" ./cmd/protoplex
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
            stages {
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
                stage('linux/mips') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/mips64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/mips64le') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/ppc64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/ppc64le') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('linux/s390x') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('windows/386') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('windows/amd64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('windows/arm') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('freebsd/386') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('freebsd/amd64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('freebsd/arm') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('netbsd/386') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('netbsd/amd64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('netbsd/arm') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('darwin/368') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('darwin/amd64') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('darwin/arm') {
                    steps {
                        sh buildCommand
                    }
                }
                stage('darwin/arm64') {
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
