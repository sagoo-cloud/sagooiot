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

English| [简体中文](README_ZH.md)

## Copyright Notice

Open source software is not the same as free. SagooIOT is released under the [GPL-3.0](LICENSE) open source license and provides technical exchange and learning. However, according to this license, modified or derived code from SagooIOT may not be released or sold as closed-source commercial software. If you need to use SagooIOT for any commercial purposes locally, please contact the project manager for commercial licensing to ensure your use complies with the GPL license.

## About SagooIOT

SagooIOT is a lightweight IoT platform developed in Go. It supports cross-platform IoT access and management solutions. The platform implements basic functions related to IoT development, based on which a complete set of IoT-related business systems can be quickly built.

Front-end project: https://github.com/sagoo-cloud/sagooiot-ui

Official documentation: http://iotdoc.sagoo.cn/

Official QQ group: 686637608

Welcome to support us by clicking the 💘Star💘 in the upper right corner.

**After the system is running, the default username and password of the system are:**

Username: admin

Password: admin123456

## Platform overview

* A management system with front-end and back-end separation developed based on the Go Frame 2.0+Vue3+Element Plus
* The front-end uses vue-next-admin, Vue, and Element UI.

## Features

* High productivity: a backend management system can be built in a few minutes
* Modularity: a single application with multiple systems, which divides a complete application into multiple systems for easier expansion and increased code reuse.
* Authentication mechanism: adopts gtoken user status authentication and casbin authorization
* Routing mode: thanks to goframe2.0, which provides a standardized routing registration method, automatic API documentation is generated without annotations
* Interface-oriented development
* Supports object models, multi-product, and multi-device access and management.
* Blurs the complexity of network protocols, adapts to multiple access protocols (TCP, MQTT, UDP, CoAP, HTTP, gRPC, RPC, etc.), and flexibly connects to devices from different manufacturers.
* Supports cross-platform operation, can quickly implement edge computing functions, realize offline automatic alerts, and automatic execution.
* Supports cross-terminal display, can monitor device status and display data through PC, mobile phone, and tablet.
* Unique plugin system, supports cross-language access, can be quickly accessed through plugins written in C/C++, Python.
* The plugin system supports hot-plugging, supports Modbus tcp, modbus rtu, modbus ascii, iec61850, opc, and other data acquisition protocols.

## Built-in functions

1. User management: Users are system operators. This function mainly completes system user configuration.
2. Department management:** Configure system organization (company, department, group), tree structure display supports data permissions.
3. Job management:** Configure the position of the system user.
4. Menu management:** Configure system menus, operation permissions, button permission labels, etc.
5. Role management:** Role menu permission allocation, set role data scope permission division according to the organization.
6. Dictionary management:** Maintain some relatively fixed data that are frequently used in the system.
7. Parameter management:** Dynamically configure common parameters for the system.
8. Operation log:** Record and query normal system operation logs; record and query system exception information logs.
9. Login log:** Query system login logs, including login exceptions.
10. Online users:** Monitor the status of active users in the current system.
11. Scheduled tasks:** Online (add, modify, delete) task scheduling, including execution result logs.
12. Code generation:** Generate front-end and back-end code.
13. Service monitoring:** Monitor CPU, memory, disk, stack, etc. of the current system.
14. File upload, cache tags, etc.
15. Product management:** Unified management of device-type products
16. Device management:** Manage device access and data configuration
17. Data center:** Manage the creation of new data models for third-party APIs, databases, and internal data, and support rule definition.


## demo

| ![login](https://iotdoc.sagoo.cn/imgs/demo/01.png)     | ![overview](https://iotdoc.sagoo.cn/imgs/demo/02.png)                       |
|--------------------------------------------------------|-----------------------------------------------------------------------------|
| ![thing](https://iotdoc.sagoo.cn/imgs/demo/03.png)     | ![monitoring](https://iotdoc.sagoo.cn/imgs/demo/04.png)                     |
| ![deviceLog](https://iotdoc.sagoo.cn/imgs/demo/05.png) | ![video](https://iotdoc.sagoo.cn/imgs/demo/08.png)                          |
| ![NotificationConfiguration](https://iotdoc.sagoo.cn/imgs/demo/09.png)   | ![Alarm Configuration Management](https://iotdoc.sagoo.cn/imgs/demo/10.png) |
| ![Alarm Rule Configuration](https://iotdoc.sagoo.cn/imgs/demo/11.png)    | ![user](https://iotdoc.sagoo.cn/imgs/demo/12.png)                           |
| ![system monitor](https://iotdoc.sagoo.cn/imgs/demo/13.png)      | ![data hub](https://iotdoc.sagoo.cn/imgs/demo/14.png)                       |
| ![Visualization Rule Engine](https://iotdoc.sagoo.cn/imgs/demo/07.png)   | ![screen](https://iotdoc.sagoo.cn/imgs/demo/06.png)                          |

![configuration](https://iotdoc.sagoo.cn/imgs/configure.jpg)

## Disclaimer

SagooIOT Community Edition is an open source learning project and is not related to commercial activities. Users should comply with laws and regulations when using this project and should not engage in illegal activities. If SagooIOT finds that a user is engaged in illegal activities, it will cooperate with relevant authorities to investigate and report to the government. Users shall bear any legal liability arising from illegal activities themselves. If a third party is damaged due to the user's use, the user shall compensate the third party in accordance with the law. Users bear all risks in using all relevant resources of SagooIOT.


## stars

[![Star History Chart](https://api.star-history.com/svg?repos=sagoo-cloud/sagooiot&type=Date)](https://star-history.com/#sagoo-cloud/sagooiot&Date)