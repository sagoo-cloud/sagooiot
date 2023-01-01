Sagoo IOT
========

SagooIOT是一个基于golang开发的轻量级的物联网平台。支持跨平台的物联网接入及管理方案，平台实现了物联网开发相关的基础功能，基于该功能可以快速的搭建起一整套的IOT相关的业务系统。




## 平台简介
* 基于全新Go Frame 2.0+Vue3+Element Plus开发的全栈前后端分离的管理系统
* 前端采用vue-next-admin 、Vue、Element UI。



## 特征
* 高生产率：几分钟即可搭建一个后台管理系统
* 模块化：单应用多系统的模式，将一个完整的应用拆分为多个系统，后续扩展更加便捷，增加代码复用性。
* 认证机制：采用gtoken的用户状态认证及casbin的权限认证
* 路由模式：得利于goframe2.0提供了规范化的路由注册方式,无需注解自动生成api文档
* 面向接口开发
* 支持物模型，多产品、多设备接入管理。
* 屏蔽网络协议的复杂性，适配多种接入协议(TCP,MQTT,UDP,CoAP,HTTP,GRPC,RPC等),灵活接入不同厂家的不同设备。
* 支持跨平台运行，可快速实现边缘计算功能，实现离线自动预警，自动执行等相关功能。
* 支持跨终端展示，可以通过PC,手机，平板等进行设备状态的监控和数据展示
* 独特的插件系统，支持跨语言接入，可以通过C/C++,Python编写的插件进行快速接入。
* 插件系统支持热插拔，支持Modbus tcp,modbus rtu,modbus ascii,iec61850,opc等数据采集协议


## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 在线用户：当前系统中活跃用户状态监控。
11. 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
12. 代码生成：前后端代码的生成。
13. 服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
14. 文件上传,缓存标签等。
15. 产品管理：对设备类产品进行统一管理
16. 设备管理：对设备进行接入与数据配置管理
17. 数据中心：对第三方api或是数据库及内部数据进行数据新建模管理，支持规则定义。


## 环境需求

**操作系统**

Linux / macOS /Win

**Go 编译环境**

依赖 `Go1.16+` 编译环境，可前往[官方网站](https://golang.org/dl/) 或 [国内镜像](https://golang.google.cn/dl/) 下载安装。

**MySQL**

MySQL 5.6+



## 技术栈

开发语言：golang

服务端基础架构：[GoFrame](https://github.com/gogf/gf) 【 [中文文档](https://goframe.org/index) 】

前端框架 [vue-next-admin](https://gitee.com/lyt-top/vue-next-admin)【[中文文档](https://lyt-top.gitee.io/vue-next-admin-doc-preview/)】

swaggo https://github.com/swaggo/swag

gtoken https://github.com/goflyfox/gtoken

casbin https://github.com/casbin/casbin

## 工程目录

```
/
├── api
├── hack
├── internal
│   ├── cmd
│   ├── consts
│   ├── controller
│   ├── dao
│   ├── logic
│   ├── model
│   |   ├── do
│   │   └── entity
│   └── service
├── manifest
├── resource
├── utility
├── go.mod
└── main.go

```

| 目录/文件名称   | 说明     | 描述                                                         |
| :-------------- | :------- | :----------------------------------------------------------- |
| `api`           | 对外接口 | 对外提供服务的输入/输出数据结构定义。考虑到版本管理需要，往往以`api/v1...`存在。 |
| `hack`          | 工具脚本 | 存放项目开发工具、脚本等内容。例如，`CLI`工具的配置，各种`shell/bat`脚本等文件。 |
| `internal`      | 内部逻辑 | 业务逻辑存放目录。通过`Golang internal`特性对外部隐藏可见性。 |
| ` - cmd`        | 入口指令 | 命令行管理目录。可以管理维护多个命令行。                     |
| ` - consts`     | 常量定义 | 项目所有常量定义。                                           |
| ` - controller` | 接口处理 | 接收/解析用户输入参数的入口/接口层。                         |
| ` - dao`        | 数据访问 | 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 `CURD` 方法 |
| ` - logic`      | 业务封装 | 业务逻辑封装管理，特定的业务逻辑实现和封装。往往是项目中最复杂的部分。 |
| ` - model`      | 结构模型 | 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义。 |
| `  - do`        | 领域对象 | 用于`dao`数据操作中业务模型与实例模型转换，由工具维护，用户不能修改。 |
| `  - entity`    | 数据模型 | 数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改。 |
| ` - service`    | 业务接口 | 用于业务模块解耦的接口定义层。具体的接口实现在`logic`中进行注入。 |
| `manifest`      | 交付清单 | 包含程序编译、部署、运行、配置的文件。常见内容如下：         |
| ` - config`     | 配置管理 | 配置文件存放目录。                                           |
| ` - docker`     | 镜像文件 | `Docker`镜像相关依赖文件，脚本文件等等。                     |
| ` - deploy`     | 部署文件 | 部署相关的文件。默认提供了`Kubernetes`集群化部署的`Yaml`模板，通过`kustomize`管理。 |
| `resource`      | 静态资源 | 静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中。 |
| `go.mod`        | 依赖管理 | 使用`Go Module`包管理的依赖描述文件。                        |
| `main.go`       | 入口文件 | 程序入口文件。                                               |

## 本地源码运行

修改config下的config.toml文件，并配置相关项

请跟据注释进行配置修改，包括服务相关配置，日志相关

`go run main.go`

## 关于build.sh编译脚本

可以使用build.sh进行程序编译，如果在使用build.sh脚本进行程序编译时，提示

```
fatal: No names found, cannot describe anything.
./build.sh linux|windows|mac

```
是因为源码没有进行git版本的标签设置。

支持将git的tag编译到程序中。需要创建git的tag。只有创建了tag，编译的程序才会显示版本号。

```
git tag v0.0.1

git push origin v0.0.1
```

## 编译后执行脚本

编译后生成的可执行文件放在bin目录下，将bin目录下的文件放到目标服务器，执行`./curl.sh start` 运行即可。

```
curl.sh脚本参数：

start|stop|restart|status|tail

```

分别对应 启动、停止、重启、状态、显示动态日志运行信息

## 插件编译

如果要使用插件，需要提前将插件进行编译。直接使用plugins下面的编译脚本直接执行就可以。


## 其它问题

### 如果在macOS下遇到 Warning :`IOMasterPort`：
```
warning: ‘IOMasterPort‘ is deprecated: first deprecated in macOS 12.0 [-Wdeprecated-declarations]
```
**原因**

依赖包跟MacOS的版本有兼容问题。

解决方案
切换CGO编译方式
```
go env -w CGO_ENABLED="0"
```

### 如果采用HTTPS方式时SSE不工作，需要如下配置：

```Nginx
    proxy_set_header Connection '';
    proxy_http_version 1.1;
    chunked_transfer_encoding off; 
```
