// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/sagoo-cloud/sagooiot/internal/dao/internal"
)

// internalDataTemplateBusiDao is internal type for wrapping internal DAO implements.
type internalDataTemplateBusiDao = *internal.DataTemplateBusiDao

// dataTemplateBusiDao is the data access object for table data_template_busi.
// You can define custom methods on it to extend its functionality as you wish.
type dataTemplateBusiDao struct {
	internalDataTemplateBusiDao
}

var (
	// DataTemplateBusi is globally public accessible object for table data_template_busi operations.
	DataTemplateBusi = dataTemplateBusiDao{
		internal.NewDataTemplateBusiDao(),
	}
)

// Fill with you ideas below.
