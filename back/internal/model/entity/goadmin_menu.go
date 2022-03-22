// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoadminMenu is the golang structure for table goadmin_menu.
type GoadminMenu struct {
	Id         uint        `json:"id"         description:""`
	ParentId   uint        `json:"parentId"   description:""`
	Type       uint        `json:"type"       description:""`
	Order      uint        `json:"order"      description:""`
	Title      string      `json:"title"      description:""`
	Icon       string      `json:"icon"       description:""`
	Uri        string      `json:"uri"        description:""`
	Header     string      `json:"header"     description:""`
	PluginName string      `json:"pluginName" description:""`
	Uuid       string      `json:"uuid"       description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:""`
}