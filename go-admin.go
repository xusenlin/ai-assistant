package main

import (
	"go-admin/global"
	"go-admin/models"
	"go-admin/routes"
	"go-admin/service/serviceSensitiveWord"
)

func main() {

	global.InitDb()
	models.AutoMigrate()

	global.InitSensitiveWords()
	err := serviceSensitiveWord.AddAllWord()
	if err != nil {
		panic(err)
	}

	r := routes.InitRouter()
	r.Run(":8088")
}
