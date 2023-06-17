package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/service/serviceSensitiveWord"
	"net/http"
)

func MigrateSensitiveWords(c *gin.Context) {
	if err := serviceSensitiveWord.Migrate(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	if err := serviceSensitiveWord.AddAllWord(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   "",
		"Msg":    "sensitive words migration success",
	})
	return
}
