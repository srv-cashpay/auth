package main

import (
	"log"
	"net/http"

	"github.com/srv-cashpay/auth/routes"
	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func main() {

// 	e := routes.New()

// 	e.Use(middleware.CORS())

// 	e.Logger.Fatal(e.Start(":2345"))
// }

func main() {

	e := routes.New()
	e.Use(middleware.CORS())

	// Autocert untuk HTTPS otomatis dari Let's Encrypt
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("cashpay.my.id"), // GANTI dengan domain kamu
		Cache:      autocert.DirCache("certs"),              // Folder lokal untuk simpan cert
	}

	// Redirect HTTP ke HTTPS di port 80
	go func() {
		log.Println("Running HTTP redirect on :80")
		http.ListenAndServe(":80", m.HTTPHandler(nil))
	}()

	// Jalankan server HTTPS di port 443
	server := http.Server{
		Addr:      ":443",
		Handler:   e,
		TLSConfig: m.TLSConfig(),
	}

	log.Println("Running HTTPS server on :443")
	log.Fatal(server.ListenAndServeTLS("", ""))
	e.Logger.Fatal(e.Start(":2345"))

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
