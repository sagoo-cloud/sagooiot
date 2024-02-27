package consts

// 系统参数KEY常量
const (
	SysUploadFileDomain            = "sys.uploadFile.domain"
	IsAutoRunJob                   = "sys.auto.run.job"
	SysOpenapiSecretkey            = "sys.openapi.secretkey"
	SysMapLngAndLat                = "sys.map.lngAndLat"  //地图中心点经纬度
	SysMapAccessKey                = "sys.map.access.key" //百度地图访问密钥
	SysSystemName                  = "sys.system.name"
	HomePageRoute                  = "homePageRoute"
	SysColumnSwitch                = "sys.column.switch" //列表开关
	SysButtonSwitch                = "sys.button.switch" //按钮开关
	SysApiSwitch                   = "sys.api.switch"    //api开关
	SysSystemCopyright             = "sys.system.copyright"
	SysSystemLogo                  = "sys.system.logo"
	SysSystemLoginPic              = "sys.system.login.pic"
	SysSystemLogoMini              = "sys.system.logo.mini"
	SysIsSingleLogin               = "sys.is.single.login"                 //是否单一登录
	SysTokenExpiryDate             = "sys.token.expiry.date"               //TOKEN过期时间
	SysPasswordChangePeriod        = "sys.password.change.period"          //密码更换周期
	SysPasswordChangePeriodSwitch  = "sys.password.change.period.switch"   //密码更换周期开关
	SysPasswordErrorNum            = "sys.password.error.num"              //密码输入错误次数
	SysAgainLoginDate              = "sys.again.login.date"                //允许再次登录时间
	SysPasswordMinimumLength       = "sys.password.minimum.length"         //密码长度
	SysRequireComplexity           = "sys.require.complexity"              //是否包含复杂字符
	SysRequireDigit                = "sys.require.digit"                   //是否包含数字
	SysRequireLowercaseLetter      = "sys.require.lowercase.letter"        //是否包含小写字母
	SysRequireUppercaseLetter      = "sys.require.uppercase.letter"        //是否包含大写字母
	SysChangePasswordForFirstLogin = "sys.change.password.for.first.login" //首次登录是否更改密码开关
	SysIsSecurityControlEnabled    = "sys.is.security.control.enabled"     //是否启用安全控制
	SysIsRsaEnabled                = "sys.is.rsa.enabled"                  //是否启用RSA
	SYSUPLOADFILEWAY               = "sys.uploadFile.way"                  //文件上传方式
)

// MINIO
const (
	MinioDomain          = "minio.domain"
	MinioApiDomain       = "minio.api.domain"
	MinioAccessKeyId     = "minio.accessKeyId"
	MinioSecretAccessKey = "minio.secretAccessKey"
	MinioUseSsl          = "minio.useSSL"
	MinioBucketName      = "minio.bucketName"
	MinioLocation        = "minio.location"
)

// 设备相关配置的参数
const (
	DeviceDataDelayedStorageTime = "device.data.delayed.storage.time" //延迟落库时间
	DeviceDefaultTimeoutTime     = "device.default.timeout.time"      //设备默认超时时间
)
