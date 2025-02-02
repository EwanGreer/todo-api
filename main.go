package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/EwanGreer/todo/api"
	"github.com/EwanGreer/todo/pages"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var env *string

func init() {
	env = flag.String("env", "local", "Environment file to read")
	flag.Parse()
}

func main() {
	err := godotenv.Load(".env." + *env)
	if err != nil {
		slog.Info("No .env file found")
	}

	l := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With("service", "todo-api")
	slog.SetDefault(l)

	srv := echo.New()
	srv.HideBanner = true

	srv.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
	)

	pagesHandler := pages.NewPagesHandler()
	apiHandler := api.NewApiHandler()

	api := srv.Group("/api/v1")

	srv.GET("/", pagesHandler.HandleRootPage)
	api.GET("/user/:id", apiHandler.HandleGetUserById)

	if err := srv.Start("0.0.0.0:" + os.Getenv("PORT")); err != nil {
		slog.Error("Start", "error", err)
		os.Exit(1)
	}
}
