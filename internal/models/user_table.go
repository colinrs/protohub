package models

import (
	"gorm.io/gorm"
)

type UserTableModel struct {
	gorm.Model

	Identity string `gorm:"column:identity;type:varchar(60);default:not null"`
	Name     string `gorm:"column:name;type:varchar(32);default:not null"`
	Password string `gorm:"column:password;type:varchar(100);default:not null"`
	Email    string `gorm:"column:email;type:varchar(100);default:not null"`
}
