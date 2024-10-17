package auth

type SigninRequest struct {
	Whatsapp string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
}

type SigninResponse struct {
	ID            string `gorm:"primary_key" json:"id"`
	Whatsapp      string `gorm:"uniqueIndex;type:varchar(20)" json:"whatsapp"`
	ProfileID     string `json:"profile_id"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
}

type AuthUnverifiedResponse struct {
	Email         string `json:"email"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}
