package auth

import (
	"context"
	"errors"
	"fmt"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
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

	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		user = &entity.AccessDoor{
			ID:       email, // bisa diganti UUID
			Email:    email,
			FullName: name,
			Provider: "google",
		}
		_ = s.Repo.Create(user)
	}

	// Simulasi pembuatan token (gunakan JWT sungguhan di produksi)
	token := fmt.Sprintf("simulated-jwt-for-%s", user.Email)

	return &dto.AuthResponse{Token: token}, nil
}
