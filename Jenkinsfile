pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }
    stages {
        stage("building") {
            steps {
                script {
                    def dockerHome = tool 'myDocker'
                    env.PATH = "${dockerHome}/bin:${env.PATH}"
                }
                sh "rm -rv *"
                sh 'git clone https://github.com/MarlonDeOcampo/golangAPI.git'
                sh "ls -1"
                sh "docker build -t alhon05/payment-service:$BUILD_ID ." 
                // sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                // sh "docker push alhon05/payment-service:$BUILD_ID ." 
            }
            // steps {
            //     sh "docker build -t alhon05/payment-service:$BUILD_ID ." 
            // }
            // stage('Login') {
            //     steps {
            //         sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            //     }
            // }
            // steps {
            //     sh "docker push alhon05/payment-service:$BUILD_ID ." 
            // }
        }
        // stage('deploying') {
        //     steps{
        //         sh "docker stack deploy -c alhon05/payment-service:$BUILD_ID main"
        //     }
        // }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}