package auth

import (
	"sync"

	dto "github.com/cashpay/cashpay-auth-srv/dto/auth"

	"github.com/cashpay/cashpay-auth-srv/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
}

type authRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
