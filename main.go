package main

import (
	"Api-go/lib"
	"Api-go/model"
	"Api-go/router"
)

func main() {
	lib.InitDb()

	db := lib.DBConn()
	_ = db.AutoMigrate(&model.User{})
	//go socket.SendToAll(socket.ChatChan)
	router.InitRouter()

}
