package models

import (
	"go-admin/global"
)

const ( //用户状态
	StatusDisabled = iota
	StatusEnabled
)

func AutoMigrate() {
	global.DB.AutoMigrate(
		&User{},
		&SensitiveWord{},
	)
}
