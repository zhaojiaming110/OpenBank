## OpenBank微服务

此项目是由Go-micro加其它开源框架实现的一套微服务框架，其项目背景来自于我曾经做过的开放银行系统。

此项目纯属是为了提高golang编程技巧及熟悉Golang技术栈而开发的，目前已完成主体框架，还在进一步完善中。

### 项目整体架构图

![](https://tva1.sinaimg.cn/large/007S8ZIlly1gekjjz3ucxj30w00qr3zn.jpg)

### 组件

- Micro API网关：外界访问微服的唯一入口，对外提供了HTTP入口。
- Micro Web：内部访问微服的入口，对外提供HTTP入口
- 定时器：通过micro web调用内部的定时任务。
- Nsq消息中间件：主要用来实现交易的异步模式。
- Etcd集群：使用了ETCD作为微服的服务发现。
- 统一配置中心：使用gRpc实现了配置的统一管理。
- Mysql集群
- Redis集群
- 文件服务系统（规划中）
- Elastic日志系统（规划中）

### 技术栈

- golang
- go-micro
- etcd
- grpc
- protobuf
- mysql
- Redis
- Nsq

### 目录结构

- account-web  openbank微服务api入口
- account-srv  openbank微服务
- asynchronous-server 异步微服务 通过nsq消息队列实现
- timing-server 定时任务微服务， 通过外界调用http
- basic 基础配置
- config-srv  grpc注册中心
- micro micro api网关
- plugins 插件

### 已实现的功能

- 同异步模式
- 定时任务
- 注册中心
- 链路追踪
- 日志持久化
- 服务的熔断、降级