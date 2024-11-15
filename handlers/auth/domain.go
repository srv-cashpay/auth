package handlers

import (
	s "github.com/srv-cashpay/auth/services/auth"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Signup(c echo.Context) error        //masuk
	Signin(c echo.Context) error        //masuk
	Signout(c echo.Context) error       //keluar
	Authenticator(c echo.Context) error //Authenticator
}

type domainHandler struct {
	serviceAuth s.AuthService
}

func NewAuthHandler(service s.AuthService) DomainHandler {
	return &domainHandler{
		serviceAuth: service,
	}
}
