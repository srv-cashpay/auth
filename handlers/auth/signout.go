package handlers

import (
	"fmt"
	"net/http"
	"time"

	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (u *domainHandler) Signout(c echo.Context) error {
	cookies := http.Cookie{
		Name:     "fullName",
		Value:    "",
		Path:     "/dashboard",
		Expires:  time.Now().Add(-7 * 24 * time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(c.Response().Writer, &cookies)

	err := res.SuccessResponse(nil).Send(c)
	if err != nil {
		// Log the error
		fmt.Println("Error sending response:", err)
	}

	return err
}
