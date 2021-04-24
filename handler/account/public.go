package account

import (
	"Api-go/lib"
	"Api-go/model"
)

func ValidatePassword(name string, password string) (model.User, bool) {
	var curUser model.User
	db := lib.DBConn()
	var err = db.Where("username = ?", name).First(&curUser).Error
	if err != nil || curUser.Password != password {
		return curUser, false
	}
	return curUser, true
}

