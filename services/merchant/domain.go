package product

import (
	dto "github.com/srv-cashpay/auth/dto/merchant"
	entity "github.com/srv-cashpay/auth/entity"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/auth/repositories/merchant"
)

type MerchantService interface {
	Get() ([]entity.MerchantDetail, error)
	Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error)
}

type merchantService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMerchantService(Repo r.DomainRepository, jwtS m.JWTService) MerchantService {
	return &merchantService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
