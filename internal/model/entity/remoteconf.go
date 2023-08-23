// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Remoteconf is the golang structure for table remoteconf.
type Remoteconf struct {
	Id              string      `json:"id"              description:"配置ID"`                      // 配置ID
	ConfigName      string      `json:"configName"      description:"配置名称"`                      // 配置名称
	ConfigFormat    string      `json:"configFormat"    description:"配置格式，json等"`                // 配置格式，json等
	ConfigContent   string      `json:"configContent"   description:"配置内容"`                      // 配置内容
	ConfigSize      int         `json:"configSize"      description:"配置文件大小（按字节计算）"`             // 配置文件大小（按字节计算）
	ProductKey      string      `json:"productKey"      description:"产品key"`                     // 产品key
	Scope           string      `json:"scope"           description:"配置范围：产品=product 设备=device"` // 配置范围：产品=product 设备=device
	Status          int         `json:"status"          description:"状态： 0=停用 1=启用"`             // 状态： 0=停用 1=启用
	ContainedOssUrl int         `json:"containedOssUrl" description:"包含OssURL"`                  // 包含OssURL
	OssPath         string      `json:"ossPath"         description:"Oss文件位置"`                   // Oss文件位置
	OssUrl          string      `json:"ossUrl"          description:"Oss链接"`                     // Oss链接
	Sign            string      `json:"sign"            description:"签名"`                        // 签名
	SignMethod      string      `json:"signMethod"      description:"签名方式，sha256等"`              // 签名方式，sha256等
	GmtCreate       string      `json:"gmtCreate"       description:"创建时间"`                      // 创建时间
	UtcCreate       *gtime.Time `json:"utcCreate"       description:"UTC格式的创建时间"`                // UTC格式的创建时间
}
