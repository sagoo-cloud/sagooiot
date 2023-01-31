package consts

const (
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	DefaultPageSize          = 10                   //默认分页条数
)

// 系统启动OR禁用常量
const (
	Start    = 1 //启动
	Disabled = 0 //禁用
)

// 权限类型常量
const (
	Menu   = "menu"
	Button = "button"
	Column = "column"
	Api    = "api"
)

// 组态图常量
const (
	FolderTypesTopology   = "topology"   //图纸文件夹类型
	FolderTypesComponents = "components" //自定义组件文件夹类型
)

// WebSocket
const (
	ConfigureDiagram = "configureDiagram" //组态拓扑图
	MonitorServer    = "monitorServer"    //服务监控
)

// 错误代码
const (
	ErrorInvalidFunction = 1   //功能不正确
	ErrorAccessDenied    = 2   //访问被拒绝
	ErrorInvalidData     = 3   //数据无效
	ErrorGetApiData      = 4   //获取接口错误
	ErrorInvalidRole     = 5   //未设置角色
	ErrorNotLogged       = 401 //未登录，或是token失效
)

const (
	Weather        = 1 //天气
	LoopRegulation = 2 //环路监管
	LoopMap        = 3 //分布图
	Energy         = 4 //能耗分析
)

// 服务状态
const (
	ServerStatusOffline = 0
	ServerStatusOnline  = 1
)

// ServerListLimit 服务限制
const (
	ServerListLimit = 10000
)

// 业务单元
const (
	PLOT  = "plot"
	Floor = "floor"
	Unit  = "unit"
)

// 系统参数KEY常量
const (
	IsAutoRunJob        = "sys.auto.run.job"
	IsOpenAccessControl = "sys.access.control"
)

// 默认的插件协议
const (
	DefaultProtocol = "SagooMqtt"
)
