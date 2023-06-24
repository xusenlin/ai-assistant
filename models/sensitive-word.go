package models

import (
	"errors"
	"go-admin/global"
	"gorm.io/gorm"
)

type SensitiveWordPaginate struct {
	PageInfo
	List []SensitiveWord
}

type SensitiveWord struct {
	gorm.Model
	Name   string `gorm:"type:varchar(200);comment:敏感词" binding:"required,min=1,max=200"`
	Status int    `gorm:"default:2"`
}

func (s *SensitiveWord) Destroy(id string) (err error) {
	return global.DB.First(s, id).Delete(&s).Error
}

func (s *SensitiveWord) Add(word string) error {

	if global.DB.Where("name = ?", word).First(&s).Error != gorm.ErrRecordNotFound {
		return errors.New("敏感词已经存在")
	}
	global.DB.Error = nil

	s.Name = word
	return global.DB.Save(&s).Error
}
