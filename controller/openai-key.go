package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service/serviceOpenai"
	"net/http"
)

func OpenaiKeyFindAll(c *gin.Context) {
	var keys []models.OpenaiKey

	err := global.DB.Find(&keys).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   keys,
		"Msg":    "success",
	})
}

func OpenaiKeyDestroy(c *gin.Context) {
	id := c.Query("id")

	err := new(models.OpenaiKey).Destroy(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   id,
		"Msg":    "",
	})
}

func OpenaiKeyAdd(c *gin.Context) {
	var key models.OpenaiKey
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.Trans(err).Error(),
		})
		return
	}

	_, err := serviceOpenai.Ping(c, key.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}
	if err := key.Save(); err != nil {
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
		"Msg":    "success",
	})
}

func OpenaiKeyPing(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "查询的key不能为空",
		})
		return
	}

	var openaiKey models.OpenaiKey

	if err := global.DB.Where("value = ?", key).First(&openaiKey).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	test, err := serviceOpenai.Ping(c, openaiKey.Value)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		openaiKey.Status = models.StatusDisabled
		openaiKey.ExceptionReason = err.Error()
		global.DB.Save(openaiKey)
		return
	}

	openaiKey.Status = models.StatusEnabled
	global.DB.Save(openaiKey)

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   test,
		"Msg":    "success",
	})

}