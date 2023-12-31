package service

import (
	"bufio"
	"errors"
	"github.com/importcjj/sensitive"
	"go-admin/global"
	"go-admin/models"
	"io"
	"strings"
)

func SensitiveWordFind() ([]models.SensitiveWord, error) {
	var sensitiveWords []models.SensitiveWord
	err := global.DB.Where("status = ?", models.StatusEnabled).Find(&sensitiveWords).Error
	if err != nil {
		return sensitiveWords, err
	}
	return sensitiveWords, nil
}

func SensitiveWordReset() error {
	words, err := SensitiveWordFind()

	if err != nil {
		return err
	}

	global.SensitiveWordsFilter = sensitive.New()

	for _, word := range words {
		global.SensitiveWordsFilter.AddWord(word.Name)
	}
	return nil
}

func SensitiveWordMigrate() error {
	var count int64

	if err := global.DB.Model(&models.SensitiveWord{}).Count(&count).Error; err != nil {
		return err
	}
	if count != 0 {
		return errors.New("敏感词迁移需要运维人员删除表的全部数据才能重新迁移")
	}

	buf := bufio.NewReader(strings.NewReader(global.SensitiveWords))

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		m := models.SensitiveWord{
			Name:   string(line),
			Status: models.StatusEnabled,
		}

		if err := global.DB.Create(&m).Error; err != nil {
			return err
		}

	}

	return nil
}
