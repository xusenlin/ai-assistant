package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/helper"
	"go-admin/models"
	"go-admin/service"
	"net/http"
	"strconv"
)

type LoginVal struct {
	Username string `binding:"required,min=2,max=18"`
	Password string `binding:"required,min=6,max=16"`
}

func CustomerLogin(c *gin.Context) {

	var val LoginVal
	if err := c.ShouldBindJSON(&val); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.Trans(err).Error(),
		})
		return
	}
	var user models.User
	user.Username = val.Username

	if !user.CustomerExists() {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "当前用户不存在。",
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "当前用户未激活，请先激活。",
		})
		return
	}
	if user.Status == models.StatusDisabled {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "账户已经被管理员禁用了",
		})
		return
	}

	if user.Password != helper.DigestString(models.PasswordSalt+val.Password) {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "密码错误，登录失败",
		})
		return
	}

	token, err := service.JwtGenToken(&models.Claims{
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
		IsAdmin:  user.IsAdmin,
	})

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
		"Data": struct {
			models.User
			Token string
		}{user, token},
		"Msg": "login successful",
	})

}

func CustomerActive(c *gin.Context) {

	var val LoginVal
	if err := c.ShouldBindJSON(&val); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.Trans(err).Error(),
		})
		return
	}
	var user models.User
	user.Username = val.Username

	if !user.CustomerExists() {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "当前用户不存在。",
		})
		return
	}
	if user.Password != "" {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "用户已经激活，不需要重复激活",
		})
		return
	}

	md5 := helper.DigestString(models.PasswordSalt + val.Password)

	err := new(models.User).UpdateField(strconv.Itoa(int(user.ID)), "password", md5)

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
		"Data":   user.ID,
		"Msg":    "success",
	})
}

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
