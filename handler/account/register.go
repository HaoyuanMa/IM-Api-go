package account

import (
	"Api-go/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	if model.CreateUser(model.User{
		Username: user.Username,
		Password: user.Password,
		}) {
		c.JSON(200, gin.H{
			"status ": 200,
			"message": "register success",
		})
	} else {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "register failed",
		})
	}
}
