package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"strings"
)

func OpenaiNewClient() (*openai.Client, *models.OpenaiKey, error) {
	var optionKeys []models.OpenaiKey

	err := global.DB.Where("status = ?", models.StatusEnabled).Find(&optionKeys).Error
	if err != nil {
		return nil, nil, err
	}
	if len(optionKeys) == 0 {
		//已经没有可用的key,请联系管理员配置
		return nil, nil, errors.New("AI小加睡着了～呼噜～")
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

func OpenaiPing(c *gin.Context, key string) (string, error) {
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

func OpenaiUpdateUserUsage(u *models.User, answer string) {
	token := 0
	tkm, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
	if err == nil {
		token = len(tkm.Encode(answer, nil, nil))
	}
	u.RemainingDialogueCount = u.RemainingDialogueCount - 1
	u.TokenConsumed = u.TokenConsumed + token
	global.DB.Save(u)
}
