package service

import (
	"go-admin/global"
	"go-admin/models"
	"gorm.io/gorm"
)

func OptionInitOpenai() error {

	err := global.DB.Where("option_key = ?", models.OptionKeyOpenaiUrl).First(&models.Option{}).Error

	if err == gorm.ErrRecordNotFound {
		openaiUrl := models.Option{
			OptionKey:   models.OptionKeyOpenaiUrl,
			OptionValue: "https://api.openai.com/v1",
		}
		global.DB.Error = nil
		err = global.DB.Save(&openaiUrl).Error
		if err != nil {
			return err
		}
	}

	err = global.DB.Where("option_key = ?", models.OptionKeyOpenaiSysPrompt).First(&models.Option{}).Error

	if err == gorm.ErrRecordNotFound {
		openaiKeys := models.Option{
			OptionKey:   models.OptionKeyOpenaiSysPrompt,
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

func OptionGetValue(key string) string {
	var option models.Option
	if err := global.DB.Where("option_key = ?", key).First(&option).Error; err != nil {
		return ""
	}
	return option.OptionValue
}
