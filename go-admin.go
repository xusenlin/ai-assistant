package main

import (
	"go-admin/global"
	"go-admin/models"
	"go-admin/routes"
	"go-admin/service/serviceSensitiveWord"
)

func main() {

	global.InitDb()
	err := models.AutoMigrate()
	if err != nil {
		panic(err)
	}

	err = models.InitOpenaiOption()
	if err != nil {
		panic(err)
	}

	global.InitSensitiveWords()
	err = serviceSensitiveWord.AddAllWord()
	if err != nil {
		panic(err)
	}

	r := routes.InitRouter()
	r.Run(":8088")
}
