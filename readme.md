### 目录结构

``` shell
dousheng
├─ app # 业务开发代码目录
│	├─api # 控制层开发代码目录
│	├─golbal #返回信息存放以及配置、数据库、日志初始化目录
│	├─internal	#关键业务代码目录
│	│	├─middleware #中间件目录
│	│	├─service #业务层代码目录
│	│	└─model #持久层代码目录
│	└─ router #路由配置初始化
├─ manifest # 配置目录，不应该缓存到git
├─ boot # 数据库、OSS对象存储、日志服务初始化、定时任务等初始化全局业务相关代码目录
├─ imgs # 系统头像存放目录
├─ log # 系统日志存放目录
├─ test # 单元测试目录
└─ utils # 各种全局可用的包，依赖，工具
```

### 快速开始

```
git clone https://github.com/let-s-go-qxy/dousheng.git
```

#### 下载依赖

```
go mod tidy
```

##### 运行

```
go build tiktok

go run tiktok
```

### 技术选型

**整体的后端架构为hertz+gorm：**

gorm是go后端框架中使用最为广泛的全功能 ORM，开发者友好的特性，多功能的各类API，可靠的性能和报错处理，都是我们选择它的理由。

hertz作为字节开源的一款十分优秀的Golang 微服务 HTTP 框架，吸收了很多传统框架如gin, fasthttp, echo的优势，并结合字节跳动内部的需求，使其具有高易用性、高性能、高扩展性等特点开发而来，经受了字节内部的企业级项目考验，其可靠性无需多言。在性能方面，Hertz 默认使用自研的高性能网络库 Netpoll，在一些特殊场景相较于 go net，Hertz 在 QPS、时延上均具有一定优势。

**在数据库选型上，我们采取了最广泛使用的MySQL，处理数据缓存和备份上，我们使用了Redis：**

之所以采取mysql+redis，是由于我们根据实际业务需求，发现一些场景下，我们无需多次去查询数据库，而是可以直接通过将数据缓存在Redis中。例如用户的喜欢列表，由于一个视频的点赞操作可能会十分频繁，而对于一个喜欢列表，在一些场景下我们无需实时更新其数据。这种情况下，我们可以采用使用Redis存储喜欢列表查询的数据缓存，定时更新该缓存，而无需每次都去数据库中查询大量的点赞记录。

**日志记录和存储方面，我们使用了Zap日志库和后端框架自带的log打印：**

我们在开发项目过程中，有可能遇到一些报错，并没有即时处理和记录，或者并没有当场发现错误产生，导致我们在往往会错过一些很有价值的报错记录和性能异常等记录。作为一款对性能和内存分配做了极致的优化的开源日志库Zap，我们可以利用其强大的日志处理和记录功能来处理我们的运行记录。另一方面，使用hertz和gorm的日志显示功能，我们也可以即时在运行过程中查看到运行过程和请求，数据库处理结果。

**视频流播放处理模块，采用了阿里云OSS对象存储服务：**

阿里云对象存储OSS（Object Storage Service）是一款海量、安全、低成本、高可靠的云存储服务，可提供99.9999999999%（12个9）的数据持久性，99.995%的数据可用性。多种存储类型供选择，全面优化存储成本。并且其提供的api调用和相关JDK都十分方便，对开发者友好。我们打算先使用该服务将视频流处理分割出整体服务中去，在后期对项目进行升级优化时进一步处理视频流的上传，下载和播放等功能。

**用户聊天模块，采用RabbitMQ解决消息重复拉取的问题，同时利用MQ集群削峰填谷提高系统可用性：**

在聊天系统中我们使用RabbitMQ针对消息重复拉取的问题进行了优化和解决。配合HAProxy实现RabbitMQ集群搭建，起到削峰填谷、负载均衡的作用。
