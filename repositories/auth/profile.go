package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (u *authRepository) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	var existingUser entity.AccessDoor

	if err := u.DB.Where("id = ?", req.UserID).Find(&existingUser).Error; err != nil {
		return dto.ProfileResponse{}, err
	}

	resp := dto.ProfileResponse{
		ID:       existingUser.ID,
		FullName: existingUser.FullName,
		Whatsapp: existingUser.Whatsapp,
	}

	return resp, nil
}
