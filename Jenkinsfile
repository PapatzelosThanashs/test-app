node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    

    stage('test'){

            sh 'sleep 600'
       
            // Verify Go installation
            sh 'go version'
            
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