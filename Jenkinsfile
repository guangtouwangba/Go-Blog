pipeline{
    agent any

    stages{
        stage('Checkout SCM') {
          steps {
            checkout([
              $class: 'GitSCM',
              branches: [[name: 'master']],
              userRemoteConfigs: [[
                url: 'git@github.com:819110812/Go-Blog.git',
                credentialsId: '',
              ]]
             ])
           }
        }
        stage('Build'){
            steps{
                sh 'echo "Hello World"'
            }
        }

        stage('Test'){
            steps{
                sh 'echo "Hello World"'
            }
        }
        stage('Deploy'){
            steps{
                sh 'echo "Hello World"'
            }
        }
    }
}