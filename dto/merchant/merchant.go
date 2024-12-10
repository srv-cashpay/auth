package entity

import (
	"time"

	"gorm.io/gorm"
)

type GetByIdRequest struct {
	ID string `query:"id" validate:"required, id"`
}

type GetMerchantRequest struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	MerchantName string         `json:"merchant_name"`
	Address      string         `json:"address"`
	City         string         `json:"city"`
	Zip          int            `json:"zip"`
	Phone        string         `json:"phone"`
	Description  string         `json:"description"`
	UpdatedBy    string         `json:"update_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type GetMerchantResponse struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	MerchantName string         `json:"merchant_name"`
	Address      string         `json:"address"`
	City         string         `json:"city"`
	Zip          int            `json:"zip"`
	Phone        string         `json:"phone"`
	Description  string         `json:"description"`
	UpdatedBy    string         `json:"update_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type UpdateMerchantRequest struct {
	ID           string    `query:"id" validate:"required, id"`
	UserID       string    `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Zip          int       `json:"zip"`
	Phone        string    `json:"phone"`
	Description  string    `json:"description"`
	UpdatedBy    string    `json:"update_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateMerchantResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Zip          int       `json:"zip"`
	Phone        string    `json:"phone"`
	Description  string    `json:"description"`
	UpdatedBy    string    `json:"update_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}
