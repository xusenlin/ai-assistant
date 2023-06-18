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
		r.Static("/admin", "frontendAdmin/dist")
		r.Static("/customer", "frontendCustomer/dist")

		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/customer")
		})
	}

	v1Public := r.Group("/v1")
	{
		v1Public.POST("/login", controller.UserLogin)
		v1Public.POST("/register", controller.UserRegister)
	}

	v1 := r.Group("/v1").Use(middlewares.AdminAuth())
	{

		//user
		v1.GET("/users/list", controller.UsersFind)
		v1.GET("/user/destroy", controller.UserDestroy)
		v1.GET("/user/updateStatus", controller.UserUpdateStatus)
		v1.GET("/user/updatePassword", controller.UserUpdatePassword)
		v1.GET("/user/updateRemainingDialogueCount", controller.UserUpdateRemainingDialogueCount)
		//option
		v1.GET("/option/get", controller.OptionGet)
		v1.GET("/option/set", controller.OptionSet)
		//SensitiveWords
		v1.GET("/sensitiveWords/migrate", controller.SensitiveWordsMigrate)
		v1.GET("/sensitiveWords/list", controller.SensitiveWordsList)
		v1.GET("/sensitiveWords/destroy", controller.SensitiveWordsDestroy)
		v1.GET("/sensitiveWords/add", controller.SensitiveWordsAdd)

	}

	customer := r.Group("/api").Use(middlewares.CustomerAuth())
	{
		customer.GET("/gpt", controller.UsersFind)
	}

	return r
}
