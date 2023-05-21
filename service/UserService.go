package service

import (
	"work/config"
	"work/dao"
	"work/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func UserService(user *model.User, ctx *gin.Context) model.User {
	// 创建用户
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	db := dao.NewUser()
	// 判断手机号是否存在
	if isTelephoneExist(user.Telephone) {
		newUser := model.User{}
		return newUser
	}
	newUser := model.User{
		Name:      user.Name,
		Telephone: user.Telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)
	return newUser
}

func isTelephoneExist(telephone string) bool {
	DB := config.GetDB()
	var user model.User
	DB.Where("telephone = ?", telephone).Find(&user)
	return user.ID > 0
}
