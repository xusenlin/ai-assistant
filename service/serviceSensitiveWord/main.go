package serviceSensitiveWord

import (
	"bufio"
	"go-admin/global"
	"go-admin/models"
	"io"
	"strings"
)

func FindAll() ([]models.SensitiveWord, error) {
	var sensitiveWords []models.SensitiveWord
	err := global.DB.Where("status = ?", models.StatusEnabled).Find(&sensitiveWords).Error
	if err != nil {
		return sensitiveWords, err
	}
	return sensitiveWords, nil
}

func AddAllWord() error {
	words, err := FindAll()

	if err != nil {
		return err
	}
	for _, word := range words {
		global.SensitiveWordsFilter.AddWord(word.Name)
	}
	return nil
}

func Migrate() error {
	var sensitiveWords []models.SensitiveWord
	buf := bufio.NewReader(strings.NewReader(global.SensitiveWords))
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		sensitiveWords = append(sensitiveWords, models.SensitiveWord{
			Name:   string(line),
			Status: models.StatusEnabled,
		})
	}
	if err := global.DB.Create(&sensitiveWords).Error; err != nil {
		return err
	}
	return nil
}
