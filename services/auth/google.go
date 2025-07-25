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

	// Encrypt email untuk disimpan
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}

	// Cari user by email terenkripsi
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		// Kalau belum ada user → buat user baru
		secureID, err := generateSecureID()
		if err != nil {
			return nil, errors.New("failed to generate secure ID")
		}

		user = &entity.AccessDoor{
			ID:           secureID,
			Email:        encryptedEmail,
			FullName:     name,
			Provider:     "google",
			AccessRoleID: "e9Wl2JyVeBM_", // ganti sesuai role
		}

		if err := s.Repo.Create(user); err != nil {
			return nil, err
		}

	}

	// Simulasi token, ganti dengan JWT di produksi
	token, err := s.jwt.GenerateToken(user.ID, user.FullName, user.Merchant.ID)
	refreshtoken, err := s.jwt.GenerateRefreshToken(user.ID, user.FullName, user.Merchant.ID)

	return &dto.AuthResponse{
		ID:            user.ID,
		MerchantID:    user.Merchant.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Token:         token,
		RefreshToken:  refreshtoken,
		TokenVerified: user.Verified.Token,
	}, nil
}
