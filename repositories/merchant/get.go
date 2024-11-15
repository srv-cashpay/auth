package merchant

import (
	"github.com/srv-cashpay/auth/entity"
)

func (r *merchantRepository) Get() ([]entity.MerchantDetail, error) {
	var data []entity.MerchantDetail

	if err := r.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
