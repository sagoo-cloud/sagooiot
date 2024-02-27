package consts

const (
	ContextKey = "ContextKey" // 上下文变量存储键名，前后端系统共享
	PageSize   = 10           //默认分页条数
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
	Weather = 1 //天气

)

// 服务状态
const (
	ServerStatusOnline = 1
)

// ServerListLimit 服务限制
const (
	ServerListLimit = 10000
)

const ApiTypes = "api_types"

// 默认的插件协议
const (
	DefaultProtocol = "SagooMqtt"
)

const (
	TokenAuth = "token_auth"
	AKSK      = "aksk"
)

// 文件路径
const (
	LogPath      = "./resource/log/run/"
	RunLogPath   = "./var/"
	MysqlLogPath = "./resource/log/sql/"
)

// RSA 公私钥文件路径
const (
	RsaPublicKeyFile  = "resource/rsa/public.pem"
	RsaPrivateKeyFile = "resource/rsa/private.pem"
	RsaOAEP           = "OAEP"
	RsaPKCS1v15       = "PKCS1v15"
)
