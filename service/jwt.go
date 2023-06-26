package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-admin/global"
	"go-admin/models"
	"time"
)

// GenToken 生成JWT
func JwtGenToken(claims *models.Claims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(models.TokenExpireDuration)), // 过期时间
		Issuer:    "senLin",                                                       // 签发人
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(models.TokenSecret)
}

func JwtParseToken(tokenString string) (*models.Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return models.TokenSecret, nil
	})
	if err != nil {
		return nil, errors.New("Token 无效或已经过期，请重新登录")
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("Token 无效，请重新登录")
}

func JwtGetClaimsByContext(c *gin.Context) (*models.Claims, error) {

	jwtClaims, hasClaims := c.Get(models.JwtClaimsKey)
	if !hasClaims {
		return &models.Claims{}, errors.New("用户信息已丢失")
	}

	claims, ok := jwtClaims.(*models.Claims)
	if !ok {
		return &models.Claims{}, errors.New("用户信息断言失败")
	}
	return claims, nil
}

func JwtGetUserByContext(c *gin.Context) (*models.User, error) {
	claims, err := JwtGetClaimsByContext(c)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = global.DB.Where("id = ?", claims.ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
