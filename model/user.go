package model

import (
	"Api-go/lib"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username	string `json:"username" gorm:"unique"`
	Password	string `json:"password"`
}

func CreateUser(user User) bool {
	db := lib.DBConn()
	err := db.Create(&user).Error
	if err != nil {
		return false
	}
	return true
}

