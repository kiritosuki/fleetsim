pipeline {
    agent any

    environment {
        // Docker Hub 登录信息，建议在 Jenkins 全局凭据里配置并使用凭据ID
        DOCKERHUB_CREDENTIALS = 'kirito-docker-hub'
        IMAGE_NAME = 'kirito693/fleetsim-backend'
        IMAGE_TAG = 'v0.1'

        // 数据库环境变量，从 Jenkins Secret Text 注入
        DB_USER = credentials('db_user')
        DB_PASSWORD = credentials('db_password')
        DB_HOST = credentials('db_host')
        DB_PORT = credentials('db_port')
        DB_NAME = credentials('db_name')
    }

    options {
        skipDefaultCheckout(true)
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
                    // 在当前会话中禁用buildkit
                    sh '''
                        export DOCKER_BUILDKIT=0
                        echo BuildKit disabled: $DOCKER_BUILDKIT
                        docker build --network host -t ${IMAGE_NAME}:${IMAGE_TAG} .
                    '''
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
                        sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_USER --password-stdin'
                                                    }
                    // 推送镜像
                    sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
                }
            }
        }

        stage('Deploy Container') {
            steps {
                script {
                    // 读取环境变量配置并写入.env.prod文件
                    sh """
                    cat > .env.prod <<EOF
                    DB_USER=${DB_USER}
                    DB_PASSWORD=${DB_PASSWORD}
                    DB_HOST=${DB_HOST}
                    DB_PORT=${DB_PORT}
                    DB_NAME=${DB_NAME}
                    IMAGE_NAME=${IMAGE_NAME}
                    IMAGE_TAG=${IMAGE_TAG}
                    EOF
                    """

                    sh '''
                        NETWORK_NAME=fleetsim-net
                        # 判断网络是否存在
                        if ! docker network ls --format '{{.Name}}' | grep -w $NETWORK_NAME > /dev/null; then
                            echo "创建 Docker 网络 $NETWORK_NAME"
                            docker network create $NETWORK_NAME
                        else
                            echo "Docker 网络 $NETWORK_NAME 已存在"
                        fi
                    '''

                    // 先停止并删除旧容器（如果存在）
                    sh '''
                        if [ $(docker ps -aq -f name=fleetsim-backend) ]; then
                            docker stop fleetsim-backend
                            docker rm fleetsim-backend
                        fi
                    '''

                    // 启动新容器并注入数据库环境变量
                    sh '''
                        echo "更新镜像"
                        docker compose pull
                        echo "构建容器"
                        docker compose up -d
                    '''
                }
            }
        }
    }

    post {
        success {
            echo "Docker image built and pushed successfully: ${IMAGE_NAME}:${IMAGE_TAG}"
        }
        failure {
            echo 'Build or push failed!'
        }
    }
}
