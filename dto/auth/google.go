package auth

type GoogleSignInRequest struct {
	IdToken  string `json:"idToken" validate:"required"`
	Whatsapp string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
}

type AuthResponse struct {
	ID            string `json:"id"`
	MerchantID    string `json:"merchant_id"`
	Whatsapp      string `json:"whatsapp"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	RefreshToken  string `json:"refresh_token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
	Status        bool `json:"status"`
}
