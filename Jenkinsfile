node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    

    stage('test'){

            // Temporarily set the PATH environment variable
           sh 'PATH=$PATH:/home/jenkins/go/bin'
           sh 'sleep 120'
       
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