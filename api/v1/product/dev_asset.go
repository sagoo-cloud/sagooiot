package product

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetDevAssetListReq 获取数据列表
type GetDevAssetListReq struct {
	g.Meta     `path:"/dev_asset/list" method:"get" summary:"获取档案记录列表" tags:"档案管理"`
	ProductKey string `json:"productKey" dc:"对应产品key"` // 产品key
	common.PaginationReq
}
type GetDevAssetListRes struct {
	Data []GetDevAssetByDevKey
	common.PaginationRes
}

type GetDevAssetByDevKey struct {
	Id             int             `json:"id" v:"required#id必填"`
	Data           []MetaDataValue `json:"data"`
	ProductKey     string          `json:"productKey"          description:"产品key"`
	DeviceName     string          `json:"deviceName"          description:"设备名称"`
	DeviceNumber   string          `json:"deviceNumber"          description:"设备编号"`
	DeviceCategory string          `json:"deviceCategory"          description:"设备类型"`
	DeviceKey      string          `json:"deviceKey"              description:"设备key"`
	InstallTime    string          `json:"installTime"          description:"安装时间"`
	DeptId         string          `json:"deptId"          description:"部门ID"`
	Area           string          `json:"area"          description:"所在区域"`
}

// GetDevAssetByDevKeyReq 获取指定deviceKey的数据
type GetDevAssetByDevKeyReq struct {
	g.Meta    `path:"/dev_asset/get" method:"get" summary:"获取档案记录" tags:"档案管理"`
	DeviceKey string `json:"deviceKey"        description:"设备key" v:"required#deviceKey不能为空"`
}
type GetDevAssetByDevKeyRes struct {
	GetDevAssetByDevKey
}

type MetaDataValue struct {
	ProductKey string `json:"productKey"    description:"产品标识"`
	Name       string `json:"name"          description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"         description:"字段类型"`
	Title      string `json:"title"         description:"字段标题"`
	Value      string `json:"value"         description:"值"`
	FieldName  string `json:"fieldName"          description:"关联字段名称"`
}

// AddDevAssetReq 添加数据
type AddDevAssetReq struct {
	g.Meta         `path:"/dev_asset/add" method:"post" summary:"添加档案记录" tags:"档案管理"`
	Data           []MetaDataValue `json:"data"`
	ProductKey     string          `json:"productKey"  v:"required#产品key不能为空"         description:"产品key"`
	DeviceName     string          `json:"deviceName"     v:"required#设备名不能为空"       description:"设备名称"`
	DeviceNumber   string          `json:"deviceNumber"          description:"设备编号"`
	DeviceCategory string          `json:"deviceCategory"          description:"设备类型"`
	DeviceKey      string          `json:"deviceKey"        v:"required#设备key不能为空"        description:"设备key"`
	InstallTime    string          `json:"installTime"          description:"安装时间"`
	DeptId         string          `json:"deptId"          description:"部门ID"`
	Area           string          `json:"area"          description:"所在区域"`
}

type AddDevAssetRes struct{}

// EditDevAssetReq 编辑数据api
type EditDevAssetReq struct {
	g.Meta         `path:"/dev_asset/edit" method:"put" summary:"编辑档案记录" tags:"档案管理"`
	Id             int             `json:"id" v:"required#id必填"`
	Data           []MetaDataValue `json:"data"`
	ProductKey     string          `json:"productKey"          description:"产品key"`
	DeviceName     string          `json:"deviceName"          description:"设备名称"`
	DeviceNumber   string          `json:"deviceNumber"          description:"设备编号"`
	DeviceCategory string          `json:"deviceCategory"          description:"设备类型"`
	DeviceKey      string          `json:"deviceKey"              description:"设备key"`
	InstallTime    string          `json:"installTime"          description:"安装时间"`
	DeptId         string          `json:"deptId"          description:"部门ID"`
	Area           string          `json:"area"          description:"所在区域"`
}
type EditDevAssetRes struct{}

// DeleteDevAssetReq 删除数据
type DeleteDevAssetReq struct {
	g.Meta `path:"/dev_asset/delete" method:"delete" summary:"删除档案记录" tags:"档案管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteDevAssetRes struct{}
