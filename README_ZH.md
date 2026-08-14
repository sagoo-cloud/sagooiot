Sagoo IOT Community V2
========

<div align="center">
<img width="120px" src="https://raw.githubusercontent.com/sagoo-cloud/.github/main/profile/logo.svg">

[![GoFrame](https://img.shields.io/badge/goframe-2.9-green)](https://goframe.org/pages/viewpage.action?pageId=1114119)
[![vue](https://img.shields.io/badge/vue.js-vue3.x-green)](https://v3.vuejs.org/)
[![typescript](https://img.shields.io/badge/typescript-%3E4.0.0-blue)](https://www.tslang.cn/)
[![vite](https://img.shields.io/badge/vite-%3E2.0.0-yellow)](https://vitejs.dev/)
[![LICENSE](https://img.shields.io/badge/license-GPL3.0-success)](https://github.com/sagoo-cloud/sagooiot/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.23.0+-blue)](https://golang.org)
[![GitHub stars](https://img.shields.io/github/stars/sagoo-cloud/sagooiot?style=social)](https://github.com/sagoo-cloud/sagooiot)

</div>

[English](./README.md) | 简体中文

---

## 🚀 快速导航

| 📖 文档 | 💻 前端  | 💬 社区 | 🐛 反馈 |
|--------|--------|--------|--------|
| [官方文档](http://iotdoc.sagoo.cn/) | [UI源码](https://github.com/sagoo-cloud/sagooiot-ui) | [QQ群:686637608](http://qq.sagoo.cn) | [Issues](https://github.com/sagoo-cloud/sagooiot/issues) |

> 💡 **欢迎点右上角 ⭐Star⭐ 支持我们！**

---

## 📝 版权声明

开源软件并不等同于免费。SagooIOT 遵循 [GPL-3.0](LICENSE) 开源协议，并提供技术交流学习。但根据该协议，修改或衍生自 SagooIOT 的代码，**不得以闭源的商业软件形式发布或销售**。如果您需要将 SagooIOT 在本地用于任何商业目的，请联系项目负责人进行商业授权，以确保您的使用符合 GPL 协议。

---

## 📋 关于 SagooIOT

**SagooIOT** 是一个基于 Go 语言开发的**轻量级企业级物联网平台**。它提供完整的物联网接入、管理、分析和应用解决方案，支持跨平台独立部署或是分布式部署，可快速搭建起包含设备管理、数据处理、告警通知、规则引擎、视频监控、边缘计算等功能的完整 IoT 业务系统。

### 🎯 核心价值

- **🚀 快速部署** - 开箱即用，几分钟即可启动完整的 IoT 平台
- **📱 全栈分离** - GoFrame 2.9 后端 + Vue 3 前端，架构清晰易维护
- **🔌 灵活接入** - 支持 TCP、UDP、HTTP、Websocket、MQTT、CoAP、OPC UA、Modbus、SNMP、ICE104、JT808、GB212等多种协议
- **🧩 插件驱动** - 独特的热插拔插件系统，支持 C/C++、Python、Go 跨语言开发
- **⚡ 高效数据处理** - 集成 TDengine 时序数据库，支持百万级数据点秒级处理
- **🏭 边缘友好** - 支持离线部署、本地规则执行、自动预警，适合边缘计算场景

### 📌 快速信息

| 项目 | 信息 |
|-----|------|
| **前端工程** | [sagooiot-ui](https://github.com/sagoo-cloud/sagooiot-ui) |
| **官方文档** | http://iotdoc.sagoo.cn/ |
| **技术社区** | QQ 群: 686637608 |
| **默认账户** | admin / admin123456 |
| **开源协议** | [GPL-3.0](LICENSE) |

---

## ⚙️ 系统要求

### 最低配置
- **操作系统** - Linux、macOS、Windows
- **Go 版本** - 1.23.0 或更高版本
- **内存** - 4GB（推荐 8GB+）
- **磁盘** - 20GB（推荐 50GB+）
- **网络** - 稳定的网络连接

### 核心依赖
| 组件 | 版本          | 用途                          |
|-----|-------------|-----------------------------|
| **MySQL** | 5.7+ / 8.0+ | 关系数据库，存储业务数据                |
| **PostgreSQL** | 16.x        | 关系数据库，存储业务数据（可选）            |
| **Redis** | 6.0+        | 缓存和消息队列                     |
| **TDengine** | 3.0+        | 时序数据库（可选，用于高效存储设备时序数据）      |
| **Influxdb** | 2.x         | 时序数据库（可选，用于高效存储设备时序数据）      |
| **MQTT Broker** | -           | 消息中间件（可选，推荐 Mosquitto 或云服务） |
| **MinIO** | -           | 对象存储（可选，用于文件管理）             |

---

## 🚀 快速开始

### 方式一：直接运行

#### 1. 前置准备
```bash
# 克隆项目
git clone https://github.com/sagoo-cloud/sagooiot.git
cd sagooiot

# 创建必要的目录
mkdir -p resource/log

# 初始化数据库（参考 manifest/sql/）
# 需要在 MySQL 中导入数据库初始化脚本
```

#### 2. 配置环境变量
在项目根目录创建 `.env` 文件或修改 `manifest/config/config.dev.yaml`：

```yaml
# 数据库配置
database:
  mysql:
    host: localhost
    port: 3306
    name: sagooiot
    user: root
    password: your_password

# Redis 配置
redis:
  host: localhost
  port: 6379
  password: ''

# MQTT 配置（如需要）
mqtt:
  broker: mqtt://localhost:1883
```

#### 3. 启动应用
```bash
# 安装依赖
go mod download

# 运行应用
go run main.go

# 应用将在 http://localhost:8000 启动
```

### 方式二：Docker Compose 一键启动

#### 1. 准备 Docker 环境
```bash
# 确保已安装 Docker 和 Docker Compose
docker --version
docker-compose --version
```

#### 2. 启动完整栈
```bash
cd manifest/docker-compose

# 使用 docker-compose 启动所有服务（MySQL、Redis、SagooIOT 等）
docker-compose -f docker-compose.yml up -d

# 查看日志
docker-compose logs -f sagooiot

# 停止服务
docker-compose down
```

#### 3. 首次访问
- **Web UI** - http://localhost:8000
- **API** - http://localhost:8000/api
- **默认账户** - admin / admin123456

> ⚠️ **首次登录后强烈建议修改默认密码！**

---

## 🏗️ 平台架构

### 整体架构设计

```
┌────────────────────────────────────────────────────────────┐
│                    前端层 (Web UI)                         │
│            Vue 3 + Element Plus + TypeScript               │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│                 GoFrame 2.9 应用服务层                      │
│  ┌──────────────┬──────────────┬──────────────────────────┐ │
│  │ 设备管理模块  │  数据处理模块  │    告警通知模块          │ │
│  │ 物模型管理    │  规则引擎      │    数据分析              │ │
│  │ 权限管理      │  数据中心      │    实时推送(WebSocket)  │ │
│  │ 系统管理      │  任务调度      │    其他业务模块          │ │
│  └──────────────┴──────────────┴──────────────────────────┘ │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│              协议接入层 (Multi-Protocol)                    │
│  TCP │ MQTT │ UDP │ CoAP │ HTTP  │ RPC              │
│              ↓                                              │
│        插件系统 (C/C++/Python/Go)                           │
│  Modbus TCP/RTU/ASCII │ IEC61850 │ OPC │ Canopen         │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│               存储和消息中间件层                             │
│  ┌──────────┬──────────┬──────────┬──────────┐             │
│  │  MySQL   │  Redis   │ TDengine │  MinIO   │             │
│  │  关系数据 │  缓存/队列 │ 时序数据 │  文件存储  │             │
│  └──────────┴──────────┴──────────┴──────────┘             │
│  MQTT Broker │ Elasticsearch │ 其他存储组件                 │
└────────────────────────────────────────────────────────────┘
```

### 核心模块说明

| 模块 | 位置 | 功能 | 说明 |
|-----|------|------|------|
| **Controller** | `internal/controller/` | API 路由层 | 接收和响应 HTTP 请求 |
| **Service** | `internal/service/` | 业务逻辑层 | 实现业务规则和流程 |
| **DAO** | `internal/dao/` | 数据访问层 | 数据库操作和查询 |
| **Model** | `internal/model/` | 数据模型 | 定义数据结构 |
| **Network** | `network/core/` | 协议接入核心 | 多协议适配和处理 |
| **MQTT** | `internal/mqtt/` | MQTT 消息处理 | MQTT 发布/订阅管理 |
| **Tasks** | `internal/tasks/` | 后台任务 | 定时任务、异步任务 |
| **Workers** | `internal/workers/` | 工作线程 | 高并发处理和队列 |
| **Pkg** | `pkg/` | 工具库 | 缓存、MQTT、插件、OAuth 等 |

### 数据流向

```
设备端 → 协议接入 → 协议解析 → 数据验证 → 业务处理 → 数据存储
                    ↓
                实时推送 → WebSocket → 前端展示
                
                ↓
              规则引擎 → 告警判断 → 通知发送
              
                ↓
              数据分析 → 可视化报表
```

---

## ✨ 功能模块

### A. 设备与物联网核心能力

1. **物模型管理** - 定义设备的属性、事件、服务，支持 JSON 格式导入，灵活适配各类设备
2. **产品管理** - 统一管理设备类型产品，支持产品版本、批量操作、产品分类
3. **设备管理** - 完整的设备生命周期管理：注册、配置、上线/下线、激活、禁用、删除
4. **设备树** - 树形展示设备关系与分组，支持多维度分类和权限控制
5. **设备标签** - 灵活的标签系统，便于设备分类、查询、权限隔离
6. **实时数据** - 设备实时状态展示、历史数据查询、数据导出（Excel/CSV）

### B. 协议与接入管理

1. **多协议适配** - 原生支持 TCP、UDP、HTTP、Websocket、MQTT、CoAP 等通信协议
2. **协议网关** - 协议转换和网关管理，支持异构设备接入（提供二次开发SDK，快速开发新协议接入）
3. **网络隧道** - 内网穿透和安全加密，支持 P2P 连接和代理转发
4. **插件系统** - 独特的热插拔插件架构：
   - 支持 C/C++、Python、Go 等多语言编写
   - 跨进程通信（gRPC）确保隔离性
   - 插件生命周期管理和热更新
   - 丰富的插件示例和文档
5. **数据采集协议** - 集成 Modbus TCP/RTU/ASCII、IEC61850、OPC UA、ICE104 等工业协议

### C. 数据处理与分析

1. **数据中心** - 构建自定义数据模型、第三方 API 集成、数据库直连、数据管理
2. **规则引擎** - 可视化规则编辑器、支持 JavaScript 表达式、灵活的条件组合
3. **数据转换** - ETL 管道、数据格式转换、字段映射
4. **实时分析** - 流数据处理、聚合计算、窗口函数
5. **时序数据库** - 深度集成 TDengine、InfluxDB 等
6. **数据导出** - 重要数据支持导出（Excel、CSV、JSON）

### D. 告警与通知

1. **告警规则** - 创建、编辑、启用/禁用规则，支持多条件触发（值范围、变化率、频率等）
2. **告警级别** - 自定义告警等级（严重、重要、普通、提示等）和颜色标记
3. **告警日志** - 查询告警历史、统计和分析、支持时间范围和条件筛选
4. **告警处理** - 告警确认、升级、处理记录、处理意见备注
5. **通知模板** - 邮件、短信、webhook、钉钉、企业微信等多种通知方式
6. **通知配置** - 通知渠道管理、接收人配置、通知时间限制
7. **消息推送** - WebSocket 实时推送告警、支持多种推送策略
8. **告警统计** - 告警数据可视化报表，支持按时间、设备、级别统计分析
9. **告警治理** - 支持告警抑制、重复告警合并、告警降噪等功能

### E. 边缘计算与离线处理

1. **边缘计算** - 支持在边缘侧本地部署和运行，实现离线自治
2. **离线预警** - 设备离线自动预警和通知，支持离线时长配置
3. **自动执行** - 设备命令自动执行，支持延迟执行和重试
4. **场景管理** - 创建和管理应用场景（多设备联动规则），支持场景触发和编排

### F. 系统管理与权限

1. **用户管理** - 用户创建、编辑、删除、禁用/启用、部门分配、角色分配、密码重置
2. **部门管理** - 树形组织结构、部门权限隔离、数据权限控制、部门级联
3. **角色管理** - 角色创建、菜单权限分配、按钮权限分配、数据权限分级
4. **岗位管理** - 岗位配置、与用户关联、岗位权限
5. **菜单管理** - 动态菜单配置、按钮权限标识、API 权限映射、菜单排序
6. **权限管理** - 基于 RBAC 的权限控制、支持 Casbin 高级权限策略、权限缓存
7. **字典管理** - 系统字典维护、枚举值定义、字典分类、数据字典导出
8. **参数管理** - 系统参数动态配置、参数分类、参数编辑和版本管理
9. **组织管理** - 多级组织架构支持、组织权限隔离、组织级数据汇总

### G. 日志与监控

1. **操作日志** - 记录用户操作（增删改查）、变更历史、操作详情、IP 地址跟踪
2. **登录日志** - 登录成功/失败记录、异常登录告警、登录 IP 限制、登录地区显示
3. **设备日志** - 设备上报日志、通信日志、错误日志、设备命令执行日志
4. **系统监控** - 实时监控 CPU、内存、磁盘、网络、系统进程、数据库连接池
5. **在线用户** - 实时用户状态展示、会话管理、强制下线、在线人数统计
6. **定时任务** - 任务调度和管理、执行记录、失败重试、任务日志查询

### H. 开发与扩展工具

1. **代码生成** - 前后端代码自动生成、CRUD 模板、包含完整 API 和 UI
2. **API 文档** - Swagger 自动生成、在线测试工具、API 版本管理
3. **文件管理** - 文件上传下载、文件分享、版本控制、MinIO 集成
4. **模块扩展** - 模块化架构、快速插件开发、SDK 提供
5. **开发示例** - 提供丰富的开发示例和文档，帮助快速上手二次开发
6. **OpenAPI** - 完整的 OpenAPI 规范支持，方便第三方集成
7. **北向接口** - 提供标准化的北向接口，支持与其他系统对接

### I. 身份与授权

1. **OAuth 2.0** - 第三方 OAuth 集成（GitHub、微信、钉钉、企业微信等）
2. **OAuth 提供者** - 作为 OAuth 提供方支持第三方应用接入
3. **登录认证** - 账号密码、手机号、邮箱、二维码扫码等多种认证方式
4. **会话管理** - gToken JWT 令牌管理、会话超时配置、强制退出、并发登录控制
5. **审计追踪** - 用户行为审计、敏感操作记录、操作溯源
6. **SK/AK** - 支持基于访问密钥的身份认证，适用于 API 调用

---

## 🛠️ 技术栈

### 后端技术栈

| 技术                    | 版本        | 用途 | 优势 |
|-----------------------|-----------|------|------|
| **Go**                | 1.23.0+   | 核心语言 | ⚡ 高性能、易并发、编译快速 |
| **GoFrame**           | 2.2+      | Web 框架 | 🎯 规范化、功能完整、自动化 API 文档 |
| **MySQL**             | 5.7+/8.0+ | 关系数据库 | 📊 可靠性高、应用广泛、数据安全 |
| **PostgreSQL（可选）**    | 16.x      | 关系数据库 | 📊 可扩展、应用广泛、数据安全 |
| **Redis**             | 6.0+      | 缓存/队列 | ⚡ 高速读写、多数据结构、持久化 |
| **TDengine**          | 3.0+      | 时序数据库 | 📈 高性能时间序列存储、自动聚合 |
| **InfluxDB（可选）**      | 2.x       | 时序数据库 | 📈 高性能时间序列存储、自动聚合 |
| **MQTT**              | 原生支持      | 物联网协议 | 📡 轻量级、低带宽、广泛应用 |
| **Casbin**            | -         | 权限管理 | 🔐 灵活 RBAC/ABAC、高效缓存 |
| **gToken**            | -         | 会话管理 | 🎫 JWT 令牌、会话超时、多终端 |
| **Asynq**             | 0.24+     | 任务队列 | ⏳ Redis 原生、可靠投递、重试机制 |
| **Gorilla WebSocket** | -         | 实时推送 | 📲 标准 WebSocket、高性能 |
| **MinIO**             | -         | 对象存储 | 💾 S3 兼容、自托管、高可用 |

### 前端技术栈

| 技术 | 版本 | 用途 | 说明 |
|------|------|------|------|
| **Vue** | 3.x | 前端框架 | 🎨 现代化、声明式、易学易用 |
| **Element Plus** | - | UI 组件库 | 🧩 丰富组件、风格统一、文档完善 |
| **TypeScript** | 4.0+ | 类型语言 | ✅ 类型安全、开发体验优、易维护 |
| **Vite** | 2.0+ | 构建工具 | ⚡ 极速冷启动、热更新快、最小化产出 |
| **Pinia** | - | 状态管理 | 📦 轻量级、TypeScript 友好 |
| **Vue Router** | - | 路由管理 | 🗂️ 单页应用、懒加载、权限控制 |
| **Axios** | - | HTTP 客户端 | 🌐 简洁 API、请求拦截、超时控制 |

### 关键依赖详解

| 依赖包 | 作用 | 应用场景 |
|--------|------|----------|
| **paho.mqtt.golang** | MQTT 客户端 | 设备接入、消息处理、发布/订阅 |
| **gogf/gf** | Go Web 框架 | API 路由、自动文档、规范开发 |
| **hashicorp/go-plugin** | 跨进程插件系统 | 多语言插件、热插拔、隔离运行 |
| **taosdata/driver-go** | TDengine 驱动 | 高效时序数据存储、自动聚合 |
| **redis/go-redis** | Redis 客户端 | 缓存、会话存储、消息队列 |
| **hibiken/asynq** | 任务队列框架 | 定时任务、异步处理、失败重试 |
| **gorilla/websocket** | WebSocket 服务端 | 实时推送、双向通信、告警推送 |
| **minio/minio-go** | 对象存储客户端 | 文件上传下载、分布式存储 |
| **golang-jwt/jwt** | JWT 令牌库 | 身份认证、会话管理、令牌签名 |

---

## 📊 与主流开源平台对比

| 特性        | SagooIOT                     | EdgeX Foundry | Mainflux | ThingsBoard |
|-----------|------------------------------|---------------|----------|------------|
| **开发语言**  | Go                           | Go | Go | Java |
| **部署复杂度** | ⭐⭐                           | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **学习成本**  | ⭐⭐                           | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| **物模型**   | ✅ 完整                         | ✅ 完整 | ❌ | ✅ 完整 |
| **插件系统**  | ✅ 热插拔(多语言)                   | ✅ 模块化 | ✅ NATS | ❌ |
| **边缘计算**  | ✅ 原生支持                       | ✅ 原生支持 | ❌ | ✅ 有限 |
| **协议适配**  | TCP/MQTT/CoAP/HTTP/Websocket | 多种 | MQTT主要 | MQTT/CoAP |
| **权限管理**  | ✅ Casbin + RBAC              | ✅ 完整 | ✅ 完整 | ✅ 完整 |
| **规则引擎**  | ✅ 可视化 + JS                   | ✅ 有 | ❌ | ✅ 有 |
| **组态工具**  | ✅ 可视化组态工具                    | ❌ | ❌ | ❌  |
| **可视化大屏** | ✅ 数据可视化工具                    | ❌ | ❌ | ❌  |
| **视频监控**  | ✅ 流媒体服务                      | ❌ | ❌ | ❌  |
| **数据中心**  | ✅ 灵活自定义                      | ❌ | ❌ | ❌ |
| **轻量级**   | ✅ 高                          | ❌ 较重 | ✅ 高 | ❌ 较重 |
| **二次开发**  | ⭐⭐⭐⭐⭐                        | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

### 🎯 SagooIOT 核心竞争力

| 优势          | 说明                                    |
|-------------|---------------------------------------|
| **极简轻量**    | 核心功能齐全，部署占用资源少，适合中小型项目快速落地            |
| **插件驱动**    | 独特的热插拔插件系统，支持 C/C++/Python/Go，灵活扩展无上限 |
| **物模型优先**   | 完整的物模型支持，灵活的属性/事件/服务定义                |
| **全栈现代化**   | 规范的前后端分离架构，基于 GoFrame 2.9 和 Vue 3     |
| **快速定制**    | 模块化设计，快速进行二次开发和功能扩展，降低开发成本            |
| **开发效率**    | 自动化 API 文档生成、代码生成工具、快速迭代              |
| **灵活的数据模型** | 数据中心提供灵活的数据模型构建，支持自定义业务数据             |
| **边缘友好**    | 支持离线部署、本地规则执行、自动预警，适合 IoT 边缘场景        |
| **数据可视化**   | 集成数据可视化工具，提供丰富的图表和数据可视化功能             |
| **分布式架构**     | 支持单体单机部署、分布式部署、混合部署，满足不同场景需求            |                                  |

---

## 📸 平台演示

| ![login](https://iotdoc.sagoo.cn/imgs/demo/01.png)                     | ![overview](https://iotdoc.sagoo.cn/imgs/demo/02.png)                       |
|------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| ![thing](https://iotdoc.sagoo.cn/imgs/demo/03.png)                     | ![monitoring](https://iotdoc.sagoo.cn/imgs/demo/04.png)                     |
| ![deviceLog](https://iotdoc.sagoo.cn/imgs/demo/05.png)                 | ![video](https://iotdoc.sagoo.cn/imgs/demo/08.png)                          |
| ![NotificationConfiguration](https://iotdoc.sagoo.cn/imgs/demo/09.png) | ![Alarm Configuration Management](https://iotdoc.sagoo.cn/imgs/demo/10.png) |
| ![Alarm Rule Configuration](https://iotdoc.sagoo.cn/imgs/demo/11.png)  | ![user](https://iotdoc.sagoo.cn/imgs/demo/12.png)                           |
| ![system monitor](https://iotdoc.sagoo.cn/imgs/demo/13.png)            | ![data hub](https://iotdoc.sagoo.cn/imgs/demo/14.png)                       |
| ![Visualization Rule Engine](https://iotdoc.sagoo.cn/imgs/demo/07.png) | ![screen](https://iotdoc.sagoo.cn/imgs/demo/06.png)                         |

![configuration](https://iotdoc.sagoo.cn/imgs/configure.jpg)

---

## 📚 常见应用场景

### 1. 工业生产监控
- 工厂设备实时监控、生产数据采集
- 设备运维告警、故障预警和处理
- 生产效率分析和优化

### 2. 智慧楼宇
- 建筑设备（HVAC、照明、安防）联动控制
- 能耗管理和优化
- 环境监测（温湿度、空气质量等）

### 3. 环境监测
- 实时环境数据采集和监测
- 污染物预警和告警
- 环保数据上报和合规性验证

### 4. 智慧农业
- 田间传感器部署和数据采集
- 环境监测（温湿度、光照、土壤含水量）
- 自动化灌溉和施肥控制

### 5. 新能源管理
- 光伏、风电、储能设备监控
- 发电量实时统计和分析
- 电池管理系统（BMS）集成

### 6. 车联网平台
- 车辆设备接入和管理
- 车辆位置追踪和路线规划
- 故障预警和维护提醒


## 👥 社区与贡献

### 🤝 贡献指南

我们欢迎各种形式的贡献！无论是代码、文档、Bug 报告还是功能建议，都能帮助我们改进项目。

#### 提交 Bug 或功能建议

1. 访问 [Issues](https://github.com/sagoo-cloud/sagooiot/issues) 页面
2. 点击 "New Issue" 按钮
3. 选择相应的模板（Bug 报告或功能建议）
4. 详细描述问题或建议
5. 提交 Issue

#### 提交代码贡献

1. **Fork 项目**
   ```bash
   点击右上角的 Fork 按钮
   ```

2. **克隆到本地**
   ```bash
   git clone https://github.com/your-username/sagooiot.git
   cd sagooiot
   ```

3. **创建特性分支**
   ```bash
   git checkout -b feature/AmazingFeature
   ```

4. **进行修改和测试**
   ```bash
   # 修改代码
   # 运行测试
   go test ./...
   ```

5. **提交变更**
   ```bash
   git add .
   git commit -m 'Add some AmazingFeature'
   ```

6. **推送到分支**
   ```bash
   git push origin feature/AmazingFeature
   ```

7. **开启 Pull Request**
   - 在 GitHub 上打开 PR
   - 描述修改内容和目的
   - 等待维护者审核和合并

### 📋 贡献规范

- **代码风格** - 遵守 Go 官方编码规范，使用 `gofmt` 格式化代码
- **提交信息** - 使用清晰、简洁的提交信息，格式：`[类型] 简短描述`
  - `[feat]` - 新功能
  - `[fix]` - Bug 修复
  - `[docs]` - 文档更新
  - `[refactor]` - 代码重构
  - `[test]` - 测试相关
  - `[ci]` - CI/CD 相关
- **测试** - 提交 PR 前请确保所有测试通过
- **文档** - 新功能请附带相应的文档更新
- **分支命名** - 使用有意义的分支名称，如 `feature/xxx` 或 `fix/xxx`

### 🎯 我们需要的帮助

- 📝 **文档翻译** - 帮助我们将文档翻译为其他语言
- 🐛 **Bug 修复** - 报告和修复发现的问题
- ✨ **新功能** - 实现新的功能或改进现有功能
- 📚 **教程和示例** - 编写教程和集成示例
- 🌍 **本地化** - 支持更多语言和地区
- 🎨 **UI/UX 改进** - 改进用户界面和体验

### 📞 联系与交流

| 方式 | 地址 | 说明 |
|-----|------|------|
| **QQ 群** | 686637608 | 实时技术交流、问题讨论 |
| **GitHub Issues** | [Issues](https://github.com/sagoo-cloud/sagooiot/issues) | 报告 Bug、功能建议 |
| **GitHub Discussions** | [Discussions](https://github.com/sagoo-cloud/sagooiot/discussions) | 问题讨论、经验分享 |
| **官方文档** | http://iotdoc.sagoo.cn/ | 详细的功能文档和教程 |

### 🏆 致谢

感谢所有为 SagooIOT 做出贡献的开发者和用户！您的支持和反馈推动了项目的发展。

---

## 📖 文档与资源

| 资源          | 描述                 | 链接 |
|-------------|--------------------|------|
| **官方文档**    | 完整的功能文档和 API 参考    | http://iotdoc.sagoo.cn/ |
| **前端工程**    | Web UI 完整源码        | https://github.com/sagoo-cloud/sagooiot-ui |
| **云端网关SDK** | 解析网关SDK，用于快速开发私有协议 | https://github.com/sagoo-cloud/iotgateway |
| **云端网关示例**  | 基于网关SDK开发云网关示例     | https://github.com/sagoo-cloud/iotgateway-example |
| **插件示例**    | SagooIOT平台插件示例   | https://github.com/sagoo-cloud/sagooiot-plugins |

---

## 📋 常见问题 (FAQ)

### Q: SagooIOT 免费吗？
**A:** SagooIOT 遵循 GPL-3.0 开源协议，个人学习和使用完全免费。如需商业授权，请联系项目团队。

### Q: 支持哪些设备接入？
**A:** SagooIOT 支持任何支持 TCP、MQTT、UDP、CoAP、HTTP、Websocket 等协议的设备接入，也可通过插件扩展支持更多协议。

### Q: 数据量很大，性能能跟上吗？
**A:** SagooIOT 集成 TDengine 时序数据库，可支持百万级数据点秒级处理。同时支持数据分片、缓存优化等性能优化方案。

### Q: 能否部署在云平台上？
**A:** 完全支持。SagooIOT 可部署在 AWS、阿里云、腾讯云、华为云等任何云平台上，支持 Docker 和 Kubernetes。

### Q: 如何实现设备远程控制？
**A:** 通过设备管理模块中的设备命令功能，可实现远程控制。支持实时控制和定时控制。

### Q: 是否支持私有部署？
**A:** 完全支持私有部署。您可以在自己的服务器或内网上部署 SagooIOT，数据完全属于您。

### Q: 如何实现边缘计算？
**A:** SagooIOT 支持离线部署和本地规则执行，可在边缘设备上运行，实现边缘计算和离线自治。

### Q: 能否集成第三方系统？
**A:** 支持。SagooIOT 提供完整的 OpenAPI 和 北向接口，可轻松集成第三方系统。

---

## 📄 免责声明

SagooIOT 社区版是一个开源学习项目，与商业行为无关。用户在使用该项目时，应遵循法律法规，不得进行非法活动。

- 如果 SagooIOT 发现用户有违法行为，将会配合相关机关进行调查并向政府部门举报
- 用户因非法行为造成的任何法律责任均由用户自行承担
- 如因用户使用造成第三方损害的，用户应当依法予以赔偿
- 使用 SagooIOT 所有相关资源均由用户自行承担风险
- 本项目不提供任何担保或保证

---

## 📈 项目成长

感谢所有 Star 和贡献者的支持！

[![Star History Chart](https://star-history.dera.page/svg?repos=sagoo-cloud/sagooiot&type=Date)](https://star-history.dera.page/#sagoo-cloud/sagooiot&Date)

---

## 📜 开源协议

本项目采用 [GPL-3.0](LICENSE) 开源协议。详见 [LICENSE](LICENSE) 文件。

---

## 💬 联系我们

- **QQ 官方群** - 686637608
- **官方网站** - http://www.sagoo.cn
- **GitHub** - https://github.com/sagoo-cloud/sagooiot
- **Gitee** - https://gitee.com/sagoo-cloud/sagooiot
- **文档中心** - http://iotdoc.sagoo.cn/

---

<div align="center">

**Made with ❤️ by the SagooIOT Team**

如果这个项目对您有帮助，请考虑点击右上角的 ⭐ **Star** 来支持我们！

</div>

