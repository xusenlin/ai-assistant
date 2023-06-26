package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
)

func DialogAddRecord(c *gin.Context, question, answer string) {
	user, err := JwtGetUserByContext(c)
	if err != nil {
		return
	}
	var dialog = models.Dialog{
		Username:    user.Username,
		Question:    question,
		Response:    answer,
		UserIP:      c.Request.RemoteAddr,
		BrowserInfo: c.Request.UserAgent(),
	}
	err = global.DBRecord.Save(&dialog).Error
	fmt.Println(err)

}
