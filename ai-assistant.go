package main

import (
	"fmt"
	"go-admin/global"
	"go-admin/models"
	"go-admin/routes"
	"go-admin/service"
)

func main() {

	global.InitParseParams()

	global.InitDb()
	err := models.AutoMigrate()
	if err != nil {
		panic(err)
	}

	err = service.OptionInitOpenai()
	if err != nil {
		panic(err)
	}

	global.InitSensitiveWords()
	err = service.SensitiveWordReset()
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
