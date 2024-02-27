package analysis

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ProductCountReq 按年度每月设备消息统计
type ProductCountReq struct {
	g.Meta `path:"/productCount" method:"get" summary:"产品数量统计" tags:"IOT数据分析"`
}
type ProductCountRes struct {
	Total   int `json:"total" dc:"产品总数"`
	Enable  int `json:"enable" dc:"启用产品数"`
	Disable int `json:"disable" dc:"禁用产品数"`
	Added   int `json:"added" dc:"新增产品数"`
}

// DeviceCountForProductReq 获取属于该产品下的设备数量
type DeviceCountForProductReq struct {
	g.Meta     `path:"/deviceCountForProduct" method:"get" summary:"获取属于该产品下的设备数量" tags:"IOT数据分析"`
	ProductKey string `json:"productKey" v:"required#ProductKey，产品key不能为空" dc:"产品key"`
}
type DeviceCountForProductRes struct {
	Number int `json:"number" dc:"设备数"`
}
