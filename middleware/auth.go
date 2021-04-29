package middleware

import (
	"Api-go/lib"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			log.Printf("no token")
			c.AbortWithStatus(401)
			return
		}

		_, tokenValidated := lib.ParserToken(token)

		if !tokenValidated {
			log.Printf("token error")
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
