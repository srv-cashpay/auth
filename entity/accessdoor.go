package entity

import (
	"time"

	"gorm.io/gorm"
)

type AccessDoor struct {
	ID                  string         `gorm:"primary_key" json:"id"`
	FullName            string         `gorm:"full_name" json:"full_name"`
	Whatsapp            string         `gorm:"uniqueIndex;type:varchar(200)" json:"whatsapp"`
	Email               string         `gorm:"uniqueIndex;type:varchar(60)" json:"email"`
	Password            string         `gorm:"password" json:"password"`
	Status              bool           `gorm:"status" json:"status"`
	Merchant            MerchantDetail `json:"merchant" gorm:"foreignKey:UserID"`
	ProfilePicture      ProfilePicture `json:"profile_picture" gorm:"foreignKey:UserID"`
	Verified            UserVerified   `json:"verified" gorm:"foreignKey:UserID"`
	File                []File         `json:"file" gorm:"foreignKey:UserID"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FailedLoginAttempts int            `json:"failed_login_attempts" gorm:"default:0"`
	SuspendedUntil      *time.Time     `json:"suspended_until" gorm:"default:null"`
}
