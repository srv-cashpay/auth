package entity

import (
	"time"

	"gorm.io/gorm"
)

type MerchantDetail struct {
	ID           string         `gorm:"primary_key" json:"id"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantName string         `gorm:"type:varchar(50)" json:"merchant_name"`
	Address      string         `gorm:"address" json:"address"`
	Country      string         `gorm:"country" json:"country"`
	City         string         `gorm:"city" json:"city"`
	Zip          string         `gorm:"zip" json:"zip"`
	Phone        string         `gorm:"phone" json:"phone"`
	CurrencyID   int            `gorm:"currency_id" json:"currency_id"`
	Description  string         `gorm:"description" json:"description"`
	UpdatedBy    string         `gorm:"update_by" json:"update_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
