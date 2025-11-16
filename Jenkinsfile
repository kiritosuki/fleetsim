pipeline {
    agent any

    environment {
        // Docker Hub 登录信息，建议在 Jenkins 全局凭据里配置并使用凭据ID
        DOCKERHUB_CREDENTIALS = 'kirito-docker-hub'
        IMAGE_NAME = 'kirito693/fleetsim-backend'
        IMAGE_TAG = 'v0.1'
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
                // 使用 withCredentials 安全注入数据库 secret
                withCredentials([
                    string(credentialsId: 'db_user', variable: 'DB_USER'),
                    string(credentialsId: 'db_password', variable: 'DB_PASSWORD'),
                    string(credentialsId: 'db_host', variable: 'DB_HOST'),
                    string(credentialsId: 'db_port', variable: 'DB_PORT'),
                    string(credentialsId: 'db_name', variable: 'DB_NAME')
                ]) {
                    script {
                        // 写入 .env.prod 文件
                        sh """
                            cat > .env.prod <<EOF
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_HOST=$DB_HOST
DB_PORT=$DB_PORT
DB_NAME=$DB_NAME
EOF
                        """

                        // 创建网络（如果不存在）
                        sh '''
                            NETWORK_NAME=fleetsim-net
                            if ! docker network ls --format '{{.Name}}' | grep -w $NETWORK_NAME > /dev/null; then
                                echo "创建 Docker 网络 $NETWORK_NAME"
                                docker network create $NETWORK_NAME
                            else
                                echo "Docker 网络 $NETWORK_NAME 已存在"
                            fi
                        '''

                        // 停止并删除旧容器
                        sh '''
                            if [ $(docker ps -aq -f name=fleetsim-backend) ]; then
                                docker stop fleetsim-backend
                                docker rm fleetsim-backend
                            fi
                        '''

                        // 启动新容器，并安全注入环境变量
                        sh '''
                            echo "更新镜像"
                            docker compose pull
                            echo "构建容器"
                            docker compose up -d
                        '''

                        // 可选：构建完成后删除 .env.prod 文件
                        sh 'rm -f .env.prod'
                    }
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
