package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/auth/repositories/auth"
)

type AuthService interface {
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
}

type authService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthService(Repo r.DomainRepository, jwtS m.JWTService) AuthService {
	return &authService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
