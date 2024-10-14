package auth

import (
	dto "github.com/cashpay/cashpay-auth-srv/dto/auth"

	r "github.com/cashpay/cashpay-auth-srv/repositories/auth"
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
