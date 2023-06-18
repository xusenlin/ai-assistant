package models

import (
	"go-admin/global"
	"gorm.io/gorm"
)

const OptionKeyOpenaiUrl = "openai_url"
const OptionKeyOpenaiKeys = "openai_keys"

type Option struct {
	gorm.Model
	OptionKey   string `gorm:"type:varchar(45);unique;comment:OptionKey" binding:"required,min=2,max=18"`
	OptionValue string `gorm:"type:varchar(1000);comment:OptionValue"`
}

func InitOpenaiOption() error {

	err := global.DB.Where("option_key = ?", OptionKeyOpenaiUrl).First(&Option{}).Error

	if err != gorm.ErrRecordNotFound {
		return err
	}

	global.DB.Error = nil

	openaiUrl := Option{
		OptionKey:   OptionKeyOpenaiUrl,
		OptionValue: "https://api.openai.com/v1",
	}
	err = global.DB.Save(&openaiUrl).Error
	if err != nil {
		return err
	}

	err = global.DB.Where("option_key = ?", OptionKeyOpenaiKeys).First(&Option{}).Error

	if err != gorm.ErrRecordNotFound {
		return err
	}

	openaiKeys := Option{
		OptionKey:   OptionKeyOpenaiKeys,
		OptionValue: "",
	}
	err = global.DB.Save(&openaiKeys).Error
	if err != nil {
		return err
	}

	return nil
}
