package models

import (
	"go-admin/global"
	"gorm.io/gorm"
	"time"
)

type OpenaiKey struct {
	gorm.Model
	Value          string `binding:"required,min=20,max=88"`
	isCardBound    bool
	ExpirationTime int64
	Status         int `gorm:"default:1"`
}

func (o *OpenaiKey) Destroy(id string) (err error) {
	return global.DB.Where("id = ?", id).Delete(&o).Error
}

func (o *OpenaiKey) Save() (err error) {
	//TODO
	o.ExpirationTime = time.Now().AddDate(0, 0, 100).Unix()
	return global.DB.Save(&o).Error
}
