package entity

type RoleUser struct {
	ID           string `gorm:"primary_key,omitempty" json:"id"`
	RoleID       string `gorm:"type:varchar(36);index,omitempty" json:"role_id"`
	UserID       string `gorm:"type:varchar(36);index,omitempty" json:"user_id"`
	PermissionID string `gorm:"type:varchar(36);index,omitempty" json:"permission_id"`
}
