package auth

type SigninRequest struct {
	Whatsapp string `json:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type SigninResponse struct {
	ID            string `json:"id"`
	MerchantID    string `json:"merchant_id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
	Status        bool `json:"status"`
}

type AuthUnverifiedResponse struct {
	Whatsapp      string `json:"whatsapp"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}
