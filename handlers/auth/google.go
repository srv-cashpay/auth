package handlers

import (
	"net/http"

	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) GoogleSignIn(c echo.Context) error {
	var req dto.GoogleSignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.serviceAuth.SignInWithGoogle(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return res.SuccessResponse(resp).Send(c)

}
