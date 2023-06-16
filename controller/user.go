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
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	if err := newUser.Register(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "registration success",
	})

}

func UserLogin(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	token, err := serviceUser.GenToken(&models.Claims{
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	type userAllInfo struct {
		models.User
		Token string
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   userAllInfo{user, token},
		"msg":    "login successful",
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
			"status": false,
			"data":   global.DB.Error.Error(),
			"msg":    "Database query error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
		"msg":    "search successful",
	})
}

func UserDestroy(c *gin.Context) {

	id := c.Query("id")
	var user models.User

	if global.DB.
		Where("id = ? AND username = ?", id, models.SuperAdministrator).
		First(&models.User{}).Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Super administrator cannot be deleted",
		})
		return
	}

	if global.DB.Where("id = ?", id).Delete(&user).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    global.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
	})

}
