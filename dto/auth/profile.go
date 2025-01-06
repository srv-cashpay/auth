package auth

type ProfileRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type ProfileResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Whatsapp string `json:"whatsapp"`
}
