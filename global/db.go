package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error

	dsn := fmt.Sprintf(
		"%v:%v@tcp(127.0.0.1:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		CmdParams.DBUserName, CmdParams.DBPwd, CmdParams.DBPort, CmdParams.DBName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}
