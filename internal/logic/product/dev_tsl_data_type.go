package product

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type sDevTSLDataType struct{}

func init() {
	service.RegisterDevTSLDataType(devTSLDataTypeNew())
}

func devTSLDataTypeNew() *sDevTSLDataType {
	return &sDevTSLDataType{}
}

func (s *sDevTSLDataType) DataTypeValueList(ctx context.Context) (out *model.DataTypeOutput, err error) {
	tMax := 0
	tString := ""
	tDecimals := 2
	tTrueText := "是"
	tFalseText := "否"
	tTrueValue := true
	tFalseValue := false

	// 基础类型
	baseType := []model.DataTypeValueBase{
		{Title: "int(整数型)", Type: "int", TSLParamBase: model.TSLParamBase{Max: &tMax, Min: &tMax, Unit: &tString}},
		{Title: "long(长整数型)", Type: "long", TSLParamBase: model.TSLParamBase{Max: &tMax, Min: &tMax, Unit: &tString}},
		{Title: "float(单精度浮点型)", Type: "float", TSLParamBase: model.TSLParamBase{Decimals: &tDecimals, Max: &tMax, Min: &tMax, Unit: &tString}},
		{Title: "double(双精度浮点型)", Type: "double", TSLParamBase: model.TSLParamBase{Decimals: &tDecimals, Max: &tMax, Min: &tMax, Unit: &tString}},
		{Title: "text(字符串)", Type: "string", TSLParamBase: model.TSLParamBase{MaxLength: &tMax}},
		{Title: "bool(布尔型)", Type: "boolean", TSLParamBase: model.TSLParamBase{TrueText: &tTrueText, FalseText: &tFalseText, TrueValue: &tTrueValue, FalseValue: &tFalseValue}},
	}

	tEnum := []model.TSLEnumType{
		{Value: "枚举值", Text: "枚举文本"},
	}

	tArray := &model.TSLArrayType{TSLValueType: model.TSLValueType{Type: "int", TSLParam: model.TSLParam{TSLParamBase: model.TSLParamBase{Max: &tMax, Min: &tMax, Unit: &tString}}}}

	tObject := []model.TSLObjectType{
		{Key: "参数标识", Name: "参数名称", Desc: "描述", ValueType: model.TSLValueType{Type: "int", TSLParam: model.TSLParam{TSLParamBase: model.TSLParamBase{Max: &tMax, Min: &tMax, Unit: &tString}}}},
	}

	// 扩展类型
	extensionType := []model.DataTypeValueExtension{
		{Title: "date(2006-01-02 15:04:05或者2006-01-02 15:04:05.000)", Type: "date"},
		{Title: "timestamp(时间戳/毫秒)", Type: "timestamp"},
		{Title: "enum(枚举)", Type: "enum", TSLParamExtension: model.TSLParamExtension{Elements: tEnum}},
		{Title: "array(数组)", Type: "array", TSLParamExtension: model.TSLParamExtension{ElementType: tArray}},
		{Title: "object(结构体)", Type: "object", TSLParamExtension: model.TSLParamExtension{Properties: tObject}},
	}

	out = &model.DataTypeOutput{
		BaseType:      baseType,
		ExtensionType: extensionType,
	}

	return
}
