package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
)

const (
	DataSourceFromApi    = iota + 1 // api数据源
	DataSourceFromDb                // 数据库数据源
	DataSourceFromFile              // 文件数据源
	DataSourceFromDevice            // 设备数据源
)

const (
	DataSourceStatusOff int = iota // 数据源未发布
	DataSourceStatusOn             // 数据源已发布
)

const (
	DataSourceDbQueryType    = "tableName" // 数据库源获取数据方式：表
	DataSourceDbQueryTypeSql = "sql"       // 数据库源获取数据方式：sql
)

// 规则配置
type DataSourceRule struct {
	Expression string `json:"expression" dc:"正则表达式"`
	Replace    string `json:"replace" dc:"替换内容"`
}

// 数据源
type DataSource struct {
	Key  string `json:"key" dc:"数据源标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入数据源标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name string `json:"name" dc:"数据源名称" v:"required#请输入数据源名称"`
	Desc string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	From int    `json:"from" dc:"数据来源:1=api导入,2=数据库,3=文件,4=设备" v:"required|in:1,2,3,4#请选择数据来源|未知数据来源，请正确选择"`

	Rule []DataSourceRule `json:"rule" dc:"规则配置"`
}

// api 数据源配置
type DataSourceConfigApi struct {
	Method        string                        `json:"method" dc:"请求方法(get、post、put)"`
	Url           string                        `json:"url" dc:"请求地址" v:"url"`
	RequestParams [][]DataSourceApiRequestParam `json:"requestParams" dc:"请求参数"`

	// 数据更新间隔，cron 格式
	CronExpression string `json:"cronExpression" dc:"任务执行表达式" v:"required#任务表达式不能为空"`
}

// api 请求参数
type DataSourceApiRequestParam struct {
	Type  string `json:"type" dc:"参数类型(header、body、param)" v:"required|in:header,body,param#请选择参数类型|未知参数类型，请正确选择"`
	Key   string `json:"key" dc:"参数标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入参数标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name  string `json:"name" dc:"参数标题"  v:"required#请输入参数标题"`
	Value string `json:"value" dc:"参数值" v:"required#请输入参数值"`
}

// 搜索数据源
type DataSourceSearchInput struct {
	Key  string `json:"key" dc:"数据源标识"`
	Name string `json:"name" dc:"数据源名称"`
	From int    `json:"from" dc:"数据来源" d:"1"`
	PaginationInput
}
type DataSourceSearchOutput struct {
	List []entity.DataSource `json:"list" dc:"数据源列表"`
	PaginationOutput
}

// 添加 api 数据源
type DataSourceApiAddInput struct {
	DataSource

	Config DataSourceConfigApi `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 编辑 api 数据源
type DataSourceApiEditInput struct {
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Name     string `json:"name" dc:"数据源名称" v:"required#请输入数据源名称"`
	Desc     string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Key      string `json:"key" dc:"数据源标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`

	Rule []DataSourceRule `json:"rule" dc:"规则配置"`

	Config DataSourceConfigApi `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 数据源详情
type DataSourceOutput struct {
	*entity.DataSource

	SourceRule []*DataSourceRule `json:"sourceRule" dc:"数据源规则配置"`

	ApiConfig    *DataSourceConfigApi    `json:"apiConfig,omitempty" dc:"api配置"`
	DeviceConfig *DataSourceConfigDevice `json:"deviceConfig,omitempty" dc:"设备配置"`
	DbConfig     *DataSourceConfigDb     `json:"dbConfig,omitempty" dc:"数据库配置"`
}

// 设备 数据源配置
type DataSourceConfigDevice struct {
	ProductKey string `json:"productKey" dc:"产品标识"`
	DeviceKey  string `json:"deviceKey" dc:"设备标识"`
}

// 添加 设备 数据源
type DataSourceDeviceAddInput struct {
	DataSource

	Config DataSourceConfigDevice `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 编辑 设备 数据源
type DataSourceDeviceEditInput struct {
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Name     string `json:"name" dc:"数据源名称" v:"required#请输入数据源名称"`
	Desc     string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Key      string `json:"key" dc:"数据源标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`

	Rule []DataSourceRule `json:"rule" dc:"规则配置"`

	Config DataSourceConfigDevice `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 数据库 数据源配置
type DataSourceConfigDb struct {
	Type   string `json:"type" dc:"数据库类型(mysql/mssql)" v:"required#请配置数据库类型"`
	Host   string `json:"host" dc:"主机" v:"required#请配置主机地址"`
	Port   int    `json:"port" dc:"端口" v:"required#请配置端口号"`
	User   string `json:"user" dc:"用户名" v:"required#请配置用户名"`
	Passwd string `json:"passwd" dc:"密码" v:"required#请配置密码"`
	DbName string `json:"dbName" dc:"数据库名称" v:"required#请配置数据库名称"`

	QueryType string `json:"queryType" dc:"数据获取方式" v:"required|in:tableName,sql#请选择数据获取方式|请正确选择数据获取方式"`
	TableName string `json:"tableName" dc:"表名称" v:"required#请配置表名称或sql语句"`

	Pk  string `json:"pk" dc:"主键字段"`
	Num int    `json:"num" dc:"每次获取数量" d:"100"`

	PkMax uint64 `json:"pkmax" dc:"主键最大值"`

	// 数据更新间隔，cron 格式
	CronExpression string `json:"cronExpression" dc:"任务执行表达式" v:"required#任务表达式不能为空"`
}

// 添加 数据库 数据源
type DataSourceDbAddInput struct {
	DataSource

	Config DataSourceConfigDb `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 编辑 数据库 数据源
type DataSourceDbEditInput struct {
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Name     string `json:"name" dc:"数据源名称" v:"required#请输入数据源名称"`
	Desc     string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Key      string `json:"key" dc:"数据源标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`

	Rule []DataSourceRule `json:"rule" dc:"规则配置"`

	Config DataSourceConfigDb `json:"config" dc:"数据源配置" v:"required#请配置数据源"`
}

// 数据源获取数据的内网方法列表，供大屏使用
type AllSourceOut struct {
	SourceId uint64 `json:"sourceId" dc:"数据源ID"`
	Name     string `json:"name" dc:"数据源名称"`
	Path     string `json:"path" dc:"接口地址"`
}

type SourceDataAllInput struct {
	SourceId uint64                 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Param    map[string]interface{} `json:"param" dc:"搜索哪些字段的数据"`
}
type SourceDataAllOutput struct {
	List string `json:"data" dc:"源数据记录"`
}

type DataSourceDataInput struct {
	SourceId uint64                 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Param    map[string]interface{} `json:"param" dc:"搜索哪些字段的数据"`
	PaginationInput
}
type DataSourceDataOutput struct {
	List string `json:"data" dc:"源数据记录"`
	PaginationOutput
}
