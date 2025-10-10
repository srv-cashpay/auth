package main

import (
	"log"
	"os"

	"github.com/srv-cashpay/auth/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := routes.New()

	e.Use(middleware.CORS())

	// Sertifikat Let's Encrypt
	// Sertifikat Let's Encrypt
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile == "" || keyFile == "" {
		log.Fatal("CERT_FILE atau KEY_FILE tidak ditemukan di environment")
	}

	log.Printf("Starting HTTPS server on :2345 (cert: %s)", certFile)
	if err := e.StartTLS(":2345", certFile, keyFile); err != nil {
		log.Fatal("StartTLS error:", err)
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
