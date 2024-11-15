package merchant

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/merchant"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Update(c echo.Context) error {
	var req dto.UpdateMerchantRequest
	updatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	idUint, err := res.QueryParam(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	req.ID = idUint
	req.UpdatedBy = updatedBy

	err = c.Bind(&req)
	if err != nil {
		return res.Response(c, http.StatusBadRequest, res.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	result, err := b.serviceMerchant.Update(req)
	if err != nil {
		return res.Response(c, http.StatusBadRequest, res.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return res.SuccessResponse(result).Send(c)

}
