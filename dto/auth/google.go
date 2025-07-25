package auth

type GoogleSignInRequest struct {
	IdToken string `json:"idToken" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
