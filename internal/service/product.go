// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IDevCategory interface {
		Detail(ctx context.Context, id uint) (out *model.ProductCategoryOutput, err error)
		GetNameByIds(ctx context.Context, categoryIds []uint) (names map[uint]string, err error)
		// ListForPage 产品分类列表
		ListForPage(ctx context.Context, page, limit int, name string) (out []*model.ProductCategoryTreeOutput, total int, err error)
		List(ctx context.Context, name string) (out []*model.ProductCategoryTreeOutput, err error)
		Add(ctx context.Context, in *model.AddProductCategoryInput) (err error)
		Edit(ctx context.Context, in *model.EditProductCategoryInput) (err error)
		Del(ctx context.Context, id uint) (err error)
	}
	IDevDataReport interface {
		// Event 设备事件上报
		Event(ctx context.Context, deviceKey string, data model.ReportEventData, subKey ...string) error
		// Property 设备属性上报
		Property(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) error
	}
	IDevDevice interface {
		// Get 获取设备详情
		Get(ctx context.Context, key string) (out *model.DeviceOutput, err error)
		// GetAll 获取所有设备
		GetAll(ctx context.Context) (out []*entity.DevDevice, err error)
		Detail(ctx context.Context, key string) (out *model.DeviceOutput, err error)
		ListForPage(ctx context.Context, in *model.ListDeviceForPageInput) (out *model.ListDeviceForPageOutput, err error)
		// List 已发布产品的设备列表
		List(ctx context.Context, productKey string, keyWord string) (list []*model.DeviceOutput, err error)
		Add(ctx context.Context, in *model.AddDeviceInput) (deviceId uint, err error)
		Edit(ctx context.Context, in *model.EditDeviceInput) (err error)
		// UpdateDeviceStatusInfo 更新设备状态信息，设备上线、离线、注册
		UpdateDeviceStatusInfo(ctx context.Context, deviceKey string, status int, timestamp time.Time) (err error)
		// BatchUpdateDeviceStatusInfo 批量更新设备状态信息，设备上线、离线、注册
		BatchUpdateDeviceStatusInfo(ctx context.Context, deviceStatusLogList []iotModel.DeviceStatusLog) (err error)
		UpdateExtend(ctx context.Context, in *model.DeviceExtendInput) (err error)
		Del(ctx context.Context, keys []string) (err error)
		// Deploy 设备启用
		Deploy(ctx context.Context, key string) (err error)
		// Undeploy 设备禁用
		Undeploy(ctx context.Context, key string) (err error)
		// TotalByProductKey 统计产品下的设备数量
		TotalByProductKey(ctx context.Context, productKeys []string) (totals map[string]int, err error)
		// RunStatus 运行状态
		RunStatus(ctx context.Context, deviceKey string) (out *model.DeviceRunStatusOutput, err error)
		// GetLatestProperty 获取设备最新的属性值
		GetLatestProperty(ctx context.Context, key string) (list []model.DeviceLatestProperty, err error)
		// GetProperty 获取指定属性值
		GetProperty(ctx context.Context, in *model.DeviceGetPropertyInput) (out *model.DevicePropertiy, err error)
		// GetPropertyList 设备属性详情列表
		GetPropertyList(ctx context.Context, in *model.DeviceGetPropertyListInput) (out *model.DeviceGetPropertyListOutput, err error)
		// GetData 获取设备指定日期属性数据
		GetData(ctx context.Context, in *model.DeviceGetDataInput) (list []model.DevicePropertiyOut, err error)
		// BindSubDevice 网关绑定子设备
		BindSubDevice(ctx context.Context, in *model.DeviceBindInput) error
		// UnBindSubDevice 网关解绑子设备
		UnBindSubDevice(ctx context.Context, in *model.DeviceBindInput) error
		// BindList 已绑定列表(分页)
		BindList(ctx context.Context, in *model.DeviceBindListInput) (out *model.DeviceBindListOutput, err error)
		// ListForSub 子设备
		ListForSub(ctx context.Context, in *model.ListForSubInput) (out *model.ListDeviceForPageOutput, err error)
		// CheckBind 检查网关、子设备绑定关系
		CheckBind(ctx context.Context, in *model.CheckBindInput) (bool, error)
		// DelSub 子设备删除
		DelSub(ctx context.Context, key string) (err error)
		// AuthInfo 获取认证信息
		AuthInfo(ctx context.Context, in *model.AuthInfoInput) (*model.AuthInfoOutput, error)
		// GetDeviceOnlineTimeOut 获取设备在线超时时长
		GetDeviceOnlineTimeOut(ctx context.Context, deviceKey string) (timeOut int)
		// ExportDevices 导出设备
		ExportDevices(ctx context.Context, req *product.ExportDevicesReq) (res product.ExportDevicesRes, err error)
		// ImportDevices 导入设备
		ImportDevices(ctx context.Context, req *product.ImportDevicesReq) (res product.ImportDevicesRes, err error)
		SetDevicesStatus(ctx context.Context, req *product.SetDeviceStatusReq) (res product.SetDeviceStatusRes, err error)
		// GetDeviceDataList 获取设备属性聚合数据列表
		GetDeviceDataList(ctx context.Context, in *model.DeviceDataListInput) (out *model.DeviceDataListOutput, err error)
		// GetAllForProduct 获取指定产品所有设备
		GetAllForProduct(ctx context.Context, productKey string) (list []*entity.DevDevice, err error)
		// CacheDeviceDetailList 缓存所有设备详情数据
		CacheDeviceDetailList(ctx context.Context) (err error)
	}
	IDevDeviceFunction interface {
		// Do 执行设备功能
		Do(ctx context.Context, in *model.DeviceFunctionInput) (out *model.DeviceFunctionOutput, err error)
	}
	IDevDeviceLog interface {
		// LogType 日志类型
		LogType(ctx context.Context) (list []string)
		// Search 日志搜索
		Search(ctx context.Context, in *model.DeviceLogSearchInput) (out *model.DeviceLogSearchOutput, err error)
	}
	IDevDeviceProperty interface {
		// Set 设备属性设置
		Set(ctx context.Context, in *model.DevicePropertyInput) (out *model.DevicePropertyOutput, err error)
	}
	IDevDeviceTag interface {
		Add(ctx context.Context, in *model.AddTagDeviceInput) (err error)
		Edit(ctx context.Context, in *model.EditTagDeviceInput) (err error)
		Del(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, deviceId uint, list []model.AddTagDeviceInput) (err error)
	}
	IDevDeviceTree interface {
		// List 设备树列表
		List(ctx context.Context) (out []*model.DeviceTreeListOutput, err error)
		// Change 更换上下级
		Change(ctx context.Context, infoId, parentInfoId int) error
		// Detail 信息详情
		Detail(ctx context.Context, infoId int) (out *model.DetailDeviceTreeInfoOutput, err error)
		// Add 添加设备树基本信息
		Add(ctx context.Context, in *model.AddDeviceTreeInfoInput) error
		// Edit 修改设备树基本信息
		Edit(ctx context.Context, in *model.EditDeviceTreeInfoInput) error
		// Del 删除设备树基本信息
		Del(ctx context.Context, infoId int) error
	}
	IDevInit interface {
		// InitProductForTd 产品表结构初始化
		InitProductForTd(ctx context.Context) (err error)
		// InitDeviceForTd 设备表结构初始化
		InitDeviceForTd(ctx context.Context) (err error)
	}
	IDevProduct interface {
		Detail(ctx context.Context, key string) (out *model.DetailProductOutput, err error)
		GetInfoById(ctx context.Context, id uint) (out *entity.DevProduct, err error)
		GetNameByIds(ctx context.Context, productIds []uint) (names map[uint]string, err error)
		ListForPage(ctx context.Context, in *model.ListForPageInput) (out *model.ListForPageOutput, err error)
		List(ctx context.Context) (list []*model.ProductOutput, err error)
		Add(ctx context.Context, in *model.AddProductInput) (err error)
		Edit(ctx context.Context, in *model.EditProductInput) (err error)
		UpdateExtend(ctx context.Context, in *model.ExtendInput) (err error)
		Del(ctx context.Context, keys []string) (err error)
		// Deploy 产品发布
		Deploy(ctx context.Context, productKey string) (err error)
		// Undeploy 产品停用
		Undeploy(ctx context.Context, productKey string) (err error)
		// ListForSub 子设备类型产品
		ListForSub(ctx context.Context) (list []*model.ProductOutput, err error)
		// UpdateScriptInfo 脚本更新
		UpdateScriptInfo(ctx context.Context, in *model.ScriptInfoInput) (err error)
		// ConnectIntro 获取设备接入信息
		ConnectIntro(ctx context.Context, productKey string) (out *model.DeviceConnectIntroOutput, err error)
	}
	IDevTSLDataType interface {
		DataTypeValueList(ctx context.Context) (out *model.DataTypeOutput, err error)
	}
	IDevTSLEvent interface {
		Detail(ctx context.Context, deviceKey string, eventKey string) (event *model.TSLEvent, err error)
		ListEvent(ctx context.Context, in *model.ListTSLEventInput) (out *model.ListTSLEventOutput, err error)
		AllEvent(ctx context.Context, key string) (list []model.TSLEvent, err error)
		AddEvent(ctx context.Context, in *model.TSLEventAddInput) (err error)
		EditEvent(ctx context.Context, in *model.TSLEventAddInput) (err error)
		DelEvent(ctx context.Context, in *model.DelTSLEventInput) (err error)
	}
	IDevTSLFunction interface {
		ListFunction(ctx context.Context, in *model.ListTSLFunctionInput) (out *model.ListTSLFunctionOutput, err error)
		AllFunction(ctx context.Context, key string, inputsValueTypes string) (list []model.TSLFunction, err error)
		AddFunction(ctx context.Context, in *model.TSLFunctionAddInput) (err error)
		EditFunction(ctx context.Context, in *model.TSLFunctionAddInput) (err error)
		DelFunction(ctx context.Context, in *model.DelTSLFunctionInput) (err error)
	}
	IDevTSLImport interface {
		// Export 导出物模型
		Export(ctx context.Context, key string) (err error)
		// Import 导入物模型
		Import(ctx context.Context, key string, file *ghttp.UploadFile) (err error)
	}
	IDevTSLParse interface {
		// ParseData 基于物模型解析上报数据
		ParseData(ctx context.Context, deviceKey string, data []byte) (res iotModel.ReportPropertyData, err error)
		// HandleProperties 处理属性
		HandleProperties(ctx context.Context, device *model.DeviceOutput, properties map[string]interface{}) (reportDataInfo iotModel.ReportPropertyData, err error)
		// HandleEvents 处理事件上报
		HandleEvents(ctx context.Context, device *model.DeviceOutput, events map[string]sagooProtocol.EventNode) (res []iotModel.ReportEventData, err error)
	}
	IDevTSLProperty interface {
		ListProperty(ctx context.Context, in *model.ListTSLPropertyInput) (out *model.ListTSLPropertyOutput, err error)
		AllProperty(ctx context.Context, key string) (list []model.TSLProperty, err error)
		AddProperty(ctx context.Context, in *model.TSLPropertyInput) (err error)
		EditProperty(ctx context.Context, in *model.TSLPropertyInput) (err error)
		DelProperty(ctx context.Context, in *model.DelTSLPropertyInput) (err error)
	}
	IDevTSLTag interface {
		ListTag(ctx context.Context, in *model.ListTSLTagInput) (out *model.ListTSLTagOutput, err error)
		AddTag(ctx context.Context, in *model.TSLTagInput) (err error)
		EditTag(ctx context.Context, in *model.TSLTagInput) (err error)
		DelTag(ctx context.Context, in *model.DelTSLTagInput) (err error)
	}
)

