package auth

import (
	"sync"

	dto "github.com/srv-cashpay/auth/dto/auth"

	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
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
