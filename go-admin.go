package main

import (
	"go-admin/global"
	"go-admin/models"
	"go-admin/routes"
)

func main() {

	global.InitDb()
	models.AutoMigrate()

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	r := routes.InitRouter()
	r.Run(":8088")
}
