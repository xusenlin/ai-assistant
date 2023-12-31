package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/helper"
	"go-admin/models"
	"go-admin/service"
	"net/http"
	"strconv"
	"strings"
)

func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.Trans(err).Error(),
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
			"Msg":    global.Trans(err).Error(),
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
	token, err := service.JwtGenToken(&models.Claims{
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

	keyword := c.Query("keyword")

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

func UserUpdateRemainingDialogueCount(c *gin.Context) {
	id := c.Query("id")
	count := c.Query("count")

	err := new(models.User).UpdateField(id, "remaining_dialogue_count", count)

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

func UserUpdateStatus(c *gin.Context) {
	id := c.Query("id")
	status := c.Query("status")

	err := new(models.User).UpdateField(id, "status", status)

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

func UserUpdatePassword(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	if len(password) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    "密码最少6位",
		})
		return
	}

	md5 := helper.DigestString(models.PasswordSalt + password)

	err := new(models.User).UpdateField(id, "password", md5)

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

type ReqUsers struct {
	Username               string `binding:"required,min=2"`
	Password               string `binding:"omitempty,min=6,max=16"`
	Status                 int
	IsAdmin                bool
	RemainingDialogueCount int
}

func UserBatchAdd(c *gin.Context) {
	var req ReqUsers

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status": false,
			"Data":   "",
			"Msg":    global.Trans(err).Error(),
		})
		return
	}

	req.Username = strings.ReplaceAll(req.Username, "，", ",")

	usernames := strings.Split(req.Username, ",")
	var password string
	if req.Password != "" {
		password = helper.DigestString(models.PasswordSalt + req.Password)
	}

	errorInfo := map[string]string{}
	var successInfo []string

	for _, u := range helper.RemoveDuplicate(usernames) {
		name := strings.Trim(u, " ")
		user := models.User{
			Username:               name,
			Password:               password,
			Status:                 req.Status,
			IsAdmin:                req.IsAdmin,
			RemainingDialogueCount: req.RemainingDialogueCount,
		}
		if err := global.DB.Create(&user).Error; err != nil {
			errorInfo[name] = err.Error()
		} else {
			successInfo = append(successInfo, name)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"Status": true,
		"Data": map[string]any{
			"success": successInfo,
			"error":   errorInfo,
		},
		"Msg": "操作成功",
	})
}
