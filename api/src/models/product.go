package models

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model  `json:"Model"`
	ProductId   int             `gorm:"primary_key;auto_increment" json:"productid"`
	ProductName string          `gorm:"size:255;not null" json:"productname"`
	Quantity    int             `gorm:"not null" json:"quantity"`
	Price       decimal.Decimal `gorm:"type:numeric not null" json:"price"`
	CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
