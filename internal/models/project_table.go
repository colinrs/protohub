package models

import (
	"gorm.io/gorm"
)

// Project is the model entity for the Project schema.
type Project struct {
	gorm.Model
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	ProjectStatus uint8 `gorm:"column:project_status" json:"project_status,omitempty"`
	// Project name | 角色名
	ProjectName string `gorm:"column:project_name" json:"project_name,omitempty"`
	// Remark | 备注
	Remark string `gorm:"column:remark" json:"remark,omitempty"`
	// Order number | 排序编号
	Sort uint32 `gorm:"column:sort" json:"sort,omitempty"`
}

func (Project) TableName() string {
	return "projects_table"
}

const (
	ProjectStatusNormal = 1
	ProjectStatusBan    = 2
)
