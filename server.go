package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.HideBanner = true

	enableAccessLogging := os.Getenv("ENABLE_ACCESS_LOGGING")
	if enableAccessLogging == "YES" {
		e.Use(middleware.Logger())
	}

	e.Use(middleware.Recover())

	helloString := os.Getenv("HELLO_STRING")
	if helloString == "" {
		helloString = "Hello!"
	}
	e.GET("/*", func(c echo.Context) error {
		return c.HTML(http.StatusOK, helloString)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
