package models

import (
	"gorm.io/gorm"
)

type UserTableModel struct {
	gorm.Model
	Identity     string `gorm:"column:identity;type:varchar(60);default:not null"`
	Name         string `gorm:"column:name;type:varchar(32);default:not null"`
	Password     string `gorm:"column:password;type:varchar(100);default:not null"`
	Email        string `gorm:"column:email;type:varchar(100);default:not null"`
	UserStatus   uint32 `gorm:"column:user_status;type:tinyint(1);default:not null"`
	Description  string `gorm:"column:description;type:varchar(1000);default:not null"`
	Mobile       string `gorm:"column:mobile;type:varchar(30);default:not null"`
	DepartmentId uint64 `gorm:"column:department_id;type:bigint(20);default:not null"`
}

func (UserTableModel) TableName() string {
	return "users_table"
}

const (
	IdentityKey    = "identity"
	UserNameKey    = "name"
	PasswordKey    = "password"
	EmailKey       = "email"
	UserStatusKey  = "user_status"
	DescriptionKey = "description"
	MobileKey      = "mobile"
)

type UserRolesTableModel struct {
	gorm.Model
	RoleCode string `gorm:"column:role_code;type:varchar(24);default:not null"`
}

func (UserRolesTableModel) TableName() string {
	return "users_roles_table"
}

const (
	RoleCodeKey = "role_code"
)

const (
	UserStatusNormal = 1
	UserStatusBan    = 2
)
