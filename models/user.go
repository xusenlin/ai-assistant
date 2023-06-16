package models

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-admin/global"
	"go-admin/helper"
	"gorm.io/gorm"
	"time"
)

const ( //用户状态
	UserStatusEnabled  = 1
	UserStatusDisabled = 2
)

const (
	PasswordSalt        = "go-admin"     //密码加盐
	TokenExpireDuration = time.Hour * 10 //Token过期时间
	JwtClaimsKey        = "JwtClaims"    //Claims 储存在 *gin.Context 里的关键字，通过JWTAuth中间件设置。
	SuperAdministrator  = "Admin"
)

var TokenSecret = []byte("MareWood")

type Claims struct {
	ID       uint
	Username string
	Status   int
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(45);unique;comment:用户名" binding:"required,min=2,max=18"`
	Password string `json:"password"  gorm:"type:varchar(200)" binding:"required,min=6,max=16"`
	Status   int    `json:"status" gorm:"default:2"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc" gorm:"type:varchar(200)"`
}

func (u *User) Register() (err error) {

	if global.DB.Where("username = ?", u.Username).First(&User{}).Error != gorm.ErrRecordNotFound {
		return errors.New("username already exists")
	}

	if u.Username == "Admin" {
		u.Status = UserStatusEnabled
	} else {
		u.Status = UserStatusDisabled
	}
	u.Password = helper.DigestString(PasswordSalt + u.Password)

	return global.DB.Save(&u).Error
}

func (u *User) Login() error {

	password := helper.DigestString(PasswordSalt + u.Password)

	err := global.DB.Where("username = ? AND password = ?", u.Username, password).First(&u).Error
	if err != nil {
		return err
	}
	if u.Status != UserStatusEnabled {
		return errors.New("user is disabled")
	}

	//u.Password = "***"

	return nil
}
