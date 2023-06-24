package models

import (
	"go-admin/global"
	"gorm.io/gorm"
)

type OpenaiKey struct {
	gorm.Model
	Value           string `binding:"required,min=20,max=88"`
	ExceptionReason string `gorm:"type:varchar(1000);comment:异常原因"`
	Status          int    `gorm:"default:1"`
}

func (o *OpenaiKey) Destroy(id string) (err error) {
	return global.DB.Where("id = ?", id).Delete(&o).Error
}

func (o *OpenaiKey) Save() (err error) {
	return global.DB.Save(&o).Error
}
