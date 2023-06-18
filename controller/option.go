package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"net/http"
)

func OptionGet(c *gin.Context) {
	key := c.Query("key")
	option := new(models.Option)
	err := global.DB.Where("option_key = ?", key).First(&option).Error
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
		"Data":   option,
		"Msg":    "",
	})
}

func OptionSet(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("val")
	option := new(models.Option)
	err := global.DB.Where("option_key = ?", key).First(&option).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	option.OptionValue = val
	err = global.DB.Save(&option).Error
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
		"Data":   option,
		"Msg":    "success",
	})
}
