package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name" gorm:"type:varchar(20);not null"`
	Telephone string `json:"telephone" form:"telephone" gorm:"varchar(11);not null;unique"`
	Password  string `json:"password" form:"password" gorm:"size:255;not null"`
}
