FROM golang:1.23

# 设置工作目录
WORKDIR /app

# 将本地代码复制到容器中
COPY . .

# 安装依赖
RUN go mod download

# 构建项目
RUN go build -o /app/main

# 暴露端口
EXPOSE 8080

# 运行程序
CMD ["/app/main"]