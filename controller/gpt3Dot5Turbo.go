package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service"
	"io"
	"net/http"
	"strings"
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

	//if chatLen > 10 {
	//	c.String(http.StatusBadRequest, "对话已经超出限制,请重置对话")
	//	return
	//}

	lastQuestion := chat[chatLen-1].Content

	ok, word := global.SensitiveWordsFilter.Validate(lastQuestion)
	if !ok {
		c.String(http.StatusBadRequest, "提问不能包含敏感词'"+word+"'")
		return
	}

	user, err := service.JwtGetUserByContext(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if user.RemainingDialogueCount <= 0 {
		c.String(http.StatusBadRequest, "你的剩余对话次数不足")
		return
	}

	sysPrompt := service.OptionGetValue(models.OptionKeyOpenaiSysPrompt)
	if sysPrompt != "" {
		systemPrompt := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: sysPrompt,
		}
		chat = append([]openai.ChatCompletionMessage{systemPrompt}, chat...)
	}

RetryCount:
	client, key, openaiErr := service.OpenaiNewClient()
	if openaiErr != nil {
		c.String(http.StatusBadRequest, openaiErr.Error())
		return
	}

	stream, clientErr := client.CreateChatCompletionStream(c, openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo16K,
		MaxTokens: 2048,
		Stream:    true,
		Messages:  chat,
	})
	if clientErr != nil {

		if strings.Contains(clientErr.Error(), "Please reduce the length of the messages or completion") {
			c.String(http.StatusBadRequest, "对话已经超出限制,请重置对话")
			return
		}
		if strings.Contains(clientErr.Error(), "You can retry your request") {
			// That model is currently overloaded with other requests. You can retry your request, or contact us through our help center at help.openai.com if the error persists.
			c.String(http.StatusBadRequest, "AI小加正在全力解答大家的问题，请稍后再试一下吧～")
			return
		}

		key.Status = models.StatusDisabled
		key.ExceptionReason = clientErr.Error()
		if err := global.DB.Save(key).Error; err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		goto RetryCount

	}
	defer stream.Close()
	var lastAnswer = ""
	defer func() {
		service.OpenaiUpdateUserUsage(user, lastAnswer)
		service.DialogAddRecord(c, lastQuestion, lastAnswer)
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
