node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"
    def goVersion = '1.21.1'
    def goUrl = "https://go.dev/dl/go${goVersion}.linux-amd64.tar.gz"
    def goTarball = "go${goVersion}.linux-amd64.tar.gz"
    def goDir = "$WORKSPACE/go"

= "$GOROOT/bin:$PATH"
    }

    stage('test'){
        // Remove old file if it exists
        sh "rm -f ${goTarball}"

        // Download Go using curl
        sh "curl -LO ${goUrl}"
        
        // Check the downloaded file
        sh "file ${goTarball}"
        
        // Extract Go to $WORKSPACE directory
        sh "tar -C ${goDir} -xzf ${goTarball}"
        
        // Set environment variables temporarily for the duration of the pipeline
        withEnv(["GOROOT=${goDir}/go", "PATH=${goDir}/go/bin:$PATH"]) {
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