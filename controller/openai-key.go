package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"net/http"
	"time"
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

	for i, k := range keys {
		if k.ExpirationTime < time.Now().UTC().Unix() {
			keys[i].Status = models.StatusDisabled
		}
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
