// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceTree is the golang structure of table dev_device_tree for DAO operations like Where/Data.
type DevDeviceTree struct {
	g.Meta       `orm:"table:dev_device_tree, do:true"`
	Id           interface{} //
	InfoId       interface{} // 设备树信息ID
	ParentInfoId interface{} // 父ID
}
