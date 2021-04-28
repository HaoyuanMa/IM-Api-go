package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ws *websocket.Conn
var err error

func UpgradeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
