// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MerchantDao is the data access object for table merchant.
type MerchantDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns MerchantColumns // columns contains all the column names of Table for convenient usage.
}

// MerchantColumns defines and stores column names for table merchant.
type MerchantColumns struct {
	Id         string //
	Name       string //
	Address    string //
	Mobile     string //
	CreateTime string //
	UpdateTime string //
}

// merchantColumns holds the columns for table merchant.
var merchantColumns = MerchantColumns{
	Id:         "id",
	Name:       "name",
	Address:    "address",
	Mobile:     "mobile",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewMerchantDao creates and returns a new DAO object for table data access.
func NewMerchantDao() *MerchantDao {
	return &MerchantDao{
		group:   "default",
		table:   "merchant",
		columns: merchantColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MerchantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MerchantDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MerchantDao) Columns() MerchantColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MerchantDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MerchantDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MerchantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
