package serviceOpenai

import (
	"encoding/json"
	"errors"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"strings"
)

func NewOpenaiByOption() (*openai.Client, error) {
	var optionKeys models.Option
	err := global.DB.Where("option_key = ?", models.OptionKeyOpenaiKeys).First(&optionKeys).Error
	if err != nil {
		return nil, err
	}
	if len(optionKeys.OptionValue) == 0 {
		return nil, errors.New("请联系管理员配置key")
	}

	var openaiKeys []string
	err = json.Unmarshal([]byte(optionKeys.OptionValue), &openaiKeys)
	if err != nil {
		return nil, err
	}
	if len(openaiKeys) == 0 {
		return nil, errors.New("请联系管理员配置key")
	}

	var optionUrl models.Option
	err = global.DB.Where("option_key = ?", models.OptionKeyOpenaiUrl).First(&optionUrl).Error
	if err != nil {
		return nil, err
	}
	if len(optionUrl.OptionValue) == 0 || strings.HasPrefix(optionUrl.OptionValue, "http") {
		return nil, errors.New("请联系管理员配置正确的url")
	}

	config := openai.DefaultConfig(openaiKeys[0])
	config.BaseURL = optionUrl.OptionValue

	client := openai.NewClientWithConfig(config)

	return client, nil

}
