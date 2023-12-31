package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service"
	"net/http"
	"strconv"
)

func SensitiveWordsList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 20
	}

	keyword := c.Query("keyword")

	var sensitiveWords []models.SensitiveWord

	db := global.DB.Model(&models.SensitiveWord{}).Order("created_at desc")

	if keyword != "" {
		db.Where("name LIKE ?", "%"+keyword+"%")
	}

	r, perr := models.Paginate(db, pageNum, pageSize, &sensitiveWords)
	if perr != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    perr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data": models.SensitiveWordPaginate{
			PageInfo: *r,
			List:     sensitiveWords,
		},
		"Msg": "success",
	})
}

func SensitiveWordsMigrate(c *gin.Context) {
	if err := service.SensitiveWordMigrate(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	if err := service.SensitiveWordReset(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   "",
		"Msg":    "sensitive words migration success",
	})
	return
}

func SensitiveWordsDestroy(c *gin.Context) {
	id := c.Query("id")

	w := new(models.SensitiveWord)
	err := w.Destroy(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	global.SensitiveWordsFilter.DelWord(w.Name)

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   id,
		"Msg":    "",
	})

}

func SensitiveWordsAdd(c *gin.Context) {
	word := c.Query("word")

	if len(word) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "请输入敏感词",
		})
	}

	err := new(models.SensitiveWord).Add(word)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}
	global.SensitiveWordsFilter.AddWord(word)

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   word,
		"Msg":    "",
	})
}
