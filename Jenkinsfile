pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }
    stages {
        stage("building") {
            steps {
                sh "rm -rv *"
                git 'https://github.com/MarlonDeOcampo/golangAPI.git'
                sh "pwd"
                sh "ls -l"
            }
        }
    }
}
