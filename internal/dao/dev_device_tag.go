// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sagooiot/internal/dao/internal"
)

// internalDevDeviceTagDao is internal type for wrapping internal DAO implements.
type internalDevDeviceTagDao = *internal.DevDeviceTagDao

// devDeviceTagDao is the data access object for table dev_device_tag.
// You can define custom methods on it to extend its functionality as you wish.
type devDeviceTagDao struct {
	internalDevDeviceTagDao
}

var (
	// DevDeviceTag is globally public accessible object for table dev_device_tag operations.
	DevDeviceTag = devDeviceTagDao{
		internal.NewDevDeviceTagDao(),
	}
)

// Fill with you ideas below.
