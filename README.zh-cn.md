Sagoo IOT
========

<div align="center">
	<img width="140px" src="https://foruda.gitee.com/avatar/1692323731930718042/10619366_sagoo-cloud_1692323731.png!avatar100">
    <p>
        <h1>SagooIOT </h1>
    </p>
    <p align="center">
        <a href="https://goframe.org/pages/viewpage.action?pageId=1114119" target="_blank">
	        <img src="https://img.shields.io/badge/goframe-2.2-green" alt="goframe">
	    </a>
	    <a href="https://v3.vuejs.org/" target="_blank">
	        <img src="https://img.shields.io/badge/vue.js-vue3.x-green" alt="vue">
	    </a>
		<a href="https://www.tslang.cn/" target="_blank">
	        <img src="https://img.shields.io/badge/typescript-%3E4.0.0-blue" alt="typescript">
	    </a>
		<a href="https://vitejs.dev/" target="_blank">
		    <img src="https://img.shields.io/badge/vite-%3E2.0.0-yellow" alt="vite">
		</a>
		<a href="https://github.com/sagoo-cloud/sagooiot/blob/main/LICENSE" target="_blank">
		    <img src="https://img.shields.io/badge/license-GPL3.0-success" alt="license">
		</a>
	</p>
</div>

[English](README.MD)| 简体中文


## 版权声明

开源软件并不等同于免费。SagooIOT 遵循 [GPL-3.0](LICENSE) 开源协议，并提供技术交流学习。但根据该协议，修改或衍生自 SagooIOT 的代码，不得以闭源的商业软件形式发布或销售。如果您需要将 SagooIOT 在本地用于任何商业目的，请联系项目负责人进行商业授权，以确保您的使用符合 GPL 协议。

## 关于SagooIOT

SagooIOT是一个基于golang开发的轻量级的物联网平台。支持跨平台的物联网接入及管理方案，平台实现了物联网开发相关的基础功能，基于该功能可以快速的搭建起一整套的IOT相关的业务系统。

前端工程：https://github.com/sagoo-cloud/sagooiot-ui

官方文档：http://iotdoc.sagoo.cn/

官方QQ群：686637608

欢迎点右上角 💘Star💘支持我们。



**系统运行后，系统默认的用户名与密码为：**

用户：admin

密码：admin123456


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

| ![login](https://iotdoc.sagoo.cn/imgs/demo/01.png)                     | ![overview](https://iotdoc.sagoo.cn/imgs/demo/02.png)                       |
|------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| ![thing](https://iotdoc.sagoo.cn/imgs/demo/03.png)                     | ![monitoring](https://iotdoc.sagoo.cn/imgs/demo/04.png)                     |
| ![deviceLog](https://iotdoc.sagoo.cn/imgs/demo/05.png)                 | ![video](https://iotdoc.sagoo.cn/imgs/demo/08.png)                          |
| ![NotificationConfiguration](https://iotdoc.sagoo.cn/imgs/demo/09.png) | ![Alarm Configuration Management](https://iotdoc.sagoo.cn/imgs/demo/10.png) |
| ![Alarm Rule Configuration](https://iotdoc.sagoo.cn/imgs/demo/11.png)  | ![user](https://iotdoc.sagoo.cn/imgs/demo/12.png)                           |
| ![system monitor](https://iotdoc.sagoo.cn/imgs/demo/13.png)            | ![data hub](https://iotdoc.sagoo.cn/imgs/demo/14.png)                       |
| ![Visualization Rule Engine](https://iotdoc.sagoo.cn/imgs/demo/07.png) | ![screen](https://iotdoc.sagoo.cn/imgs/demo/06.png)                         |

![configuration](https://iotdoc.sagoo.cn/imgs/configure.jpg)

## 免责声明：

SagooIOT社区版是一个开源学习项目，与商业行为无关。用户在使用该项目时，应遵循法律法规，不得进行非法活动。如果SagooIOT发现用户有违法行为，将会配合相关机关进行调查并向政府部门举报。用户因非法行为造成的任何法律责任均由用户自行承担，如因用户使用造成第三方损害的，用户应当依法予以赔偿。使用SagooIOT所有相关资源均由用户自行承担风险.