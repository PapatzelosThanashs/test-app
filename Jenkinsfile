node('any'){
    def a="im into test stage"
    def b="im into build stage"
    def c="im into deploy stage"

    stage('test'){
        echo "${a} also it runs on ${env.NODE_NAME}"
    }

    stage('build'){
        echo "${b}"
    }

    stage('deploy'){
        echo "${c}"
    }
}