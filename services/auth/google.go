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

	secureID, err := generateSecureID()
	if err != nil {
		return nil, errors.New("id null")
	}
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		user = &entity.AccessDoor{
			ID:       secureID, // bisa diganti UUID
			Email:    encryptedEmail,
			FullName: name,
			Provider: "google",
		}
		_ = s.Repo.Create(user)
	}

	// Simulasi pembuatan token (gunakan JWT sungguhan di produksi)
	return &dto.AuthResponse{Token: email}, nil

}
