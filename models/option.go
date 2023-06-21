package models

import (
	"go-admin/global"
	"gorm.io/gorm"
)

const OptionKeyOpenaiUrl = "openai_url"
const OptionKeyOpenaiSysPrompt = "openai_sys_prompt"

type Option struct {
	gorm.Model
	OptionKey   string `gorm:"type:varchar(45);unique;comment:OptionKey" binding:"required,min=2,max=18"`
	OptionValue string `gorm:"type:varchar(1000);comment:OptionValue"`
}

func InitOpenaiOption() error {

	err := global.DB.Where("option_key = ?", OptionKeyOpenaiUrl).First(&Option{}).Error

	if err == gorm.ErrRecordNotFound {
		openaiUrl := Option{
			OptionKey:   OptionKeyOpenaiUrl,
			OptionValue: "https://api.openai.com/v1",
		}
		global.DB.Error = nil
		err = global.DB.Save(&openaiUrl).Error
		if err != nil {
			return err
		}
	}

	err = global.DB.Where("option_key = ?", OptionKeyOpenaiSysPrompt).First(&Option{}).Error

	if err == gorm.ErrRecordNotFound {
		openaiKeys := Option{
			OptionKey:   OptionKeyOpenaiSysPrompt,
			OptionValue: "你是一款智能ai助手，你的名字叫 'ai小护',专为护士提供各种问题的解答。",
		}
		global.DB.Error = nil
		err = global.DB.Save(&openaiKeys).Error
		if err != nil {
			return err
		}
	}

	return nil
}
