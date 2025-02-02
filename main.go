package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/EwanGreer/todo/api"
	"github.com/EwanGreer/todo/pages"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		slog.Info("No .env file found")
	}

	l := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With("service", "todo-api")
	slog.SetDefault(l)

	server := echo.New()
	server.HideBanner = true

	server.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
	)

	pagesHandler := pages.NewPagesHandler()
	apiHandler := api.NewApiHandler()

	api := server.Group("/api/v1")

	server.GET("/", pagesHandler.HandleRootPage)

	api.GET("/user/:id", apiHandler.HandleGetUserById)

	if err := server.Start(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))); err != nil {
		slog.Error("Start", "error", err)
		os.Exit(1)
	}
}
