package routes

import (
	"go-admin/controller"
	"go-admin/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode("debug")

	r := gin.Default()
	r.Use(middlewares.Cors())

	{
		r.GET("/ping", func(c *gin.Context) { //服务健康检查
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		//r.Static(config.Cfg.WebsUrl, config.Cfg.WebRootDir)
		//r.Static("/public", config.Cfg.ClientDir)
	}

	v1Public := r.Group("/v1")
	{
		v1Public.POST("/login", controller.UserLogin)
		v1Public.POST("/register", controller.UserRegister)
	}

	v1 := r.Group("/v1").Use(middlewares.AdminAuth())
	{

		////user
		v1.GET("/users", controller.UserFindAll)
		//v1.GET("/user/delete", middlewares.RoleSuperAdmin(), controller.UserDestroy)
		//v1.GET("/user/role_edit", middlewares.RoleSuperAdmin(), controller.UserRoleEdit)

		//InitSensitiveWords
		v1.GET("/migrate_sensitive_words", controller.MigrateSensitiveWords)
	}

	return r
}
