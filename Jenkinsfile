node{
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    

    stage('test'){


           // Fetch the GitHub repository
            git branch: 'master', url: 'https://github.com/PapatzelosThanashs/test-app.git'
       
           //sh 'kubectl get pods -n jenkins'
        }
    

    stage('build'){
        echo "${b}"
        // Run the build process, e.g., using Maven or Gradle
        //sh 'mvn clean install'
        echo 'Build phase!'
       
                
    }

    stage('deploy'){
        echo "${c}"
    }
}
