package models

import "gorm.io/gorm"

type Order struct {
	UserID    int       `gorm:"not null" json:"user_id" validate:"required"`
	User      User      `json:"user"`
	Products  []Product `gorm:"many2many:order_products" json:"products" validate:"required,min=1"`
	TotalCost float64   `gorm:"not null" json:"total_cost" validate:"required,gt=0"`
	gorm.Model
}
