package product

import (
	dto "github.com/srv-cashpay/auth/dto/merchant"
)

func (b *merchantService) Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error) {
	request := dto.UpdateMerchantRequest{
		ID:           req.ID,
		MerchantName: req.MerchantName,
		Description:  req.Description,
		Address:      req.Address,
		City:         req.City,
		Zip:          req.Zip,
		Phone:        req.Phone,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateMerchantResponse{
		ID:           req.ID,
		MerchantName: request.MerchantName,
		Description:  request.Description,
		Address:      request.Address,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}
