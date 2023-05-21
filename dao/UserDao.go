package dao

import (
	"work/config"
	"work/model"

	"gorm.io/gorm"
)

type UserDao interface {
	Create(user *model.User)
}

type User struct {
	DB *gorm.DB
}

func (u *User) Create(user *model.User) {
	result := u.DB.Create(user)
	err := result.Error
	if err != nil {
		panic(err)
	}
}

func NewUser() User {
	db := config.GetDB()
	db.AutoMigrate(&model.User{})
	return User{DB: db}
}
