# 第一阶段：构建可执行文件
FROM golang:1.22-alpine AS builder

WORKDIR /app

# 拷贝 go.mod 和 go.sum 并下载依赖
COPY go.mod ./
RUN go mod download

# 拷贝项目源码
COPY . .

# 构建可执行文件（假设 main.go 是入口）
RUN go build -o app main.go

# 第二阶段：创建更小的运行镜像
FROM alpine:latest

WORKDIR /app

# 拷贝可执行文件
COPY --from=builder /app/app .

# 如果有需要的配置文件或静态资源，也COPY进来
# COPY --from=builder /app/xxx ./xxx

# 启动应用
ENTRYPOINT ["./app"] 