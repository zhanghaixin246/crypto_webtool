FROM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

# 编译
RUN go build -o webtool .

###################
# 创建一个小镜像
###################
FROM scratch

COPY --from=builder /build/config /config
COPY --from=builder /build/configtx /configtx
# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/webtool /

# 运行
ENTRYPOINT ["/webtool"]