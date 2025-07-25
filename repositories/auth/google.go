package auth

import (
	"errors"

	"github.com/srv-cashpay/auth/entity"
	"gorm.io/gorm"
)

func (r *authRepository) FindByEmail(email string) (*entity.AccessDoor, error) {
	var user entity.AccessDoor
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) Create(user *entity.AccessDoor) error {
	return r.DB.Create(user).Error
}
