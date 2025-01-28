package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
)

func (u *domainHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	Id, ok := c.Get("Id").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = Id

	// Validate the refresh token (validate inside the service)
	accessToken, err := u.serviceAuth.RefreshAccessToken(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	// Prepare the response with the new access token
	resp := dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	// Return success with the new access token
	return res.SuccessResponse(resp).Send(c)
}
