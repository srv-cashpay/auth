package auth

type SignupRequest struct {
	ID            string `gorm:"primary_key" json:"id"`
	Whatsapp      string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
	TokenVerified string `json:"token_verified"`
	Otp           string `json:"token_verified"`
}

type SignupResponse struct {
	ID            string `gorm:"primary_key" json:"id"`
	Whatsapp      string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
}
