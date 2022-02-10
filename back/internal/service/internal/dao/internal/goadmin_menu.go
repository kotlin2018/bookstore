// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-10 14:06:06
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoadminMenuDao is the data access object for table goadmin_menu.
type GoadminMenuDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns GoadminMenuColumns // columns contains all the column names of Table for convenient usage.
}

// GoadminMenuColumns defines and stores column names for table goadmin_menu.
type GoadminMenuColumns struct {
	Id         string //
	ParentId   string //
	Type       string //
	Order      string //
	Title      string //
	Icon       string //
	Uri        string //
	Header     string //
	PluginName string //
	Uuid       string //
	CreatedAt  string //
	UpdatedAt  string //
}

//  goadminMenuColumns holds the columns for table goadmin_menu.
var goadminMenuColumns = GoadminMenuColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Type:       "type",
	Order:      "order",
	Title:      "title",
	Icon:       "icon",
	Uri:        "uri",
	Header:     "header",
	PluginName: "plugin_name",
	Uuid:       "uuid",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewGoadminMenuDao creates and returns a new DAO object for table data access.
func NewGoadminMenuDao() *GoadminMenuDao {
	return &GoadminMenuDao{
		group:   "default",
		table:   "goadmin_menu",
		columns: goadminMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoadminMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoadminMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoadminMenuDao) Columns() GoadminMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoadminMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoadminMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoadminMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
