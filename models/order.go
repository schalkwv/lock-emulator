package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductId uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
	UserId    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserId"`
}
