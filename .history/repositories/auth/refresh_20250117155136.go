package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (u *authRepository) RefreshToken(req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	var existingUser entity.AccessDoor

	request := dto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
		UserID:       existingUser.ID,
	}

	err := u.DB.Where("id = ?", request.UserID).Preload("Verified").Preload("Merchant").First(&existingUser).Error
	if err != nil {
		return dto.RefreshTokenResponse{}, err
	}

	response := dto.RefreshTokenResponse{}

	return response, nil
}
