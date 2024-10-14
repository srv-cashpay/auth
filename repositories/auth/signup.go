package auth

import (
	"errors"

	"github.com/srv-cashpay/auth/entity"
)

func (r *authRepository) Signup(user *entity.User) error {
	if err := r.DB.Save(user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}
