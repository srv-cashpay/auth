package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/auth/configs"
	h_auth "github.com/srv-cashpay/auth/handlers/auth"
	r_auth "github.com/srv-cashpay/auth/repositories/auth"
	s_auth "github.com/srv-cashpay/auth/services/auth"
	"github.com/srv-cashpay/middlewares/middlewares"
)

var (
	DB  = configs.InitDB()
	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)
)

func New() *echo.Echo {

	e := echo.New()

	auth := e.Group("api/auth", middlewares.ApiKeyMiddleware)
	{
		auth.POST("/signup", authH.Signup)
		auth.POST("/signin", authH.Signin)
	}

	logout := e.Group("api/logout")
	{
		logout.POST("", authH.Signout)

	}
	return e
}
