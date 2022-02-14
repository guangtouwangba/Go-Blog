pipeline{
    agent docker {
        image 'docker.io/library/alpine:latest'
    }

    triggers {
        GenericTrigger(
                genericVariables: [
                        [key: 'commit', value: '$.commits[0].id'],
                        [key: 'committer', value: '$.commits[0].committer.name'],
                        [key: 'ref', value: '$.ref']
                ],
                token: '',
                causeString: 'Triggered by github webhook on commit $commit to $ref by $committer',
                printContributedVariables: true,
                printPostContent: true,
                silentResponse: true
        )
      }



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