package auth

import (
	"context"
	"errors"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
	"google.golang.org/api/idtoken"
)

func (s *authService) SignInWithGoogle(req dto.GoogleSignInRequest) (*dto.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), req.IdToken, "")
	if err != nil {
		return nil, err
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errors.New("email not found in token")
	}
	name, _ := payload.Claims["name"].(string)

	// Enkripsi hanya sekali
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}

	// Cari user berdasarkan email terenkripsi
	user, err := s.Repo.FindByEncryptedEmail(encryptedEmail)
	if err != nil && err.Error() == "user not found" {
		// Buat user baru
		secureID, err := generateSecureID()
		if err != nil {
			return nil, errors.New("failed to generate secure ID")
		}

		encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
		if err != nil {
			return nil, err
		}

		user = &entity.AccessDoor{
			ID:           secureID,
			Email:        encryptedEmail,
			FullName:     name,
			Provider:     "google",
			AccessRoleID: "e9Wl2JyVeBM_", // default role
			Whatsapp:     encryptedWhatsapp,
		}

		if err := s.Repo.Create(user); err != nil {
			return nil, err
		}

		// Ambil user baru yang sudah disimpan dan preload merchant
		user, err = s.Repo.FindByEncryptedEmail(encryptedEmail)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		// Error lainnya
		return nil, err
	} else {
		// Jika sudah ada dan WhatsApp kosong → update
		if user.Whatsapp == "" && req.Whatsapp != "" {
			if err := s.Repo.UpdateWhatsapp(user.ID, req.Whatsapp); err != nil {
				return nil, err
			}
			user.Whatsapp = req.Whatsapp
		}
	}

	// Generate token
	token, err := s.jwt.GenerateToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		ID:            user.ID,
		MerchantID:    user.Merchant.ID,
		FullName:      user.FullName,
		Whatsapp:      user.Whatsapp,
		Email:         user.Email,
		Token:         token,
		RefreshToken:  refreshToken,
		TokenVerified: user.Verified.Token,
	}, nil
}