var (
	localDevCategory       IDevCategory
	localDevDataReport     IDevDataReport
	localDevDevice         IDevDevice
	localDevDeviceFunction IDevDeviceFunction
	localDevDeviceLog      IDevDeviceLog
	localDevDeviceProperty IDevDeviceProperty
	localDevDeviceTag      IDevDeviceTag
	localDevDeviceTree     IDevDeviceTree
	localDevInit           IDevInit
	localDevProduct        IDevProduct
	localDevTSLDataType    IDevTSLDataType
	localDevTSLEvent       IDevTSLEvent
	localDevTSLFunction    IDevTSLFunction
	localDevTSLImport      IDevTSLImport
	localDevTSLParse       IDevTSLParse
	localDevTSLProperty    IDevTSLProperty
	localDevTSLTag         IDevTSLTag
)

func DevCategory() IDevCategory {
	if localDevCategory == nil {
		panic("implement not found for interface IDevCategory, forgot register?")
	}
	return localDevCategory
}

func RegisterDevCategory(i IDevCategory) {
	localDevCategory = i
}

func DevDataReport() IDevDataReport {
	if localDevDataReport == nil {
		panic("implement not found for interface IDevDataReport, forgot register?")
	}
	return localDevDataReport
}

func RegisterDevDataReport(i IDevDataReport) {
	localDevDataReport = i
}

