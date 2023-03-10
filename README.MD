Sagoo IOT
========

SagooIOT是一个基于golang开发的轻量级的物联网平台。支持跨平台的物联网接入及管理方案，平台实现了物联网开发相关的基础功能，基于该功能可以快速的搭建起一整套的IOT相关的业务系统。

前端工程：https://github.com/sagoo-cloud/sagooiot-ui

官方文档：http://iotdoc.sagoo.cn/

官方QQ群：686637608

欢迎点右上角 💘Star💘支持我们。


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

## 演示图

| ![登录页](http://iotdoc.sagoo.cn/assets/img/01.1c18fa83.png) | ![物联概览](http://iotdoc.sagoo.cn/assets/img/02.cf806fcd.png) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ![物模型管理](http://iotdoc.sagoo.cn/assets/img/03.c53a7ad3.png) | ![设备数据监测](http://iotdoc.sagoo.cn/assets/img/04.56ad4c63.png) |
| ![设备数据日志](http://iotdoc.sagoo.cn/assets/img/05.047f7b6c.png) | ![视频监控](http://iotdoc.sagoo.cn/assets/img/08.9df1e582.png) |
| ![通知配置](http://iotdoc.sagoo.cn/assets/img/09.7cd7f1e4.png) | ![告警配置管理](http://iotdoc.sagoo.cn/assets/img/10.643301e3.png) |
| ![告警规则配置](http://iotdoc.sagoo.cn/assets/img/11.27ccb9d9.png) | ![用户管理](http://iotdoc.sagoo.cn/assets/img/12.e5f73f32.png) |
| ![系统监控](http://iotdoc.sagoo.cn/assets/img/13.9f1f168e.png) | ![数据汇聚中心](http://iotdoc.sagoo.cn/assets/img/14.098c87ed.png) |
| ![可视化规则引擎](http://iotdoc.sagoo.cn/assets/img/07.b48538a8.png) | ![可视化大屏](http://iotdoc.sagoo.cn/assets/img/06.8a674439.png) |

![组态界面](http://iotdoc.sagoo.cn/assets/img/configure.06cd9b6e.jpg)


## 免责声明：

SagooIOT社区版是一个开源学习项目，与商业行为无关。用户在使用该项目时，应遵循法律法规，不得进行非法活动。如果SagooIOT发现用户有违法行为，将会配合相关机关进行调查并向政府部门举报。用户因非法行为造成的任何法律责任均由用户自行承担，如因用户使用造成第三方损害的，用户应当依法予以赔偿。使用SagooIOT所有相关资源均由用户自行承担风险.