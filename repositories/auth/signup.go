package auth

import (
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
)

func (r *authRepository) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	user := entity.AccessDoor{
		ID:       req.ID,
		Whatsapp: req.Whatsapp,
	}
	if err := r.DB.Save(&user).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     req.TokenVerified,
		Otp:       req.Otp,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := r.DB.Save(&verified).First(&verified).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	response := dto.SignupResponse{
		ID:            user.ID,
		Whatsapp:      user.Whatsapp,
		TokenVerified: verified.Token,
	}

	return response, nil

}
