package global

import (
	"flag"
)

type Params struct {
	Port       string
	DBPort     string
	DBName     string
	DBPwd      string
	DBUserName string
}

var CmdParams *Params

func InitParseParams() {
	CmdParams = new(Params)

	flag.StringVar(&CmdParams.Port, "port", "8088", "监听端口号，默认8088")
	flag.StringVar(&CmdParams.DBPort, "dbPort", "3306", "数据库端口号，默认3306")
	flag.StringVar(&CmdParams.DBName, "dbName", "ai-admin", "数据库名字，默认ai-admin")
	flag.StringVar(&CmdParams.DBPwd, "dbPwd", "12345678", "数据库密码，默认12345678")
	flag.StringVar(&CmdParams.DBUserName, "dbUserName", "senlin", "数据库用户，默认senlin")
	flag.Parse()
}
