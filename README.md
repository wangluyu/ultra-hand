# Ultra Hand | *究极手*

一个易用的微服务项目骨架。 可以让开发人员专注业务逻辑开发。

使用一个命令即可生成一个微服务的代码。

单个服务可以独立部署，服务间随意调用。

## 目录结构

```tree
.
├── cmd                       # 入口指令，用于启动微服务，每个微服务（app）一个文件夹
│   └── ${APP_NAME}
│       └── main.go           # Application 启动入口文件
├── app                   # Application 目录
│    └── ${APP_NAME}
│       ├── transport.go      # 传输层，必须。定义请求 / 响应数据结构, 解析用户请求 / 响应。定义路由，传输域绑定到具体的传输协议，如 HTTP 或 gRPC。
│       ├── endpoint.go       # 端点层，必须。
│       └── service.go        # 服务层，必须。串联业务逻辑。负责调用 logic 层。
├── entity                    # 贯穿整个 server 的通用的数据结构实体定义层，比如数据结构，错误码或者其他各种常量定义
├── logic                     # 业务封装层，当前 server 的业务复杂逻辑。
├── pkg                       # server 通用的公共模块
├── job                       # 定时任务
├── resource                  # 静态资源文件。这些文件可以通过 资源打包/镜像编译 的形式注入到发布文件中
├── script                    # 脚本文件。
├── manifest                  # 交付清单， 包含程序编译、部署、运行、配置的文件
│   ├── config                # 配置文件
│   ├── deploy                # 部署相关的文件
│   ├── docker                # Docker 镜像相关依赖文件，脚本文件等等
│   └── protobuf              # GRPC 协议时使用的 protobuf 协议定义文件，协议文件编译后生成 go 文件到 service/${SERVICE_NAME}/pb 目录
├── doc                       # 文档目录
├── go.mod
├── go.sum
└── README.md
```

### 各层概念介绍

一个 Server 包含多个 Application，每个 Application 都可以单独启动。

1. cmd
   cmd 层负责引导程序启动，显著的工作是初始化逻辑、启动server监听、阻塞运行程序直至server退出。
2. app
3. entity
4. logic
5. pkg
6. job
7. resource
8. script
9. manifest

> 参考链接
>
>> [1] https://goframe.org/pages/viewpage.action?pageId=30740166
>
>> [2] https://go-zero.dev/docs/concepts/layout

## 准备工作

## 最佳实践

- go mod tidy
- 脚手架 自动生成 service
- 业务逻辑开始写起。
- 生成pb
  ```
  cd 项目根目录
  protoc --go_out=. --go-grpc_out=. manifest/protobuf/${文件名}.proto 
  ```
# 帮助
- go get 慢
  - Mac/linux: ```export GOPROXY=https://goproxy.cn```
  - Windows: ```SET GOPROXY="https://goproxy.cn"```

## TODO

1. ~~日志组件~~
2. 配置解析
3. 注入依赖
4. 命令创建微服务：创建文件夹、文件
5. redis
6. 本地缓存
7. mysql
8. http
9. grpc
10. 服务注册 服务发现
11. 压测
12. 定时任务

## 库

1. [gRPC-Go](https://github.com/grpc/grpc-go) golang RPC库
2. [go-kit](https://github.com/go-kit/kit) 构建微服务工具
3. [mux](https://github.com/gorilla/mux) URL 路由和调度器
4. [gin](https://github.com/gin-gonic/gin) web 框架
5. [wire](https://github.com/google/wire) 注入依赖
6. [viper](https://github.com/spf13/viper) 解析配置文件
7. [zap](https://github.com/uber-go/zap) 日志
8. [sentinel-golang](https://github.com/alibaba/sentinel-golang)
   流控 [中文文档](https://sentinelguard.io/zh-cn/docs/golang/basic-api-usage.html)
9. [fasthttp](https://github.com/valyala/fasthttp) http库
10. [chaosmonkey](https://github.com/Netflix/chaosmonkey) 环境故障模拟工具
11. [useragent](https://github.com/mssola/useragent) 解析 User Agent
12. [hey](https://github.com/rakyll/hey) 压测工具
13. [vegeta](https://github.com/tsenart/vegeta) 另一个压测工具
14. [go-redis](https://github.com/redis/go-redis) Redis 客户端
15. [clickhouse-go](https://github.com/ClickHouse/clickhouse-go) ClickHouse 客户端
16. [etcd](https://github.com/etcd-io/etcd) 分布式 K/V 存储
17. [bigcache](https://github.com/allegro/bigcache) 本地缓存 (可以再多看几个对比)
18. [uuid](https://github.com/google/uuid) uuid
19. [jaeger](https://github.com/jaegertracing/jaeger) 分布式追踪系统
20. [gocron](https://github.com/jasonlvhit/gocron) 定时任务
21. [sonic](https://github.com/bytedance/sonic) 极快的 JSON 序列化 / 反序列化
