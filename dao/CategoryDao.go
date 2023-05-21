package dao

import (
	"work/config"
	"work/model"

	"gorm.io/gorm"
)

type CategoryDao interface {
	Create(Category *model.User)
}

type Category struct {
	DB *gorm.DB
}

func (c *Category) Create(category *model.Category) bool {

	result := c.DB.Create(category)
	err := result.Error
	return err != nil
}

func NewCategory() Category {
	db := config.GetDB()
	db.AutoMigrate(&model.Category{})
	return Category{DB: db}
}
