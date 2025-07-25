package auth

import (
	"errors"

	"github.com/srv-cashpay/auth/entity"
	entitymerchant "github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
	"gorm.io/gorm"
)

func (r *authRepository) FindByEmail(email string) (*entity.AccessDoor, error) {
	var user entity.AccessDoor
	encryptedEmail, err := util.Encrypt(email)
	if err != nil {
		return nil, err
	}
	if err := r.DB.Where("email = ?", encryptedEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) Create(user *entity.AccessDoor) error {

	merchant := entitymerchant.MerchantDetail{
		ID:         util.GenerateRandomString(),
		UserID:     user.ID,
		CurrencyID: 1,
	}

	if err := r.DB.Save(&merchant).First(&merchant).Error; err != nil {
		return nil
	}

	return r.DB.Create(user).Error
}
