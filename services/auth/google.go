package auth

import (
	"context"
	"errors"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
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

	// Encrypt email untuk disimpan
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}

	encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
	if err != nil {
		return nil, err
	}

	// Cari user by email terenkripsi
	user, err := s.Repo.FindByEmail(encryptedWhatsapp)
	if err != nil {
		// Belum ada → buat user baru
		secureID, err := generateSecureID()
		if err != nil {
			return nil, errors.New("failed to generate secure ID")
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
			if err.Error() == "duplicate_entry" {
				return nil, res.ErrorResponse(&res.ErrorConstant.Duplicate)
			}
			return nil, err
		}

	} else {
		// Jika user sudah ada, update WhatsApp jika kosong
		if user.Whatsapp == "" {
			user.Whatsapp = req.Whatsapp
			if err := s.Repo.UpdateWhatsapp(user.ID, req.Whatsapp); err != nil {
				return nil, err
			}
		}

	}

	// Simulasi token, ganti dengan JWT di produksi
	token, err := s.jwt.GenerateToken(user.ID, user.FullName, user.Merchant.ID)
	refreshtoken, err := s.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)

	return &dto.AuthResponse{
		ID:            user.ID,
		MerchantID:    user.Merchant.ID,
		FullName:      user.FullName,
		Whatsapp:      user.Whatsapp,
		Email:         user.Email,
		Token:         token,
		RefreshToken:  refreshtoken,
		TokenVerified: user.Verified.Token,
	}, nil
}
