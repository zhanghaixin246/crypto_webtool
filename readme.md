crypto configtx webservice tools

# 项目文档
[项目地址](https://gitlab.trustslink.com/trustlink_blockchain/crypto_webtool) : https://gitlab.trustslink.com/trustlink_blockchain/crypto_webtool
## 1.基本介绍
基于fabric中提供的相关工具，生成可web请求完成相应功能的项目。
## 2.使用说明
```bash
# 使用 go.mod
# 安装go依赖包
go mod tidy
# 编译
go build
# 直接运行
bee run
```
- 健康检查接口
    ```shell script
       curl http://localhost:8083/
    ```
- cryptogen功能
    - ```shell script
         # generate
         # go run main.go generate --config=/Users/zhang/go/src/fabric-samples/test-network/organizations/cryptogen/crypto-config-org2.yaml --output="organizations"
         curl -X POST http://localhost:8083/cryptogen/generate \
                  -F "file=@/Users/zhang/go/src/fabric-samples/test-network/organizations/cryptogen/crypto-config-org2.yaml" \
                  -H "Content-Type: multipart/form-data"
      ```
# docker 运行
```
    docker build . -t webtool
    docker run -p 8083:8083 webtool
```

## 3.技术选型和项目架构
用`Gin`框架快速搭建基础restful风格API + `fabric`- v2.2.0
```
    ├── config          （配置包）
    ├── configtx        （生成的configtx配置文件）
    ├── core            （内核）
    ├── fabric          （fabric基础）
    ├── global          （全局对象）
    ├── go.mod
    ├── go.sum
    ├── handler         （API）
    ├── latest_log      （最新日志）
    ├── log             （日志）
    ├── main.go         （入口文件）
    ├── middleware      （中间件）
    ├── test            （测试）
    ├── tool            （fabric工具库）
    └── utils           （公共服务）
```

## 4.主要功能和计划任务
fabric的web接口

## 5.注意事项
基于版本release-2.2
