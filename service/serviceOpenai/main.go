package serviceOpenai

import (
	"errors"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"strings"
	"time"
)

func NewOpenaiByOption() (*openai.Client, error) {
	var optionKeys []models.OpenaiKey
	now := time.Now().UTC().Unix()

	err := global.DB.Where("status = ? AND expiration_time > ?", models.StatusEnabled, now).Find(&optionKeys).Error
	if err != nil {
		return nil, err
	}
	if len(optionKeys) == 0 {
		return nil, errors.New("已经没有可用的key,请联系管理员配置")
	}

	var optionUrl models.Option
	err = global.DB.Where("option_key = ?", models.OptionKeyOpenaiUrl).First(&optionUrl).Error
	if err != nil {
		return nil, err
	}

	if len(optionUrl.OptionValue) == 0 || !strings.HasPrefix(optionUrl.OptionValue, "http") {
		return nil, errors.New("请联系管理员配置正确的url")
	}

	config := openai.DefaultConfig(optionKeys[0].Value)
	config.BaseURL = optionUrl.OptionValue

	client := openai.NewClientWithConfig(config)

	return client, nil

}
