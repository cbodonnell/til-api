pipeline {
    agent any
    environment {
        GOROOT = "${tool type: 'go', name: 'go1.15.6'}/go"
    }
    stages {
        stage('build') {
            steps {
                echo 'building...'
                sh 'echo $GOROOT'
                sh '$GOROOT/bin/go build'
            }
        }
        stage('test') {
            steps {
                echo 'testing...'
            }
        }
        stage('deploy') {
            steps {
                echo 'deploying...'
                sh 'sudo systemctl stop til-api'
                sh 'sudo cp til-api /etc/til-api/til-api'
                sh 'sudo systemctl start til-api'
            }
        }
    }
    post {
        cleanup {
            deleteDir()
        }
    }
}