func DevDevice() IDevDevice {
	if localDevDevice == nil {
		panic("implement not found for interface IDevDevice, forgot register?")
	}
	return localDevDevice
}

func RegisterDevDevice(i IDevDevice) {
	localDevDevice = i
}

func DevDeviceFunction() IDevDeviceFunction {
	if localDevDeviceFunction == nil {
		panic("implement not found for interface IDevDeviceFunction, forgot register?")
	}
	return localDevDeviceFunction
}

func RegisterDevDeviceFunction(i IDevDeviceFunction) {
	localDevDeviceFunction = i
}

func DevDeviceLog() IDevDeviceLog {
	if localDevDeviceLog == nil {
		panic("implement not found for interface IDevDeviceLog, forgot register?")
	}
	return localDevDeviceLog
}

func RegisterDevDeviceLog(i IDevDeviceLog) {
	localDevDeviceLog = i
}

func DevDeviceProperty() IDevDeviceProperty {
	if localDevDeviceProperty == nil {
		panic("implement not found for interface IDevDeviceProperty, forgot register?")
	}
	return localDevDeviceProperty
}

func RegisterDevDeviceProperty(i IDevDeviceProperty) {
	localDevDeviceProperty = i
}

func DevDeviceTag() IDevDeviceTag {
	if localDevDeviceTag == nil {
		panic("implement not found for interface IDevDeviceTag, forgot register?")
	}
	return localDevDeviceTag
}

