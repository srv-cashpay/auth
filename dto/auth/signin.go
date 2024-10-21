package auth

type SigninRequest struct {
	Whatsapp string `json:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type SigninResponse struct {
	ID            string `json:"id"`
	FullName      string `json:"full_name"`
	ProfileID     string `json:"profile_id"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
}

type AuthUnverifiedResponse struct {
	Whatsapp      string `json:"whatsapp"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}
