pipeline {
    
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }

    stages {
        stage('checkout') {
            steps {
               checkout scmGit(branches: [[name: '*/dev']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/MarlonDeOcampo/golangAPI.git']])
            }
        }
        stage('build image') {
            steps {
                script {
                    sh "docker build -t alhon05/payment-service:v$BUILD_ID ."
                }
            }
        }
        stage('Login to Docker Hub') {      	
            steps{                
                    sh 'docker --version'
                    sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'                		
                    echo 'Login Completed'      
                }           
        }
        stage('upload image') {
            steps {
                sh "docker push alhon05/payment-service:v$BUILD_ID"
            }
        }
        stage('remote host redeploy') {
            steps {            
                script {
                    def remote = [:]
                    remote.name = "tns-des145"
                    remote.host = "172.24.31.39"
                    remote.allowAnyHosts = true
                    withCredentials([sshUserPrivateKey(credentialsId: 'jenkins-to-host', keyFileVariable: 'identity', passphraseVariable: 'jenkins', usernameVariable: 'tns-des145')]) {
                        remote.user = tns-des145
                        remote.identityFile = identity
                        sshCommand remote: remote, command: "pwd"
                    }
                }
            }
        }
    }
    post {
        always {
            sh 'docker logout'
            echo 'Job Completed...'   
        }
    }
}


           
	