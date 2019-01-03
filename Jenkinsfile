pipeline {
    agent { label "SLAVE01" }
        stages {
            stage("Stage 1 - Start"){
                environment {
                    MY_SECRET = credentials('secret_text_test')
                }
                steps {
                    echo "Starting build process"
                    sh '''
                    ls -ltr
                    echo "Secret password used here: ${MY_SECRET}"
                    '''
                }
            }
            stage("Stage 2 - Build"){
                steps {
                    echo "Running the first proper step..."
                    sh '''
                    go build golangchanneltest.go
                    '''
                }
            }
            stage("Stage 3 - Exec"){
                steps {
                    echo "Now we're running!"
                    sh '''
                    ./golangchanneltest "${MY_SECRET}"
                    '''
                }
            }
            stage("Stage 4"){
                steps {
                    echo "Complete."
                    sh '''
                    ls -ltr
                    '''
                }
            }
        }
}
