package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (r *authRepository) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	user := entity.AccessDoor{
		ID:       req.ID,
		Whatsapp: req.Whatsapp,
	}
	if err := r.DB.Save(user).Error; err != nil {
		return dto.SignupResponse{}, err
	}
	response := dto.SignupResponse{
		ID:       user.ID,
		Whatsapp: user.Whatsapp,
	}

	return response, nil

}
