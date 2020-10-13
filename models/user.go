package models

import (
	"github.com/rubhiauliatirta/rublog/config"
)

var Db = config.Db

func GetUser(email string) (user User, err error) {
	err = Db.Where("email = ?", email).First(&user).Error

	return user, err
}

func CreateUser(user *User) error {
	result := Db.Create(&user)
	return result.Error
}
