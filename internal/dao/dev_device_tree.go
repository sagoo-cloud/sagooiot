// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sagooiot/internal/dao/internal"
)

// internalDevDeviceTreeDao is internal type for wrapping internal DAO implements.
type internalDevDeviceTreeDao = *internal.DevDeviceTreeDao

// devDeviceTreeDao is the data access object for table dev_device_tree.
// You can define custom methods on it to extend its functionality as you wish.
type devDeviceTreeDao struct {
	internalDevDeviceTreeDao
}

var (
	// DevDeviceTree is globally public accessible object for table dev_device_tree operations.
	DevDeviceTree = devDeviceTreeDao{
		internal.NewDevDeviceTreeDao(),
	}
)

// Fill with you ideas below.
