# =========================
# *1* 编译阶段
# =========================

# 1. 指定镜像源 命名 builder 阶段
FROM golang:1.25-alpine AS builder

# 2. 指定编译环境 写入环境变量
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# 3. 创建工作目录
WORKDIR /app

# 4. 先复制依赖文件并下载依赖（为了提高缓存命中 原则上 go build 会自动下载依赖）
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

# 5. 复制完整项目并编译
COPY . ./
RUN go build -o fleetsim-backend ./cmd/main.go

# =========================
# *2* 运行阶段
# =========================

# 1. 指定镜像源 命名 runtime 阶段
FROM alpine:latest AS runtime

# 2. 创建工作目录
WORKDIR /app

# 3. copy builder阶段的产生的二进制文件
COPY --from=builder /app/fleetsim-backend ./

# 4. 添加执行权限
RUN chmod +x fleetsim-backend

# 5. 暴露端口（后端服务运行在8088）
EXPOSE 8088

# 6. 运行服务
CMD ["./fleetsim-backend"]