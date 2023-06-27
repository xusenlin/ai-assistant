package models

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-admin/global"
	"go-admin/helper"
	"gorm.io/gorm"
	"time"
)

const (
	PasswordSalt        = "go-admin"     //密码加盐
	TokenExpireDuration = time.Hour * 72 //Token过期时间
	JwtClaimsKey        = "JwtClaims"    //Claims 储存在 *gin.Context 里的关键字，通过JWTAuth中间件设置。
	SuperAdministrator  = "Admin"
)

var TokenSecret = []byte("MareWood")

type Claims struct {
	ID       uint
	Username string
	Status   int
	IsAdmin  bool
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Username               string `gorm:"type:varchar(45);unique;comment:用户名" binding:"required,min=2,max=18"`
	Password               string `gorm:"type:varchar(200)" binding:"required,min=6,max=16"`
	Status                 int    `gorm:"default:0"`
	Avatar                 string `gorm:"type:varchar(255)"`
	IsAdmin                bool   `gorm:"type:tinyint(1)"`
	TokenConsumed          int    `gorm:"default:0"`
	RemainingDialogueCount int    `gorm:"default:0"`
	Desc                   string `gorm:"type:varchar(500)"`
}

type UserPaginate struct {
	PageInfo
	List []User
}

func (u *User) Register() (err error) {

	if global.DB.Where("username = ?", u.Username).First(&User{}).Error != gorm.ErrRecordNotFound {
		return errors.New("username already exists")
	}
	global.DB.Error = nil

	if u.Username == SuperAdministrator {
		u.Status = StatusEnabled
		u.IsAdmin = true
	} else {
		u.Status = StatusDisabled
	}
	u.Password = helper.DigestString(PasswordSalt + u.Password)

	return global.DB.Save(&u).Error
}

func (u *User) Login() error {

	password := helper.DigestString(PasswordSalt + u.Password)

	err := global.DB.Where("username = ? AND password = ?", u.Username, password).First(&u).Error
	if err != nil {
		return errors.New("用户不存在")
	}
	if u.Status != StatusEnabled {
		return errors.New("用户已经被禁用")
	}

	u.Password = "***"

	return nil
}

func (u *User) Destroy(id string) (err error) {
	return global.DB.Where("id = ?", id).Delete(&u).Error
}

func (u *User) UpdateField(id string, field string, fieldContent any) (err error) {
	return global.DB.Model(&u).Where("id = ?", id).Update(field, fieldContent).Error
}
