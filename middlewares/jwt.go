package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/service/serviceUser"
	"net/http"
)

func AdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "请先登录系统",
			})
			c.Abort()
			return
		}
		claims, err := serviceUser.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Status == models.StatusDisabled {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "You have been banned from logging in",
			})
			c.Abort()
			return
		}

		if !claims.IsAdmin {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "You are not an administrator, access is not allowed",
			})
			c.Abort()
			return
		}
		c.Set(models.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}

func CustomerAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "请先登录系统",
			})
			c.Abort()
			return
		}
		claims, err := serviceUser.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Status == models.StatusDisabled {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "You have been banned from logging in",
			})
			c.Abort()
			return
		}
		c.Set(models.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}
