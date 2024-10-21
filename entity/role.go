package entity

type Role struct {
	ID   string `gorm:"primary_key,omitempty" json:"id"`
	Role string `gorm:"type:varchar(36)" json:"user_id"`
}
