package database

import (
	"app/pkg/model"
)

func CreateUser(user model.User) {
	DB.Create(user)
}

func GetUserByPhone(user model.User) model.User {
	var tmpUser model.User
	DB.Find(&tmpUser, "phone = ?", user.Phone)
	return tmpUser
}
