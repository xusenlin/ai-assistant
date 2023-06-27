package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"net/http"
	"strconv"
)

func DialogList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 20
	}

	keyword := c.Query("keyword")

	var dialogs []models.Dialog

	db := global.DBRecord.Model(&models.Dialog{}).Order("created_at desc")

	if keyword != "" {
		db.Where("username LIKE ?", "%"+keyword+"%")
	}

	r, err := models.Paginate(db, pageNum, pageSize, &dialogs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data": models.DialogPaginate{
			PageInfo: *r,
			List:     dialogs,
		},
		"Msg": "success",
	})
}
