package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (u *authRepository) RefreshToken(req dto.RefreshTokenRequest) (*entity.AccessDoor, error) {
	var existingUser entity.AccessDoor

	request := dto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
		UserID:       req.UserID,
	}

	err := u.DB.Where("id = ?", request.UserID).Preload("Verified").Preload("Merchant").First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
