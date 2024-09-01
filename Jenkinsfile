pipeline {
    agent any
    tools { go 'Go 1.22.5' }
    stages {
       
        stage('Checkout') {
            steps {
                // Checkout the source code from version control
                  echo 'SCM succeed'
            }
        }
        
        stage('Build') {
            steps {
                // Run the build process, e.g., using Maven or Gradle
                //sh 'mvn clean install'
                 echo 'Build phase!'
                
            }
        }
        
        stage('Test') {
            steps {
                // Run tests
                //sh 'mvn test'
                 echo 'Test phase!'
            }
        }
        
        stage('Deploy') {
            steps {
                // Deploy the application, e.g., copy files to a server
                //sh 'scp target/your-app.jar user@server:/path/to/deploy'
                 echo 'Deploy phase!'
            }
        }
    }
    
    post {
        always {
            // Actions that should always happen, like cleanup
            cleanWs()
        }
        
        success {
            // Actions to perform on successful build
            echo 'Build succeeded!'
        }
        
        failure {
            // Actions to perform on failed build
            echo 'Build failed!'
        }
    }
}
