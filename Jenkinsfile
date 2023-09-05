pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }
    stages {
        stage("building") {
            steps {
                sh "pwd"
                // sh 'git clone https://github.com/MarlonDeOcampo/golangAPI.git'
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
    // post {
    //     always {
    //         sh 'docker logout'
    //     }
    // }
}