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
		       sh 'go build main.go'
            }
        }
    }
}
