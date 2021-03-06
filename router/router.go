package router

import (
	"Api-go/handler/account"
	"Api-go/handler/socket"
	"Api-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Static("/UploadFiles", "C:/Users/mahaoyuan/Desktop/RealTimeWeb/Api-go/UploadFiles/")

	var user = r.Group("Account")
	{
		user.POST("/Register", account.Register)
		user.POST("/Login", account.Login)
	}

	var auth = r.Group("Socket")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/BuildConnection", socket.BuildConnection)
		//auth.GET("/setOnline",status.SetOnline)
	}

	_ = r.Run(":5202")
}
