package models

import (
	"go-admin/global"
	"gorm.io/gorm"
)

const ( //用户状态
	StatusDisabled = iota
	StatusEnabled
)

type PageInfo struct {
	PageIndex int // 当前页码
	PageSize  int // 每页数据量
	Total     int // 总数据量
}

func Paginate(db *gorm.DB, pageIndex, pageSize int, results interface{}) (pageInfo *PageInfo, err error) {
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	// 计算偏移量
	offset := (pageIndex - 1) * pageSize

	// 查询数据并进行分页
	if err := db.Offset(offset).Limit(pageSize).Find(results).Error; err != nil {
		return nil, err
	}

	pageInfo = &PageInfo{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     int(count),
	}
	return pageInfo, nil
}

func AutoMigrate() error {
	err := global.DB.AutoMigrate(
		&User{},
		&SensitiveWord{},
		&Option{},
		&OpenaiKey{},
	)
	if err != nil {
		return err
	}
	err = global.DBRecord.AutoMigrate(
		&Dialog{},
	)
	if err != nil {
		return err
	}
	return nil
}
