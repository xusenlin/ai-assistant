package middlewares

import "C"
import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/service"
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
		claims, err := service.JwtParseToken(token)
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
				"Msg":    "你已经被管理员禁用了",
			})
			c.Abort()
			return
		}

		if !claims.IsAdmin {
			c.JSON(http.StatusOK, gin.H{
				"Status": false,
				"Data":   "",
				"Msg":    "你不是管理员，不允许操作",
			})
			c.Abort()
			return
		}
		c.Set(models.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}

func CustomerAuth(responseJson bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			if responseJson {
				c.JSON(http.StatusOK, gin.H{
					"Status": false,
					"Data":   "",
					"Msg":    "请先登录系统",
				})
			} else {
				c.String(http.StatusBadRequest, "请先登录系统")
			}

			c.Abort()
			return
		}
		claims, err := service.JwtParseToken(token)
		if err != nil {
			if responseJson {
				c.JSON(http.StatusOK, gin.H{
					"Status": false,
					"Data":   "",
					"Msg":    err.Error(),
				})
			} else {
				c.String(http.StatusBadRequest, err.Error())
			}

			c.Abort()
			return
		}
		if claims.Status == models.StatusDisabled {
			if responseJson {
				c.JSON(http.StatusOK, gin.H{
					"Status": false,
					"Data":   "",
					"Msg":    "你已经被管理员禁用了",
				})
			} else {
				c.String(http.StatusBadRequest, "你已经被管理员禁用了")
			}

			c.Abort()
			return
		}
		c.Set(models.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}
