package entity

import (
	"time"
)

type UserVerified struct {
	ID        string    `gorm:"primary_key" json:"id"`
	UserID    string    `gorm:"type:varchar(36);index" json:"user_id"`
	Token     string    `gorm:"token" json:"token"`
	Verified  bool      `gorm:"verified" json:"verified"`
	Otp       string    `gorm:"otp" json:"otp"`
	ExpiredAt time.Time `gorm:"expired_at" json:"expired_at"`
}
