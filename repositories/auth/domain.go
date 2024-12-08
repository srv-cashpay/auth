package auth

import (
	"sync"

	dto "github.com/srv-cashpay/auth/dto/auth"

	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
	Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
	Signin(req dto.SigninRequest) (*entity.AccessDoor, error)
	UpdateTokenVerified(userID string, otp string, token string) (dto.SigninResponse, error)
	UpdateUser(user *entity.AccessDoor) error
	SigninByPhoneNumber(req dto.SigninRequest) (*entity.AccessDoor, error)
	RefreshToken(req dto.RefreshTokenRequest) (*entity.AccessDoor, error)
	SaveUser(user *entity.AccessDoor) error
}

type authRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.AccessDoor
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
