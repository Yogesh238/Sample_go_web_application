pipeline {
    agent any
    stages {
    stage('Checkout Source') {
        steps {
                git branch: 'main',
                credentialsId: "yogesh-github",
                url: "https://github.com/Yogesh238/Sample_go_web_application.git"
        }
    }
        stage('Compile & Build') {
            steps {
	         //sh '/usr/local/go/bin/go build main.go'
		       sh '/usr/bin/go build main.go'
            }
        }
	    stage('sonarqube analysis'){
                    steps {
           withSonarQubeEnv(credentialsId: 'sonarqube', installationName: 'sonar')
        {
              sh '/opt/sonar-scanner/bin/sonar-scanner \
  -Dsonar.projectKey=demo-test \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://52.204.215.151:9000/'
           
    }
            }
            }	    
    }
}
