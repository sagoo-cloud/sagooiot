package model

const (
	TSLAccessModeDefault  int = iota // 访问类型：读写
	TSLAccessModeReadOnly            // 访问类型：只读
)

const (
	TSLEventLevelDefault int = iota // 事件级别：普通
	TSLEventLevelWarn               // 事件级别：警告
	TSLEventLevelUrgen              // 事件级别：紧急
)

// 基础类型参数
type TSLParamBase struct {
	Max      *int    `json:"max,omitempty" dc:"最大,数字类型:int、long、float、double"`  // 最大,数字类型:int、long、float、double
	Min      *int    `json:"min,omitempty" dc:"最小,数字类型:int、long、float、double"`  // 最小,数字类型:int、long、float、double
	Decimals *int    `json:"decimals,omitempty" dc:"小数位数,数字类型:float、double"`    // 小数位数,数字类型:float、double
	Unit     *string `json:"unit,omitempty" dc:"单位,数字类型:int、long、float、double"` // 单位,数字类型:int、long、float、double

	TrueText   *string `json:"trueText,omitempty" dc:"为true时的文本,默认为'是',布尔类型:bool"`       // 为true时的文本,默认为`是`,布尔类型:bool
	FalseText  *string `json:"falseText,omitempty" dc:"为false时的文本,默认为'否',布尔类型:bool"`     // 为false时的文本,默认为`否`,布尔类型:bool
	TrueValue  *bool   `json:"trueValue,omitempty" dc:"为true时的值,默认为'true',布尔类型:bool"`    // 为true时的值,默认为`true`,布尔类型:bool
	FalseValue *bool   `json:"falseValue,omitempty" dc:"为false时的值,默认为'false',布尔类型:bool"` // 为false时的值,默认为`false`,布尔类型:bool

	MaxLength *int `json:"maxLength,omitempty" dc:"最大长度,字符类型:string"` // 最大长度,字符类型:string
}

// 扩展类型参数
type TSLParamExtension struct {
	// Format      *string         `json:"format,omitempty" dc:"时间类型:date,如:yyyy-MM-dd"` // 时间类型:date,如:yyyy-MM-dd
	Elements    []TSLEnumType   `json:"elements,omitempty" dc:"枚举类型:enum"`     // 枚举类型:enum
	ElementType *TSLArrayType   `json:"elementType,omitempty" dc:"数组类型:array"` // 数组类型:array
	Properties  []TSLObjectType `json:"properties,omitempty" dc:"对象类型:object"` // 对象类型:object
}

// 扩展类型参数:枚举型
type TSLEnumType struct {
	Value string `json:"value" dc:"枚举值"` // 枚举值
	Text  string `json:"text" dc:"枚举文本"` // 枚举文本
}

// 扩展类型参数:数组型
type TSLArrayType struct {
	TSLValueType
}

// 扩展类型参数:对象型
type TSLObjectType struct {
	Key       string       `json:"key" dc:"参数标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string       `json:"name" dc:"参数名称"`
	ValueType TSLValueType `json:"valueType" dc:"参数值"`
	Desc      string       `json:"desc" dc:"描述"`
}

// 类型参数
type TSLParam struct {
	TSLParamBase
	TSLParamExtension
}

// 属性
type TSLProperty struct {
	Key        string       `json:"key" dc:"属性标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入属性标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name       string       `json:"name" dc:"属性名称" v:"required#请输入属性名称"`
	AccessMode int          `json:"accessMode" dc:"属性访问类型:0=读写,1=只读" v:"required#请选择是否只读"`
	ValueType  TSLValueType `json:"valueType" dc:"属性值"`
	Desc       string       `json:"desc" dc:"描述"`
}

// 功能
type TSLFunction struct {
	Key     string              `json:"key" dc:"功能标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入功能标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name    string              `json:"name" dc:"功能名称" v:"required#请输入功能名称"`
	Inputs  []TSLFunctionInput  `json:"inputs" dc:"输入参数"`
	Outputs []TSLFunctionOutput `json:"outputs" dc:"输出参数"`
	Desc    string              `json:"desc" dc:"描述"`
}

// 功能:输入参数
type TSLFunctionInput struct {
	Key       string       `json:"key" dc:"参数标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string       `json:"name" dc:"参数名称"`
	ValueType TSLValueType `json:"valueType" dc:"参数值"`
	Desc      string       `json:"desc" dc:"描述"`
}

// 功能:输出参数
type TSLFunctionOutput struct {
	Key       string       `json:"key" dc:"参数标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string       `json:"name" dc:"参数名称"`
	ValueType TSLValueType `json:"valueType" dc:"参数值"`
	Desc      string       `json:"desc" dc:"描述"`
}

// 事件
type TSLEvent struct {
	Key     string           `json:"key" dc:"事件标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入事件标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name    string           `json:"name" dc:"事件名称" v:"required#请输入事件名称"`
	Level   int              `json:"level" dc:"事件级别:0=普通,1=警告,2=紧急" v:"required#请选择事件级别"`
	Outputs []TSLEventOutput `json:"outputs" dc:"输出参数"`
	Desc    string           `json:"desc" dc:"描述"`
}

// 事件:输入参数
type TSLEventOutput struct {
	Key       string       `json:"key" dc:"参数标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string       `json:"name" dc:"参数名称"`
	ValueType TSLValueType `json:"valueType" dc:"参数值"`
	Desc      string       `json:"desc" dc:"描述"`
}

// 标签
type TSLTag struct {
	Key        string       `json:"key" dc:"标签标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入标签标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name       string       `json:"name" dc:"标签名称" v:"required#请输入标签名称"`
	AccessMode int          `json:"accessMode" dc:"标签访问类型:0=读写,1=只读" v:"required#请选择是否只读"`
	ValueType  TSLValueType `json:"valueType" dc:"标签值"`
	Desc       string       `json:"desc" dc:"描述"`
}

// 物模型
type TSL struct {
	Key        string        `json:"key" dc:"产品标识"`
	Name       string        `json:"name" dc:"产品名称"`
	Properties []TSLProperty `json:"properties" dc:"属性"`
	Functions  []TSLFunction `json:"functions" dc:"功能"`
	Events     []TSLEvent    `json:"events" dc:"事件"`
	Tags       []TSLTag      `json:"tags" dc:"标签"`
}
