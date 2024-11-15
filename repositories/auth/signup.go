package auth

import (
	"errors"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
	"gorm.io/gorm"
)

func (r *authRepository) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {
	var existingUser entity.AccessDoor

	if err := r.DB.Where("email = ?", req.Email).Or("whatsapp = ?", req.Whatsapp).First(&existingUser).Error; err == nil {
		// Jika tidak ada error dan pengguna ditemukan, berarti data sudah ada
		return dto.SignupResponse{}, &res.ErrorConstant.Duplicate
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Jika error lain selain ErrRecordNotFound, kembalikan error tersebut
		return dto.SignupResponse{}, err
	}
	user := entity.AccessDoor{
		ID:       req.ID,
		FullName: req.FullName,
		Whatsapp: req.Whatsapp,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	merchant := entity.MerchantDetail{
		ID:     util.GenerateRandomString(),
		UserID: user.ID,
	}

	if err := r.DB.Save(&merchant).First(&merchant).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     req.Token,
		Otp:       req.Otp,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := r.DB.Save(&verified).First(&verified).Error; err != nil {
		return dto.SignupResponse{}, err
	}
	response := dto.SignupResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Whatsapp: user.Whatsapp,
		Email:    user.Email,
		Password: user.Password,
		Token:    verified.Token,
	}

	return response, nil
}
