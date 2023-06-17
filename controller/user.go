package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/models"
	"go-admin/service/serviceUser"
	"net/http"
	"strconv"
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

func UsersFind(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 20
	}

	keyword := c.Query("name")

	db := global.DB.Model(&models.User{}).Order("created_at asc")

	if keyword != "" {
		db = db.Where("username LIKE ?", "%"+keyword+"%")
	}
	var users []models.User
	r, perr := models.Paginate(db, pageNum, pageSize, &users)

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
		"Data": models.UserPaginate{
			PageInfo: *r,
			List:     users,
		},
		"Msg": "success",
	})
}

func UserDestroy(c *gin.Context) {
	id := c.Query("id")

	err := new(models.User).Destroy(id)

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
		"Data":   id,
		"Msg":    "",
	})

}
