package models

import (
	"gorm.io/gorm"
)

type SensitiveWord struct {
	gorm.Model
	Name   string `gorm:"type:varchar(200);comment:敏感词" binding:"required,min=1,max=200"`
	Status int    `gorm:"default:2"`
}
