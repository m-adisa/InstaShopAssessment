package models

import "gorm.io/gorm"

type User struct {
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Email    string  `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string  `gorm:"type:varchar(100);not null" json:"-"`
	Order    []Order `json:"orders"`
	gorm.Model
}
