pipeline {
    agent { label "SLAVE01" }
        stages {
            stage("Stage 1"){
                steps {
                    echo "Hello, let's begin..."
                    sh '''
                    hostname
                    ls -ltr /tmp
                    '''
                }
            }
            stage("Stage 2"){
                steps {
                    echo "Running the first proper step..."
                    sh '''
                    go build golangchanneltest.go
                    '''
                }
            }
            stage("Stage 3"){
                steps {
                    echo "Now we're running!"
                    sh '''
                    ./golangchanneltest
                    '''
                }
            }
            stage("Stage 4"){
                steps {
                    echo "I'm bored, quitting time!"
                    sh '''
                    sleep 10
                    '''
                }
            }
        }
}
