package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/conf"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	authSecret      = []byte("不断奋起，直到羔羊变雄狮！")
	tokenCookieName = "rpa-token"
)

func generateToken(uid string) (token string) {
	hours := conf.C.AuthTokenExpireHour

	c := AuthClaims{
		UID: "uid", // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(hours) * time.Hour).Unix(), // 过期时间
			Issuer:    "rpa-server",                                            // 签发人
		},
	}

	token, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(authSecret)
	return
}

func parseToken(tokenStr string) (uid string, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (i interface{}, err error) { return authSecret, nil })

	if err != nil {
		logrus.Errorf("parse token err: %s, token: %s", err.Error(), tokenStr)
		return
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		logrus.Errorf("assert token failed, token: %s", tokenStr)
		return
	}

	if token.Valid {
		logrus.Errorf("invalid token: %s", tokenStr)
		err = errors.New("invalid token")
		return
	}

	uid = claims.UID
	return
}

type AuthClaims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

func Auth(c *gin.Context) {
	// 获取token值
	tokenStr, err := c.Cookie(tokenCookieName)
	if err != nil {
		c.SetCookie(tokenCookieName, "", 0, "", "", true, true)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 解析token
	uid, err := parseToken(tokenStr)
	if err != nil {
		c.SetCookie(tokenCookieName, "", 0, "", "", true, true)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 判断用户是否存在
	if ok, err := user.Exist(uid); err != nil && !ok {
		c.SetCookie(tokenCookieName, "", 0, "", "", true, true)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
	return
}

func SetToken(c *gin.Context, id string) {
	token := generateToken(id)

	c.SetCookie(tokenCookieName, token, conf.C.AuthTokenExpireHour*3600, "", "", true, true)
	return
}

func ClearToken(c *gin.Context) {
	c.SetCookie(tokenCookieName, "", 0, "", "", true, true)
	return
}
