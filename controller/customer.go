package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/helper"
	"go-admin/models"
	"go-admin/service"
	"net/http"
	"strconv"
)

func CustomerUpdatePassword(c *gin.Context) {

	type Val struct {
		Password string `json:"password"`
	}
	var v Val
	if err := c.BindJSON(&v); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}
	if len(v.Password) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "密码最少6位",
		})
		return
	}
	claims, err := service.JwtGetClaimsByContext(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	md5 := helper.DigestString(models.PasswordSalt + v.Password)

	err = new(models.User).UpdateField(strconv.Itoa(int(claims.ID)), "password", md5)

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
		"Data":   claims.ID,
		"Msg":    "success",
	})
}

func CustomerGetInfo(c *gin.Context) {
	user, err := service.JwtGetUserByContext(c)

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
		"Data":   user,
		"Msg":    "success",
	})
}
