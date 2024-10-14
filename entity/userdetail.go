package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID          string         `gorm:"primary_key" json:"id"`
	UserID      string         `gorm:"uniqueIndex:varchar(36);index" json:"user_id"`
	ProfileID   string         `gorm:"uniqueIndex;type:varchar(255)" json:"profile_id"`
	FullName    string         `gorm:"type:varchar(50)" json:"full_name"`
	DataAccount string         `gorm:"type:varchar(10)" json:"data_account"`
	Age         int            `gorm:"age" json:"age"`
	Gender      string         `gorm:"type:varchar(40)" json:"gender"`
	Description string         `gorm:"description" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
