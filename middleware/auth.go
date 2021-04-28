package middleware

import (
	"Api-go/lib"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			//util.LogPrint("info", "middleware", "前端应携带Authorization header却未携带")
			c.AbortWithStatus(401)
			return
		}

		_, tokenValidated := lib.ParserToken(token)

		if !tokenValidated {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "token error",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthUser(token string) (string, bool) {
	user, tokenValidated := lib.ParserToken(token)
	return user, tokenValidated
}
