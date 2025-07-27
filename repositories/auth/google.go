package auth

import (
	"errors"
	"time"

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

	if err := r.DB.Save(&merchant).Error; err != nil {
		return nil
	}

	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     util.GenerateRandomString(),
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := r.DB.Save(&verified).Error; err != nil {
		return nil
	}

	return r.DB.Create(user).Error
}

func (r *authRepository) UpdateWhatsapp(userID string, phone string) error {
	return r.DB.Model(&entity.AccessDoor{}).Where("id = ?", userID).Update("whatsapp", phone).Error
}
