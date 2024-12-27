package route

import (
	"time"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/golang-jwt/jwt/v5"
)

type UserJWTModel struct {
	Username string
	jwt.RegisteredClaims
}

// jwtUserGenerate 生成jwt
func jwtUserGenerate(user string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserJWTModel{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	t, _ := token.SignedString([]byte(appconfig.JWTSecret))
	return t
}

// jwtUserParse 解析jwt
func jwtUserParse(tokenString string) (string, bool) {
	tokenAfter, e := jwt.ParseWithClaims(tokenString, &UserJWTModel{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(appconfig.JWTSecret), nil
	})
	if e != nil {
		return "", false
	}
	cliamAfter := tokenAfter.Claims.(*UserJWTModel)
	return cliamAfter.Username, true
}
