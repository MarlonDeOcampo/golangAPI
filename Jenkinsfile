pipeline {
    
    agent any

    stages {
        stage("cloning") {
            steps {
                sh 'git pull https://github.com/MarlonDeOcampo/golangAPI.git'
            }
            echo "successfully pulled the repo"
        }
        stage("building image") {
            steps {
                sh docker build -t "alhon05/payment-service:$BUILD_ID" .
            }
            echo "successfully built the image"
        }
        stage("pushing image") {
            steps {
                sh docker push "alhon05/payment-service:$BUILD_ID" .
            }
            echo "successfully pushed the image"
        }
        stage('Deploy Image') {
            steps{
                sh docker stack deploy -c "alhon05/payment-service:$BUILD_ID" main
            }
            echo "successfully re-deployed the image"
        }
    }
}