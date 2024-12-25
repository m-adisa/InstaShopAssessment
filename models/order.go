package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    int       `gorm:"not null" json:"user_id"`
	User      User      `json:"user"`
	Products  []Product `gorm:"many2many:order_products" json:"products"`
	TotalCost float64   `gorm:"not null" json:"total_cost"`
}
