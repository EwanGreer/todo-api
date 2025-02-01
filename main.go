package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]any{"msg": "hello world"})
	})

	if err := e.Start("0.0.0.0:8000"); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
