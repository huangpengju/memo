package middleware

import (
	"memo/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		// var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404 // 参数不对，没有token
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 // 无效了
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token 解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
