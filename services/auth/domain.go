package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"

	r "github.com/srv-cashpay/auth/repositories/auth"
)

type AuthService interface {
	Signup(req dto.SignupRequest) (dto.SignupRequest, error)
}

type authService struct {
	Repo r.DomainRepository
}

func NewAuthService(Repo r.DomainRepository) AuthService {
	return &authService{
		Repo: Repo,
	}
}
