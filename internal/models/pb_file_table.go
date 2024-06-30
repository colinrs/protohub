package models

import (
	"gorm.io/gorm"
)

type PbFileTableModel struct {
	gorm.Model

	ProjectName string `gorm:"column:project_name;type:varchar(60);default:not null"`
	ServiceName string `gorm:"column:service_name;type:varchar(32);default:not null"`
	Branch      string `gorm:"column:branch;type:varchar(100);default:not null"`
	FileName    string `gorm:"column:file_name;type:varchar(255);default:not null"`
	Sign        string `gorm:"column:sign;type:varchar(255);default:not null"`
	FileID      uint   `gorm:"column:file_id;type:int(11);"`
	Creator     string `gorm:"column:creator;type:varchar(100);default:not null"`
}

func (PbFileTableModel) TableName() string {
	return "pb_file_table"
}

const (
	ProjectNameKey = "project_name"
	ServiceNameKey = "service_name"
	BranchKey      = "branch"
	FileIDKey      = "file_id"
	CreatorKey     = "creator"
)
