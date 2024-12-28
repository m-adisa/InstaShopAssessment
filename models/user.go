package models

import "gorm.io/gorm"

type User struct {
	Name       string  `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Email      string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password   string  `gorm:"type:varchar(100);not null" json:"password" validate:"required"`
	Role       string  `gorm:"type:varchar(100);default:'regular'" json:"role" validate:"required,oneof=regular admin"`
	Order      []Order `json:"orders" swaggerignore:"true"`
	gorm.Model `swaggerignore:"true"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
