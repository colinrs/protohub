package models

import (
	"gorm.io/gorm"
)

const (
	IdentityKey    = "identity"
	UserNameKey    = "user_name"
	PasswordKey    = "password"
	EmailKey       = "email"
	UserStatusKey  = "user_status"
	DescriptionKey = "description"
	MobileKey      = "mobile"
)

const (
	RoleCodeKey = "role_code"
)

const (
	ProjectIDKey = "project_id"
)

const (
	UserStatusNormal = 1
	UserStatusBan    = 2
)

type UserTableModel struct {
	gorm.Model
	UserName    string `gorm:"column:user_name;type:varchar(32);default:''"`
	Password    string `gorm:"column:password;type:varchar(256);default:''"`
	Email       string `gorm:"column:email;type:varchar(100);default:''"`
	UserStatus  uint32 `gorm:"column:user_status;type:tinyint(1);default:1"`
	Description string `gorm:"column:description;type:varchar(1000);default:''"`
	Mobile      string `gorm:"column:mobile;type:varchar(30);default:''"`
}

func (UserTableModel) TableName() string {
	return "user_table"
}

type UserRolesTableModel struct {
	gorm.Model
	ProjectID uint   `gorm:"column:project_id"`
	UserID    uint   `gorm:"column:user_id"`
	RoleCode  string `gorm:"column:role_code"`
}

func (UserRolesTableModel) TableName() string {
	return "user_roles_table"
}

type UserProjectTableModel struct {
	gorm.Model
	UserID    uint `gorm:"column:user_id"`
	ProjectID uint `gorm:"column:project_id"`
}

func (UserProjectTableModel) TableName() string {
	return "user_project_table"
}

type UserTokenModel struct {
	gorm.Model
	UserID    uint   `gorm:"column:user_id"`
	Token     string `gorm:"column:token"`
	ExpiredAt int64  `gorm:"column:expired_at"`
}

func (UserTokenModel) TableName() string {
	return "user_token_table"
}
