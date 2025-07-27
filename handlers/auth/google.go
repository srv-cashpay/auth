package handlers

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) GoogleSignIn(c echo.Context) error {
	var req dto.GoogleSignInRequest
	var resp *dto.AuthResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	resp, err = h.serviceAuth.SignInWithGoogle(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(resp).Send(c)

}
