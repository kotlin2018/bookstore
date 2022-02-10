// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-10 14:06:06
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoadminSiteDao is the data access object for table goadmin_site.
type GoadminSiteDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns GoadminSiteColumns // columns contains all the column names of Table for convenient usage.
}

// GoadminSiteColumns defines and stores column names for table goadmin_site.
type GoadminSiteColumns struct {
	Id          string //
	Key         string //
	Value       string //
	Description string //
	State       string //
	CreatedAt   string //
	UpdatedAt   string //
}

//  goadminSiteColumns holds the columns for table goadmin_site.
var goadminSiteColumns = GoadminSiteColumns{
	Id:          "id",
	Key:         "key",
	Value:       "value",
	Description: "description",
	State:       "state",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewGoadminSiteDao creates and returns a new DAO object for table data access.
func NewGoadminSiteDao() *GoadminSiteDao {
	return &GoadminSiteDao{
		group:   "default",
		table:   "goadmin_site",
		columns: goadminSiteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoadminSiteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoadminSiteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoadminSiteDao) Columns() GoadminSiteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoadminSiteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoadminSiteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoadminSiteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}