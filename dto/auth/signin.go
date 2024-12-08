package auth

type SigninRequest struct {
	Whatsapp string `json:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type SigninResponse struct {
	ID            string `json:"id"`
	MerchantID    string `json:"merchant_id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	RefreshToken  string `json:"refresh_token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
	Status        bool `json:"status"`
}

type AuthUnverifiedResponse struct {
	Whatsapp      string `json:"whatsapp"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}
