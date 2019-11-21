node('jenkins-slave') {
    
     stage('test pipeline') {
        sh(script: """
            echo "hello"
           git clone https://github.com/marcel-dempers/docker-development-youtube-series.git
           cd ./docker-development-youtube-series/golang
           
           docker build . -t test
        """)
    }
}