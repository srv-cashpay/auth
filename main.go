package main

import (
	"log"
	"os"

	"github.com/srv-cashpay/auth/routes"
	"github.com/srv-cashpay/util/s/elasticsearch"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	elasticsearch.Init()
	if elasticsearch.Client == nil {
		log.Println("[ERROR] Elasticsearch client is nil after Init()")
		os.Exit(1)
	}
	log.Println("[INFO] Elasticsearch initialized successfully")

	e := routes.New()

	e.Use(middleware.CORS())

	// Sertifikat Let's Encrypt
	certFile := "/certs/fullchain.pem"
	keyFile := "/certs/privkey.pem"

	// Jalankan HTTPS langsung dari Echo
	err := e.StartTLS(":2345", certFile, keyFile)
	if err != nil {
		log.Fatal("StartTLS error: ", err)
	}
}

// CORSMiddleware ..
func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}

			return next(c)
		}
	}
}
