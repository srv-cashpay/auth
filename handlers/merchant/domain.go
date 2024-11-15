package merchant

import (
	s "github.com/srv-cashpay/auth/services/merchant"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceMerchant s.MerchantService
}

func NewMerchantHandler(service s.MerchantService) DomainHandler {
	return &domainHandler{
		serviceMerchant: service,
	}
}
