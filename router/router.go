package router

import (
	"Api-go/handler/account"
	"Api-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	var user = r.Group("account")
	{
		user.POST("/register", account.Register)
		user.POST("/login", account.Login)
	}

	var auth = r.Group("api")
	auth.Use(middleware.AuthMiddleware())
	{

	}

	_ = r.Run(":5202")
}

