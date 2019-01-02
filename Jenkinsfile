pipeline {
    agent any 
        stages {
            stage("Stage 1"){
                steps {
                    echo "Hello, let's begin..."
                    sh '''
                    env
                    '''
                }
            }
            stage("Stage 2"){
                steps {
                    echo "Running the first proper step..."
                    sh '''
                    ls -ltr /tmp
                    sleep 10
                    '''
                }
            }
            stage("Stage 3"){
                steps {
                    echo "Now we're running!"
                    sh '''
                    hostname
                    sleep 5
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
