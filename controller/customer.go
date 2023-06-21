package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service/serviceOpenai"
	"go-admin/service/serviceUser"
	"io"
	"net/http"
)

func GPT3Dot5Turbo(c *gin.Context) {
	var chat []openai.ChatCompletionMessage

	if err := c.BindJSON(&chat); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	chatLen := len(chat)
	if chatLen == 0 {
		c.String(http.StatusBadRequest, "没有对话内容")
		return
	}

	lastQuestion := chat[chatLen-1].Content

	ok, word := global.SensitiveWordsFilter.Validate(lastQuestion)
	if !ok {
		c.String(http.StatusBadRequest, "提问不能包含敏感词'"+word+"'")
		return
	}

	user, err := serviceUser.GetUserByContext(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if user.RemainingDialogueCount <= 0 {
		c.String(http.StatusBadRequest, "你的剩余对话次数不足")
		return
	}

	client, openaiErr := serviceOpenai.NewOpenaiByOption()
	if openaiErr != nil {
		c.String(http.StatusBadRequest, openaiErr.Error())
		return
	}

	var prompt models.Option
	if err := global.DB.Where("option_key = ?", models.OptionKeyOpenaiSysPrompt).First(&prompt).Error; err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	if prompt.OptionValue != "" {
		systemPrompt := openai.ChatCompletionMessage{
			Role:    "system",
			Content: prompt.OptionValue,
		}
		chat = append([]openai.ChatCompletionMessage{systemPrompt}, chat...)
	}

	stream, clientErr := client.CreateChatCompletionStream(c, openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Stream:    true,
		Messages:  chat,
	})
	if clientErr != nil {
		c.String(http.StatusBadRequest, clientErr.Error())
		return
	}
	defer stream.Close()
	var lastAnswer = ""
	defer func() {
		token := 0
		tkm, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
		if err == nil {
			token = len(tkm.Encode(lastAnswer, nil, nil))
		}
		user.RemainingDialogueCount = user.RemainingDialogueCount - 1
		user.TokenConsumed = user.TokenConsumed + token
		global.DB.Save(user)
	}()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			fmt.Fprint(c.Writer, err.Error())
			c.Writer.Flush()
			return
		}

		content := response.Choices[0].Delta.Content
		fmt.Fprint(c.Writer, content)
		lastAnswer = lastAnswer + content
		c.Writer.Flush()
	}
}
