package models

import (
	"time"
)

// Role is the model entity for the Role schema.
type Role struct {
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Role name | 角色名
	Name string `json:"name,omitempty"`
	// Role code for permission control in front end | 角色码，用于前端权限控制
	Code string `json:"code,omitempty"`
	// Default menu : dashboard | 默认登录页面
	DefaultRouter string `json:"default_router,omitempty"`
	// Remark | 备注
	Remark string `json:"remark,omitempty"`
	// Order number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
}

func (Role) TableName() string {
	return "role_table"
}
