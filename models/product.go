package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	Orders      []Order `gorm:"many2many:order_products" json:"orders"`
}
