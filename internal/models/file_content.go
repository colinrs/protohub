package models

import "gorm.io/gorm"

type FileContentTableModel struct {
	gorm.Model

	FileContent string `gorm:"column:file_content;type:mediumtext;default:null"`
	Creator     string `gorm:"column:creator;type:varchar(100);default:not null"`
}

func (FileContentTableModel) TableName() string {
	return "file_content_table"
}
