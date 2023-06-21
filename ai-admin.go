package main

import (
	"fmt"
	"go-admin/global"
	"go-admin/models"
	"go-admin/routes"
	"go-admin/service/serviceSensitiveWord"
)

func main() {

	global.InitParseParams()

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
	err = serviceSensitiveWord.ResetAllWord()
	if err != nil {
		panic(err)
	}

	if err := global.InitTrans("zh"); err != nil {
		fmt.Printf("初始化翻译器失败, err:%v\n", err)
		return
	}

	r := routes.InitRouter()

	r.Run(":" + global.CmdParams.Port)
}
