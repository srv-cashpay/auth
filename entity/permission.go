package entity

type Permission struct {
	ID         string `gorm:"primary_key,omitempty" json:"id"`
	Permission string `gorm:"permission" json:"permission"`
}
