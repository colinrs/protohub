package models

import (
	"gorm.io/gorm"
)

// Role is the model entity for the Role schema.
type Role struct {
	gorm.Model

	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	RoleStatus uint8 `gorm:"column:role_status" json:"role_status,omitempty"`
	// Role name | 角色名
	Name string `gorm:"column:file_name" json:"name,omitempty"`
	// Role code for permission control in front end | 角色码，用于前端权限控制
	Code string `gorm:"column:code" json:"code,omitempty"`
	// Remark | 备注
	Remark string `gorm:"column:remark" json:"remark,omitempty"`
	// Order number | 排序编号
	Sort uint32 `gorm:"column:sort" json:"sort,omitempty"`
}

func (Role) TableName() string {
	return "roles_table"
}

const (
	StatusKey = "status"
	NameKey   = "name"
	CodeKey   = "code"
	RemarkKey = "remark"
	SortKey   = "sort"
)

const (
	StatusNormal = 1
	StatusBan    = 2
)
