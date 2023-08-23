// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Remoteconf is the golang structure of table remoteconf for DAO operations like Where/Data.
type Remoteconf struct {
	g.Meta          `orm:"table:remoteconf, do:true"`
	Id              interface{} // 配置ID
	ConfigName      interface{} // 配置名称
	ConfigFormat    interface{} // 配置格式，json等
	ConfigContent   interface{} // 配置内容
	ConfigSize      interface{} // 配置文件大小（按字节计算）
	ProductKey      interface{} // 产品key
	Scope           interface{} // 配置范围：产品=product 设备=device
	Status          interface{} // 状态： 0=停用 1=启用
	ContainedOssUrl interface{} // 包含OssURL
	OssPath         interface{} // Oss文件位置
	OssUrl          interface{} // Oss链接
	Sign            interface{} // 签名
	SignMethod      interface{} // 签名方式，sha256等
	GmtCreate       interface{} // 创建时间
	UtcCreate       *gtime.Time // UTC格式的创建时间
}
