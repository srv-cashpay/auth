package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/auth/configs"
	h_auth "github.com/srv-cashpay/auth/handlers/auth"
	r_auth "github.com/srv-cashpay/auth/repositories/auth"
	s_auth "github.com/srv-cashpay/auth/services/auth"
	"github.com/srv-cashpay/middlewares/middlewares"

	h_verify "github.com/srv-cashpay/auth/handlers/auth/verify_account"
	r_verify "github.com/srv-cashpay/auth/repositories/auth/verify_account"
	s_verify "github.com/srv-cashpay/auth/services/auth/verify_account"

	h_merchant "github.com/srv-cashpay/auth/handlers/merchant"
	r_merchant "github.com/srv-cashpay/auth/repositories/merchant"
	s_merchant "github.com/srv-cashpay/auth/services/merchant"
)

var (
	DB  = configs.InitDB()
	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)

	merchantR = r_merchant.NewMerchantRepository(DB)
	merchantS = s_merchant.NewMerchantService(merchantR, JWT)
	merchantH = h_merchant.NewMerchantHandler(merchantS)

	verifyR = r_verify.NewVerifyRepository(DB)
	verifyS = s_verify.NewVerifyService(verifyR, JWT)
	verifyH = h_verify.NewVerifyHandler(verifyS)
)

func New() *echo.Echo {

	e := echo.New()
	e.POST("/verify", verifyH.HandleVerification)
	e.PUT("/resend-otp", verifyH.ResendVerification)
	// e.POST("/authenticator-admin", verifyH.AuthenticatorAdmin)

	auth := e.Group("api/auth", middlewares.ApiKeyMiddleware)
	{
		auth.POST("/signup", authH.Signup)
		auth.POST("/signin", authH.Signin)
		auth.POST("/authenticator", authH.Authenticator)
	}

	merchant := e.Group("api/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merchant.PUT("/update", merchantH.Update)
		merchant.GET("/get", merchantH.Get)
	}

	logout := e.Group("api/logout")
	{
		logout.POST("", authH.Signout)
	}
	return e
}
