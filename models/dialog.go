package models

import (
	"gorm.io/gorm"
)

type Dialog struct {
	gorm.Model
	Username    string `gorm:"type:varchar(20);index;not null;comment:用户名" binding:"required"`
	Question    string `gorm:"type:varchar(1000);comment:提问信息" binding:"required,min=1,max=1000"`
	Response    string `gorm:"type:varchar(1000);comment:回复信息"`
	UserIP      string `gorm:"type:varchar(45);comment:IP"`
	BrowserInfo string `gorm:"type:varchar(100);comment:浏览器信息"`
}
