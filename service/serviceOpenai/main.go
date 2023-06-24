package serviceOpenai

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"strings"
)

func NewOpenaiByOption() (*openai.Client, *models.OpenaiKey, error) {
	var optionKeys []models.OpenaiKey

	err := global.DB.Where("status = ?", models.StatusEnabled).Find(&optionKeys).Error
	if err != nil {
		return nil, nil, err
	}
	if len(optionKeys) == 0 {
		return nil, nil, errors.New("已经没有可用的key,请联系管理员配置")
	}

	var optionUrl models.Option
	err = global.DB.Where("option_key = ?", models.OptionKeyOpenaiUrl).First(&optionUrl).Error
	if err != nil {
		return nil, nil, err
	}

	if len(optionUrl.OptionValue) == 0 || !strings.HasPrefix(optionUrl.OptionValue, "http") {
		return nil, nil, errors.New("请联系管理员配置正确的url")
	}

	config := openai.DefaultConfig(optionKeys[0].Value)
	config.BaseURL = optionUrl.OptionValue

	client := openai.NewClientWithConfig(config)

	return client, &optionKeys[0], nil

}

func Ping(c *gin.Context, key string) (string, error) {
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		c,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "ping!",
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
