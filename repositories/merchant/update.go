package merchant

import (
	dto "github.com/srv-cashpay/auth/dto/merchant"
	"github.com/srv-cashpay/auth/entity"
)

func (b *merchantRepository) Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.MerchantDetail{
		MerchantName: req.MerchantName,
		Description:  req.Description,
		Address:      req.Address,
		City:         req.City,
		Zip:          req.Zip,
		Phone:        req.Phone,
		UpdatedBy:    req.UpdatedBy,
		ID:           tr.ID,
	}

	mer, err := b.GetById(tr)
	if err != nil {
		return dto.UpdateMerchantResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.MerchantDetail{
		MerchantName: request.MerchantName,
		Description:  request.Description,
		Address:      request.Address,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.UpdateMerchantResponse{}, err
	}

	response := dto.UpdateMerchantResponse{
		MerchantName: request.MerchantName,
		Description:  request.Description,
		Address:      request.Address,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
		ID:           mer.ID,
	}

	return response, nil
}
