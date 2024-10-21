package auth

type SignupRequest struct {
	ID       string `json:"id"`
	Otp      string `json:"otp" form:"otp" validate:"required,otp"`
	Whatsapp string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
	Country  string `json:"country"` // Add this field for country selection
	Token    string `json:"token"`
}

type SignupResponse struct {
	ID       string `json:"id"`
	Whatsapp string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Country  string `json:"country"` // Add this field for country selection
	Password string `json:"-" form:"password" validate:"required,password"`
	Token    string `json:"token"`
}
