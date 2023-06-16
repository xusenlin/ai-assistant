package models

import (
	"go-admin/global"
)

func AutoMigrate() {
	global.DB.AutoMigrate(
		&User{},
	)
}