func RegisterDevDeviceTag(i IDevDeviceTag) {
	localDevDeviceTag = i
}

func DevDeviceTree() IDevDeviceTree {
	if localDevDeviceTree == nil {
		panic("implement not found for interface IDevDeviceTree, forgot register?")
	}
	return localDevDeviceTree
}

func RegisterDevDeviceTree(i IDevDeviceTree) {
	localDevDeviceTree = i
}

func DevInit() IDevInit {
	if localDevInit == nil {
		panic("implement not found for interface IDevInit, forgot register?")
	}
	return localDevInit
}

func RegisterDevInit(i IDevInit) {
	localDevInit = i
}

func DevProduct() IDevProduct {
	if localDevProduct == nil {
		panic("implement not found for interface IDevProduct, forgot register?")
	}
	return localDevProduct
}

func RegisterDevProduct(i IDevProduct) {
	localDevProduct = i
}

func DevTSLDataType() IDevTSLDataType {
	if localDevTSLDataType == nil {
		panic("implement not found for interface IDevTSLDataType, forgot register?")
	}
	return localDevTSLDataType
}

func RegisterDevTSLDataType(i IDevTSLDataType) {
	localDevTSLDataType = i
}

func DevTSLEvent() IDevTSLEvent {
	if localDevTSLEvent == nil {
		panic("implement not found for interface IDevTSLEvent, forgot register?")
	}
	return localDevTSLEvent
}

func RegisterDevTSLEvent(i IDevTSLEvent) {
	localDevTSLEvent = i
}

func DevTSLFunction() IDevTSLFunction {
	if localDevTSLFunction == nil {
		panic("implement not found for interface IDevTSLFunction, forgot register?")
	}
	return localDevTSLFunction
}

func RegisterDevTSLFunction(i IDevTSLFunction) {
	localDevTSLFunction = i
}

func DevTSLImport() IDevTSLImport {
	if localDevTSLImport == nil {
		panic("implement not found for interface IDevTSLImport, forgot register?")
	}
	return localDevTSLImport
}

func RegisterDevTSLImport(i IDevTSLImport) {
	localDevTSLImport = i
}

func DevTSLParse() IDevTSLParse {
	if localDevTSLParse == nil {
		panic("implement not found for interface IDevTSLParse, forgot register?")
	}
	return localDevTSLParse
}

func RegisterDevTSLParse(i IDevTSLParse) {
	localDevTSLParse = i
}

func DevTSLProperty() IDevTSLProperty {
	if localDevTSLProperty == nil {
		panic("implement not found for interface IDevTSLProperty, forgot register?")
	}
	return localDevTSLProperty
}

func RegisterDevTSLProperty(i IDevTSLProperty) {
	localDevTSLProperty = i
}

func DevTSLTag() IDevTSLTag {
	if localDevTSLTag == nil {
		panic("implement not found for interface IDevTSLTag, forgot register?")
	}
	return localDevTSLTag
}

func RegisterDevTSLTag(i IDevTSLTag) {
	localDevTSLTag = i
}
