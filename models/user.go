package models

import "gorm.io/gorm"

type User struct {
	Name     string  `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Email    string  `gorm:"type:varchar(100);unique;not null" json:"email" validate:"required"`
	Password string  `gorm:"type:varchar(100);not null" json:"-" validate:"required"`
	UserType string  `gorm:"type:varchar(100);default:'regular'" json:"user_type" validate:"oneof=regular admin"`
	Order    []Order `json:"orders"`
	gorm.Model
}
