package main

import (
	"embed"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const pagesDir = "public/html"

//go:embed public
var publicFS embed.FS

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With("service", "todo-api")
	slog.SetDefault(l)

	e := echo.New()
	e.HideBanner = true

	e.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		b, err := publicFS.ReadFile(pagesDir + "/index.html")
		if err != nil {
			slog.Error("ReadFile", "error", err)
			return c.HTML(500, "<p>Something went wrong</p>")
		}

		return c.HTMLBlob(200, b)
	})

	// TODO: auth (JWT)
	e.GET("/api/user/:id", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"msg": map[string]any{
				"username": "user1",
				"email":    "dummy@emails.com",
			},
		})
	})

	if err := e.Start("0.0.0.0:8000"); err != nil {
		slog.Error("Start", "error", err)
		os.Exit(1)
	}
}
