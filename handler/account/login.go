package account

import (
	"Api-go/lib"
	"Api-go/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)

	var curUser, validated = ValidatePassword(user.Username, user.Password)
	if validated {
		c.JSON(200, gin.H{
			"status":   200,
			"token": lib.CreateToken(curUser.Username),
			"message":  "login success",
		})
	} else {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "password wrong",
		})
	}
}