package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service/serviceUser"
	"gorm.io/gorm"
	"net/http"
)

func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	if err := newUser.Register(); err != nil {
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
		"Msg":    "registration success",
	})

}

func UserLogin(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}

	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    err.Error(),
		})
		return
	}
	token, err := serviceUser.GenToken(&models.Claims{
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
		IsAdmin:  user.IsAdmin,
	})

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
		"Data": struct {
			models.User
			Token string
		}{user, token},
		"Msg": "login successful",
	})

}

func UserFindAll(c *gin.Context) {

	var users []models.User
	keyword := c.Query("name")

	query := global.DB.Order("created_at asc")

	if keyword != "" {
		query = query.Where("username LIKE ?", "%"+keyword+"%")
	}

	if query.Find(&users).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   global.DB.Error.Error(),
			"Msg":    "Database query error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   users,
		"Msg":    "search successful",
	})
}

func UserDestroy(c *gin.Context) {

	id := c.Query("id")
	var user models.User

	if global.DB.
		Where("id = ? AND username = ?", id, models.SuperAdministrator).
		First(&models.User{}).Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "Super administrator cannot be deleted",
		})
		return
	}

	if global.DB.Where("id = ?", id).Delete(&user).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data":   id,
		"Msg":    "successfully deleted",
	})

}
