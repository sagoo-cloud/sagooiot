// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuestbookDao is the data access object for table guestbook.
type GuestbookDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GuestbookColumns // columns contains all the column names of Table for convenient usage.
}

// GuestbookColumns defines and stores column names for table guestbook.
type GuestbookColumns struct {
	Id        string //
	Title     string // 留言标题
	Content   string // 留言内容
	Contacts  string // 联系人
	Telephone string // 联系方式
	CreatedAt string // 留言时间
}

// guestbookColumns holds the columns for table guestbook.
var guestbookColumns = GuestbookColumns{
	Id:        "id",
	Title:     "title",
	Content:   "content",
	Contacts:  "contacts",
	Telephone: "telephone",
	CreatedAt: "created_at",
}

// NewGuestbookDao creates and returns a new DAO object for table data access.
func NewGuestbookDao() *GuestbookDao {
	return &GuestbookDao{
		group:   "default",
		table:   "guestbook",
		columns: guestbookColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuestbookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuestbookDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuestbookDao) Columns() GuestbookColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuestbookDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuestbookDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuestbookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
