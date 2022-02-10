// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-10 14:06:06
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoadminSessionDao is the data access object for table goadmin_session.
type GoadminSessionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns GoadminSessionColumns // columns contains all the column names of Table for convenient usage.
}

// GoadminSessionColumns defines and stores column names for table goadmin_session.
type GoadminSessionColumns struct {
	Id        string //
	Sid       string //
	Values    string //
	CreatedAt string //
	UpdatedAt string //
}

//  goadminSessionColumns holds the columns for table goadmin_session.
var goadminSessionColumns = GoadminSessionColumns{
	Id:        "id",
	Sid:       "sid",
	Values:    "values",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewGoadminSessionDao creates and returns a new DAO object for table data access.
func NewGoadminSessionDao() *GoadminSessionDao {
	return &GoadminSessionDao{
		group:   "default",
		table:   "goadmin_session",
		columns: goadminSessionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoadminSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoadminSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoadminSessionDao) Columns() GoadminSessionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoadminSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoadminSessionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoadminSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
