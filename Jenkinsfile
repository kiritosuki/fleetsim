pipeline {
    agent any

    environment {
        // Docker Hub 登录信息，建议在 Jenkins 全局凭据里配置并使用凭据ID
        DOCKERHUB_CREDENTIALS = 'kirito-docker-hub'
        IMAGE_NAME = 'kirito693/fleetsim-backend'
        IMAGE_TAG = "v0.1"
    }

    stages {
        stage('Checkout') {
            steps {
                // 拉取代码
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // 构建镜像
                    sh "docker build --network host -t ${IMAGE_NAME}:${IMAGE_TAG} ."
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    // 登录 Docker Hub
                    withCredentials([usernamePassword(credentialsId: DOCKERHUB_CREDENTIALS,
                                                    usernameVariable: 'DOCKER_USER',
                                                    passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USER --password-stdin"
                    }
                    // 推送镜像
                    sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
                }
            }
        }
    }

    post {
        success {
            echo "Docker image built and pushed successfully: ${IMAGE_NAME}:${IMAGE_TAG}"
        }
        failure {
            echo "Build or push failed!"
        }
    }
}
