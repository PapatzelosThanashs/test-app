node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    environment {
        GOROOT = "/usr/local/go"
        GOPATH = "$WORKSPACE/go"
        PATH = "$GOROOT/bin:$PATH"
    }

    stage('test'){
        echo "${a} also it runs on ${env.NODE_NAME}"
        //install go
        sh 'wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz'
        sh 'tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz'



         // Ensure Go is installed and available
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