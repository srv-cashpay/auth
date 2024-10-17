package auth

type SignupRequest struct {
	ID       string `gorm:"primary_key" json:"id"`
	Whatsapp string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
}

type SignupResponse struct {
	ID       string `gorm:"primary_key" json:"id"`
	Whatsapp string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
	Token    string `json:"token"`
}
