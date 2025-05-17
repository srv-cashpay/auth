package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("lab.cashpay.my.id"),
		Cache:      autocert.DirCache("certs"),
	}

	go func() {
		log.Fatal(http.ListenAndServe(":80", m.HTTPHandler(nil)))
	}()

	server := &http.Server{
		Addr:      ":443",
		Handler:   e,
		TLSConfig: m.TLSConfig(),
	}

	log.Fatal(server.ListenAndServeTLS("", ""))
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
