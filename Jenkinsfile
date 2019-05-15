pipeline {
    agent { label "SLAVE01" }
        environment {
            MY_SECRET = credentials('secret_text_test')
            AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
            AWS_SECRET_ACCESS_KEY  = credentials('AWS_SECRET_ACCESS_KEY ')
            AWS_DEFAULT_REGION = "eu-west-2"
        }
        stages {
            stage("Stage 1 - Start"){
                steps {
                    echo "Starting build process"
                    sh '''
                    ls -ltr
                    '''
                }
            }
            stage("Stage 2 - Build"){
                steps {
                    echo "Running the first proper step..."
                }
            }
            stage("Stage 3 - Exec"){
                steps {
                    echo "Now we're running!"
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
