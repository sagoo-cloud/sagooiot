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

English | [简体中文](./README_ZH.md)

---

## 🚀 Quick Navigation

| 📖 Documentation | 💻 Frontend | 💬 Community | 🐛 Feedback |
|-----------------|-----------|------------|-----------|
| [Official Docs](http://iotdoc.sagoo.cn/) | [UI Source](https://github.com/sagoo-cloud/sagooiot-ui) | [QQ Group: 686637608](http://qq.sagoo.cn) | [Issues](https://github.com/sagoo-cloud/sagooiot/issues) |

> 💡 **Welcome to click ⭐Star⭐ in the upper right corner to support us!**

---

## 📝 Copyright Notice

Open source software is not the same as free. SagooIOT is released under the [GPL-3.0](LICENSE) open source license and provides technical exchange and learning. However, according to this license, modified or derived code from SagooIOT may not be released or sold as closed-source commercial software. If you need to use SagooIOT for any commercial purposes locally, please contact the project manager for commercial licensing to ensure your use complies with the GPL license.

---

## 📋 About SagooIOT

**SagooIOT** is a **lightweight enterprise-grade IoT platform** developed in Go. It provides complete IoT access, management, analysis and application solutions, supporting cross-platform standalone or distributed deployment. You can quickly build a complete IoT business system with device management, data processing, alert notifications, rule engine, video monitoring, edge computing and other functions.

### 🎯 Core Values

- **🚀 Fast Deployment** - Out-of-the-box, start a complete IoT platform in minutes
- **📱 Full Stack Separation** - GoFrame 2.9 backend + Vue 3 frontend, clear and maintainable architecture
- **🔌 Flexible Access** - Support TCP, UDP, HTTP, Websocket, MQTT, CoAP, OPC UA, Modbus, SNMP, IEC104, JT808, GB212 and other protocols
- **🧩 Plugin-Driven** - Unique hot-pluggable plugin system supporting C/C++, Python, Go multi-language development
- **⚡ High-Performance Data Processing** - Integrated TDengine time-series database supporting million-level data points at second-level processing
- **🏭 Edge-Friendly** - Support offline deployment, local rule execution, automatic alerts suitable for edge computing scenarios

### 📌 Quick Information

| Item | Info |
|------|------|
| **Frontend Project** | [sagooiot-ui](https://github.com/sagoo-cloud/sagooiot-ui) |
| **Official Documentation** | http://iotdoc.sagoo.cn/ |
| **Tech Community** | QQ Group: 686637608 |
| **Default Account** | admin / admin123456 |
| **Open Source License** | [GPL-3.0](LICENSE) |

---

## ⚙️ System Requirements

### Minimum Configuration
- **Operating System** - Linux, macOS, Windows
- **Go Version** - 1.23.0 or higher
- **Memory** - 4GB (Recommended 8GB+)
- **Disk Space** - 20GB (Recommended 50GB+)
- **Network** - Stable network connection

### Core Dependencies
| Component | Version | Purpose |
|-----------|---------|---------|
| **MySQL** | 5.7+ / 8.0+ | Relational database for business data storage |
| **PostgreSQL** | 16.x | Relational database for business data storage (Optional) |
| **Redis** | 6.0+ | Cache and message queue |
| **TDengine** | 3.0+ | Time-series database (Optional, for efficient device time-series data storage) |
| **InfluxDB** | 2.x | Time-series database (Optional, for efficient device time-series data storage) |
| **MQTT Broker** | - | Message middleware (Optional, Mosquitto or cloud service recommended) |
| **MinIO** | - | Object storage (Optional, for file management) |

---

## 🚀 Quick Start

### Method One: Direct Run

#### 1. Prerequisites
```bash
# Clone the project
git clone https://github.com/sagoo-cloud/sagooiot.git
cd sagooiot

# Create necessary directories
mkdir -p resource/log

# Initialize database (refer to manifest/sql/)
# Need to import database initialization script in MySQL
```

#### 2. Configure Environment Variables
Create `.env` file in project root or modify `manifest/config/config.dev.yaml`:

```yaml
# Database configuration
database:
  mysql:
    host: localhost
    port: 3306
    name: sagooiot
    user: root
    password: your_password

# Redis configuration
redis:
  host: localhost
  port: 6379
  password: ''

# MQTT configuration (if needed)
mqtt:
  broker: mqtt://localhost:1883
```

#### 3. Start Application
```bash
# Download dependencies
go mod download

# Run application
go run main.go

# Application will start at http://localhost:8000
```

### Method Two: Docker Compose One-Click Start

#### 1. Prepare Docker Environment
```bash
# Ensure Docker and Docker Compose are installed
docker --version
docker-compose --version
```

#### 2. Start Complete Stack
```bash
cd manifest/docker-compose

# Use docker-compose to start all services (MySQL, Redis, SagooIOT, etc.)
docker-compose -f docker-compose.yml up -d

# View logs
docker-compose logs -f sagooiot

# Stop services
docker-compose down
```

#### 3. First Access
- **Web UI** - http://localhost:8000
- **API** - http://localhost:8000/api
- **Default Account** - admin / admin123456

> ⚠️ **Strongly recommended to change default password after first login!**

---

## 🏗️ Platform Architecture

### Overall Architecture Design

```
┌────────────────────────────────────────────────────────────┐
│                   Frontend Layer (Web UI)                  │
│            Vue 3 + Element Plus + TypeScript               │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│              GoFrame 2.9 Application Service Layer          │
│  ┌──────────────┬──────────────┬──────────────────────────┐ │
│  │Device Manager│Data Processing│  Alert & Notification   │ │
│  │Thing Model   │Rule Engine    │  Data Analysis          │ │
│  │Permission    │Data Center    │  Real-time Push         │ │
│  │System Mgmt   │Task Scheduler │  Other Modules          │ │
│  └──────────────┴──────────────┴──────────────────────────┘ │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│            Protocol Access Layer (Multi-Protocol)           │
│  TCP │ MQTT │ UDP │ CoAP │ HTTP │ Websocket │ RPC         │
│              ↓                                              │
│        Plugin System (C/C++/Python/Go)                      │
│  Modbus TCP/RTU/ASCII │ IEC61850 │ OPC UA │ Canopen       │
└────────────────────────────────────────────────────────────┘
                           ↓
┌────────────────────────────────────────────────────────────┐
│         Storage & Message Middleware Layer                  │
│  ┌──────────┬──────────┬──────────┬──────────┐             │
│  │  MySQL   │  Redis   │ TDengine │  MinIO   │             │
│  │Relational│Cache/Queue│TimeSeries│File Store │             │
│  └──────────┴──────────┴──────────┴──────────┘             │
│  MQTT Broker │ Elasticsearch │ Other Storage Components    │
└────────────────────────────────────────────────────────────┘
```

### Core Modules Explanation

| Module | Location | Function | Description |
|--------|----------|----------|-------------|
| **Controller** | `internal/controller/` | API Routing Layer | Receive and respond HTTP requests |
| **Service** | `internal/service/` | Business Logic Layer | Implement business rules and processes |
| **DAO** | `internal/dao/` | Data Access Layer | Database operations and queries |
| **Model** | `internal/model/` | Data Model | Define data structures |
| **Network** | `network/core/` | Protocol Access Core | Multi-protocol adaptation and processing |
| **MQTT** | `internal/mqtt/` | MQTT Message Processing | MQTT pub/sub management |
| **Tasks** | `internal/tasks/` | Background Tasks | Scheduled tasks, async tasks |
| **Workers** | `internal/workers/` | Worker Thread Management | High-concurrency processing and queues |
| **Pkg** | `pkg/` | Tool Library | Cache, MQTT, Plugin, OAuth, etc. |

### Data Flow

```
Device → Protocol Access → Protocol Parse → Data Validation → Business Process → Data Storage
                    ↓
            Real-time Push → WebSocket → Frontend Display
                
                ↓
          Rule Engine → Alert Judgment → Notification Send
              
                ↓
          Data Analysis → Visualization Report
```

---

## ✨ Feature Modules

### A. Device & IoT Core Capabilities

1. **Thing Model Management** - Define device properties, events, services, support JSON format import, flexible adaptation for various devices
2. **Product Management** - Unified management of device type products, support product version, batch operations, product classification
3. **Device Management** - Complete device lifecycle management: registration, configuration, online/offline, activation, disable, delete
4. **Device Tree** - Tree-view display of device relationships and grouping, support multi-dimensional classification and permission control
5. **Device Tags** - Flexible tag system for easy device classification, query, and permission isolation
6. **Real-time Data** - Device real-time status display, historical data query, data export (Excel/CSV)

### B. Protocols & Access Management

1. **Multi-Protocol Adaptation** - Native support for TCP, UDP, HTTP, Websocket, MQTT, CoAP protocols
2. **Protocol Gateway** - Protocol conversion and gateway management, support heterogeneous device access (provide secondary development SDK for quick new protocol development)
3. **Network Tunnel** - Internal network penetration and secure encryption, support P2P connection and proxy forwarding
4. **Plugin System** - Unique hot-pluggable plugin architecture:
   - Support C/C++, Python, Go multi-language writing
   - Cross-process communication (gRPC) ensures isolation
   - Plugin lifecycle management and hot updates
   - Rich plugin examples and documentation
5. **Data Collection Protocols** - Integrated Modbus TCP/RTU/ASCII, IEC61850, OPC UA, IEC104 and other industrial protocols

### C. Data Processing & Analysis

1. **Data Center** - Build custom data models, third-party API integration, database direct connection, data management
2. **Rule Engine** - Visual rule editor, support JavaScript expressions, flexible condition combinations
3. **Data Transformation** - ETL pipeline, data format conversion, field mapping
4. **Real-time Analysis** - Stream data processing, aggregation computation, window functions
5. **Time-Series Database** - Deep integration with TDengine, InfluxDB
6. **Data Export** - Important data support export (Excel, CSV, JSON)

### D. Alerts & Notifications

1. **Alert Rules** - Create, edit, enable/disable rules, support multiple condition triggers (value range, change rate, frequency, etc.)
2. **Alert Levels** - Custom alert levels (critical, important, normal, tip) and color marking
3. **Alert Logs** - Query alert history, statistics and analysis, support time range and condition filtering
4. **Alert Processing** - Alert confirmation, escalation, processing records, processing notes
5. **Notification Templates** - Email, SMS, webhook, DingTalk, WeChat Enterprise multi-channel notifications
6. **Notification Configuration** - Notification channel management, recipient configuration, notification time limits
7. **Message Push** - WebSocket real-time alert push, support multiple push strategies
8. **Alert Statistics** - Alert data visualization reports, support statistics by time, device, level
9. **Alert Governance** - Support alert suppression, duplicate alert merging, alert denoising functions

### E. Edge Computing & Offline Processing

1. **Edge Computing** - Support local deployment and execution at edge side, achieve offline autonomy
2. **Offline Alert** - Automatic alert and notification when device goes offline, support offline duration configuration
3. **Auto Execution** - Automatic device command execution, support delayed execution and retry
4. **Scene Management** - Create and manage application scenarios (multi-device linkage rules), support scenario triggering and orchestration

### F. System Management & Permissions

1. **User Management** - User creation, edit, delete, disable/enable, department assignment, role assignment, password reset
2. **Department Management** - Tree organization structure, department permission isolation, data permission control, department cascade
3. **Role Management** - Role creation, menu permission assignment, button permission assignment, data permission levels
4. **Position Management** - Position configuration, association with users, position permissions
5. **Menu Management** - Dynamic menu configuration, button permission flags, API permission mapping, menu sorting
6. **Permission Management** - RBAC-based permission control, support Casbin advanced permission policies, permission caching
7. **Dictionary Management** - System dictionary maintenance, enum value definition, dictionary classification, dictionary export
8. **Parameter Management** - System parameter dynamic configuration, parameter classification, parameter editing and version management
9. **Organization Management** - Multi-level organization structure support, organization permission isolation, organization-level data aggregation

### G. Logging & Monitoring

1. **Operation Log** - Record user operations (add/delete/update/query), change history, operation details, IP address tracking
2. **Login Log** - Login success/failure records, abnormal login alerts, login IP restriction, login location display
3. **Device Log** - Device upload logs, communication logs, error logs, device command execution logs
4. **System Monitoring** - Real-time monitoring CPU, memory, disk, network, system processes, database connection pool
5. **Online Users** - Real-time user status display, session management, forced logout, online user count statistics
6. **Scheduled Tasks** - Task scheduling and management, execution records, failure retry, task log query

### H. Development & Extension Tools

1. **Code Generation** - Front-end and back-end code auto-generation, CRUD templates, complete API and UI
2. **API Documentation** - Swagger auto-generation, online testing tool, API version management
3. **File Management** - File upload download, file sharing, version control, MinIO integration
4. **Module Extension** - Modular architecture, quick plugin development, SDK provision
5. **Development Examples** - Provide rich development examples and documentation for quick start with secondary development
6. **OpenAPI** - Complete OpenAPI specification support for convenient third-party integration
7. **North-Bound Interface** - Provide standardized north-bound interfaces for integration with other systems

### I. Identity & Authorization

1. **OAuth 2.0** - Third-party OAuth integration (GitHub, WeChat, DingTalk, WeChat Enterprise, etc.)
2. **OAuth Provider** - Act as OAuth provider supporting third-party application access
3. **Login Authentication** - Account password, phone number, email, QR code scan and other authentication methods
4. **Session Management** - gToken JWT token management, session timeout configuration, forced logout, concurrent login control
5. **Audit Trail** - User behavior audit, sensitive operation records, operation traceability
6. **SK/AK** - Support access key based authentication, suitable for API calls

---

## 🛠️ Technology Stack

### Backend Technology Stack

| Technology | Version | Purpose | Advantage |
|------------|---------|---------|-----------|
| **Go** | 1.23.0+ | Core Language | ⚡ High performance, easy concurrency, fast compilation |
| **GoFrame** | 2.9+ | Web Framework | 🎯 Standardized, feature-rich, auto API documentation |
| **MySQL** | 5.7+/8.0+ | Relational Database | 📊 High reliability, widely used, data security |
| **PostgreSQL (Optional)** | 16.x | Relational Database | 📊 Extensible, widely used, data security |
| **Redis** | 6.0+ | Cache/Queue | ⚡ Fast read-write, multiple data structures, persistence |
| **TDengine** | 3.0+ | Time-Series Database | 📈 High-performance time-series storage, auto aggregation |
| **InfluxDB (Optional)** | 2.x | Time-Series Database | 📈 High-performance time-series storage, auto aggregation |
| **MQTT** | Native Support | IoT Protocol | 📡 Lightweight, low bandwidth, widely used |
| **Casbin** | - | Permission Management | 🔐 Flexible RBAC/ABAC, efficient caching |
| **gToken** | - | Session Management | 🎫 JWT token, session timeout, multi-terminal |
| **Asynq** | 0.24+ | Task Queue | ⏳ Redis native, reliable delivery, retry mechanism |
| **Gorilla WebSocket** | - | Real-time Push | 📲 Standard WebSocket, high performance |
| **MinIO** | - | Object Storage | 💾 S3 compatible, self-hosted, high availability |

### Frontend Technology Stack

| Technology | Version | Purpose | Description |
|------------|---------|---------|-------------|
| **Vue** | 3.x | Frontend Framework | 🎨 Modern, declarative, easy to learn |
| **Element Plus** | - | UI Component Library | 🧩 Rich components, unified style, comprehensive docs |
| **TypeScript** | 4.0+ | Type Language | ✅ Type safe, better developer experience, maintainable |
| **Vite** | 2.0+ | Build Tool | ⚡ Lightning-fast cold start, hot-reload fast, minimal output |
| **Pinia** | - | State Management | 📦 Lightweight, TypeScript friendly |
| **Vue Router** | - | Routing Management | 🗂️ SPA, lazy loading, permission control |
| **Axios** | - | HTTP Client | 🌐 Simple API, request intercepting, timeout control |

### Key Dependencies

| Dependency | Function | Use Case |
|------------|----------|----------|
| **paho.mqtt.golang** | MQTT Client | Device access, message processing, pub/sub |
| **gogf/gf** | Go Web Framework | API routing, auto documentation, standardized development |
| **hashicorp/go-plugin** | Cross-process Plugin System | Multi-language plugins, hot-pluggable, isolated execution |
| **taosdata/driver-go** | TDengine Driver | Efficient time-series data storage, auto aggregation |
| **redis/go-redis** | Redis Client | Cache, session storage, message queue |
| **hibiken/asynq** | Task Queue Framework | Scheduled tasks, async processing, failure retry |
| **gorilla/websocket** | WebSocket Server | Real-time push, bidirectional communication, alert push |
| **minio/minio-go** | Object Storage Client | File upload download, distributed storage |
| **golang-jwt/jwt** | JWT Token Library | Authentication, session management, token signing |

---

## 📊 Comparison with Major Open-Source Platforms

| Feature | SagooIOT | EdgeX Foundry | Mainflux | ThingsBoard |
|---------|----------|---------------|----------|-------------|
| **Development Language** | Go | Go | Go | Java |
| **Deployment Complexity** | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Learning Cost** | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| **Thing Model** | ✅ Complete | ✅ Complete | ❌ | ✅ Complete |
| **Plugin System** | ✅ Hot-pluggable (Multi-language) | ✅ Modular | ✅ NATS | ❌ |
| **Edge Computing** | ✅ Native Support | ✅ Native Support | ❌ | ✅ Limited |
| **Protocol Adaptation** | TCP/MQTT/CoAP/HTTP/Websocket | Multiple | MQTT mainly | MQTT/CoAP |
| **Permission Management** | ✅ Casbin + RBAC | ✅ Complete | ✅ Complete | ✅ Complete |
| **Rule Engine** | ✅ Visual + JS | ✅ Yes | ❌ | ✅ Yes |
| **Configuration Tool** | ✅ Visual configuration tool | ❌ | ❌ | ❌ |
| **Data Visualization** | ✅ Data visualization tools | ❌ | ❌ | ❌ |
| **Video Monitoring** | ✅ Stream media service | ❌ | ❌ | ❌ |
| **Data Center** | ✅ Flexible customization | ❌ | ❌ | ❌ |
| **Lightweight** | ✅ High | ❌ Heavy | ✅ High | ❌ Heavy |
| **Secondary Development** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

### 🎯 SagooIOT Core Competitive Advantages

| Advantage | Description |
|-----------|-------------|
| **Ultra-Lightweight** | Full core functionality, low resource consumption, quick deployment for SMEs |
| **Plugin-Driven** | Unique hot-pluggable plugin system, support C/C++/Python/Go unlimited extension |
| **Thing Model First** | Complete thing model support, flexible property/event/service definition |
| **Full-Stack Modern** | Standardized front-end and back-end separation architecture, GoFrame 2.9 and Vue 3 |
| **Quick Customization** | Modular design, quick secondary development and feature extension, lower costs |
| **Development Efficiency** | Auto API documentation generation, code generation tools, quick iteration |
| **Flexible Data Model** | Data center provides flexible data model construction supporting custom business data |
| **Edge-Friendly** | Support offline deployment, local rule execution, automatic alerts suitable for IoT edge scenarios |
| **Data Visualization** | Integrated data visualization tools providing rich charts and visualization features |
| **Distributed Architecture** | Support standalone, distributed, hybrid deployment, meet different scenarios |

---

## 📸 Platform Demo

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

## 📚 Common Application Scenarios

### 1. Industrial Production Monitoring
- Real-time factory equipment monitoring, production data collection
- Equipment maintenance alerts, fault prediction and handling
- Production efficiency analysis and optimization

### 2. Smart Buildings
- Building equipment (HVAC, lighting, security) linkage control
- Energy consumption management and optimization
- Environmental monitoring (temperature, humidity, air quality, etc.)

### 3. Environmental Monitoring
- Real-time environmental data collection and monitoring
- Pollutant alert and alarm
- Environmental data reporting and compliance verification

### 4. Smart Agriculture
- Field sensor deployment and data collection
- Environmental monitoring (temperature, humidity, light, soil moisture)
- Automated irrigation and fertilization control

### 5. New Energy Management
- Photovoltaic, wind, energy storage device monitoring
- Real-time generation statistics and analysis
- Battery management system (BMS) integration

### 6. Vehicle IoT Platform
- Vehicle device access and management
- Vehicle location tracking and route planning
- Fault alert and maintenance reminders

---

## 👥 Community & Contribution

### 🤝 Contribution Guidelines

We welcome all forms of contribution! Whether code, documentation, bug reports or feature suggestions, they all help improve our project.

#### Submit Bug or Feature Suggestion

1. Visit [Issues](https://github.com/sagoo-cloud/sagooiot/issues) page
2. Click "New Issue" button
3. Select appropriate template (Bug report or feature suggestion)
4. Describe the issue or suggestion in detail
5. Submit the issue

#### Submit Code Contribution

1. **Fork the project**
   ```bash
   Click the Fork button in upper right corner
   ```

2. **Clone to local**
   ```bash
   git clone https://github.com/your-username/sagooiot.git
   cd sagooiot
   ```

3. **Create feature branch**
   ```bash
   git checkout -b feature/AmazingFeature
   ```

4. **Make changes and test**
   ```bash
   # Modify code
   # Run tests
   go test ./...
   ```

5. **Commit changes**
   ```bash
   git add .
   git commit -m 'Add some AmazingFeature'
   ```

6. **Push to branch**
   ```bash
   git push origin feature/AmazingFeature
   ```

7. **Open Pull Request**
   - Open PR on GitHub
   - Describe changes and purpose
   - Wait for maintainer review and merge

### 📋 Contribution Standards

- **Code Style** - Follow Go official coding standards, use `gofmt` format code
- **Commit Message** - Use clear, concise commit messages, format: `[type] short description`
  - `[feat]` - New feature
  - `[fix]` - Bug fix
  - `[docs]` - Documentation update
  - `[refactor]` - Code refactoring
  - `[test]` - Test related
  - `[ci]` - CI/CD related
- **Testing** - Ensure all tests pass before submitting PR
- **Documentation** - New features should include corresponding documentation updates
- **Branch Naming** - Use meaningful branch names, e.g., `feature/xxx` or `fix/xxx`

### 🎯 Help We Need

- 📝 **Documentation Translation** - Help translate documentation to other languages
- 🐛 **Bug Fixing** - Report and fix discovered issues
- ✨ **New Features** - Implement new features or improve existing ones
- 📚 **Tutorials & Examples** - Write tutorials and integration examples
- 🌍 **Localization** - Support more languages and regions
- 🎨 **UI/UX Improvement** - Improve user interface and experience

### 📞 Contact & Communication

| Channel | Address | Description |
|---------|---------|-------------|
| **QQ Group** | 686637608 | Real-time tech discussion and issue discussion |
| **GitHub Issues** | [Issues](https://github.com/sagoo-cloud/sagooiot/issues) | Report bugs, feature suggestions |
| **GitHub Discussions** | [Discussions](https://github.com/sagoo-cloud/sagooiot/discussions) | Issue discussion, experience sharing |
| **Official Documentation** | http://iotdoc.sagoo.cn/ | Detailed feature documentation and tutorials |

### 🏆 Acknowledgments

Thanks to all developers and users who contributed to SagooIOT! Your support and feedback drive project development.

---

## 📖 Documentation & Resources

| Resource | Description | Link |
|----------|-------------|------|
| **Official Documentation** | Complete feature documentation and API reference | http://iotdoc.sagoo.cn/ |
| **Frontend Project** | Complete Web UI source code | https://github.com/sagoo-cloud/sagooiot-ui |
| **Cloud Gateway SDK** | Gateway SDK for quick development of private protocols | https://github.com/sagoo-cloud/iotgateway |
| **Cloud Gateway Example** | Example cloud gateway developed based on gateway SDK | https://github.com/sagoo-cloud/iotgateway-example |
| **Plugin Examples** | SagooIOT platform plugin examples | https://github.com/sagoo-cloud/sagooiot-plugins |

---

## 📋 Frequently Asked Questions (FAQ)

### Q: Is SagooIOT free?
**A:** SagooIOT follows GPL-3.0 open source license, personal learning and use are completely free. For commercial licensing, please contact the project team.

### Q: What devices are supported?
**A:** SagooIOT supports access of any devices supporting TCP, MQTT, UDP, CoAP, HTTP, Websocket and other protocols. Additional protocols can be supported through plugins.

### Q: Can performance keep up with large data volumes?
**A:** SagooIOT integrates TDengine time-series database supporting million-level data points at second-level processing. Also supports data sharding and cache optimization solutions.

### Q: Can it be deployed on cloud platforms?
**A:** Completely supported. SagooIOT can be deployed on AWS, Alibaba Cloud, Tencent Cloud, Huawei Cloud and any cloud platform, supporting Docker and Kubernetes.

### Q: How to implement device remote control?
**A:** Through device command function in device management module for remote control. Support real-time and scheduled control.

### Q: Does it support private deployment?
**A:** Fully supported. You can deploy SagooIOT on your own servers or intranet, data belongs completely to you.

### Q: How to implement edge computing?
**A:** SagooIOT supports offline deployment and local rule execution, can run on edge devices achieving edge computing and offline autonomy.

### Q: Can it integrate third-party systems?
**A:** Supported. SagooIOT provides complete OpenAPI and north-bound interfaces for easy integration with third-party systems.

---

## 📄 Disclaimer

SagooIOT Community Edition is an open source learning project not related to commercial activities. Users should comply with laws and regulations when using this project and must not engage in illegal activities.

- If SagooIOT discovers users engaged in illegal activities, it will cooperate with relevant authorities to investigate and report to government
- Users bear all legal liability arising from illegal activities themselves
- If a third party is damaged due to user use, users must compensate accordingly under law
- Users bear all risks in using all resources of SagooIOT
- This project provides no guarantees or warranties

---

## 📈 Project Growth

Thanks for all the support from stars and contributors!

[![Star History Chart](https://star-history.dera.page/svg?repos=sagoo-cloud/sagooiot&type=Date)](https://star-history.dera.page/#sagoo-cloud/sagooiot&Date)

---

## 📜 Open Source License

This project is licensed under [GPL-3.0](LICENSE). See [LICENSE](LICENSE) file for details.

---

## 💬 Contact Us

- **QQ Official Group** - 686637608
- **Official Website** - http://www.sagoo.cn
- **GitHub** - https://github.com/sagoo-cloud/sagooiot
- **Gitee** - https://gitee.com/sagoo-cloud/sagooiot
- **Documentation Center** - http://iotdoc.sagoo.cn/

---

<div align="center">

**Made with ❤️ by the SagooIOT Team**

If this project helps you, please consider clicking ⭐ **Star** in the upper right corner to support us!

</div>

