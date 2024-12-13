package jwt

import (
	"e-learning-platform/package/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IdentifyTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		msg := ""
		if token == "" {
			msg = "未登录"
		} else {
			_, role, err := jwt.ParseToken(token)
			if err != nil {
				msg = err.Error()
			} else if role != "teacher" {
				msg = "没有操作权限"
			}
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
		}
		c.Next()
	}
}
