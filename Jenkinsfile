node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    

    stage('test'){


           // Fetch the GitHub repository
            git branch: 'master', url: 'https://github.com/PapatzelosThanashs/test-app.git'
       
            // Verify Go installation
            sh 'go version'

            sh 'go mod init github.com/PapatzelosThanashs/test-app'
            
            // If using Go modules, ensure the dependencies are up to date
            sh 'go mod tidy'
        }
    

    stage('build'){
        echo "${b}"
        // Run the build process, e.g., using Maven or Gradle
        //sh 'mvn clean install'
        echo 'Build phase!'
        sh 'go run main.go'
                
    }

    stage('deploy'){
        echo "${c}"
    }
}