package jwt

import (
	"e-learning-platform/config"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	jwtSecret = []byte(config.Config.JWT.Secret)
)

func GetTokenString(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  role,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Id:        "justAId",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "e-learning-platform",
		NotBefore: time.Now().Unix(),
		Subject:   username,
	})
	return token.SignedString(jwtSecret)
}

// ParseToken 解析Token，并返回用户名以及用户身份
func ParseToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Subject, claims.Audience, nil
	} else {
		return "", "", err
	}
}
