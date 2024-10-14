package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
)

func (h *domainHandler) Signup(c echo.Context) error {
	var req dto.SignupRequest
	var resp dto.SignupResponse

	// if err := c.Bind(&req); err != nil {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	// }

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	// if err := h.serviceAuth.Signup(&req); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	// }

	resp, err = h.serviceAuth.Signup(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}
