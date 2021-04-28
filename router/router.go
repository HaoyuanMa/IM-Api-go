package router

import (
	"Api-go/handler/account"
	"Api-go/handler/socket"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(middleware.CORSMiddleware())

	var user = r.Group("account")
	{
		user.POST("/register", account.Register)
		user.POST("/login", account.Login)
	}

	var auth = r.Group("status")
	//auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/buildConnection", socket.BuildConnection)
		//auth.GET("/setOnline",status.SetOnline)
	}

	_ = r.Run(":5202")
}
