package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
)

func (u *domainHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	var resp dto.RefreshTokenResponse
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId

	// Validate the refresh token (validate inside the service)
	_, err := u.serviceAuth.RefreshAccessToken(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	// Prepare the response with the new access token
	response := dto.RefreshTokenResponse{
		AccessToken: resp.AccessToken,
	}

	// Return success with the new access token
	return res.SuccessResponse(response).Send(c)
}
