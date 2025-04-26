package models

import "time"

type Product struct {
	ID          uint    `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	ProductName string  `json:"product_name" gorm:"column:product_name; type:varchar(255)"`
	Total       int     `json:"total" gorm:"column:total; type:int"`
	Price       float32 `json:"price" gorm:"column:price; type:float"`
	UserID      uint    `json:"user_id" gorm:"column:user_id; type:int"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "product"
}
