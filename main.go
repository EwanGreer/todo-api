package main

import (
	"flag"
	"fmt"
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

	fmt.Println(*env)
}

func main() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", *env))
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

	api := srv.Group("/api/v1")

	srv.GET("/", pagesHandler.HandleRootPage)
	api.GET("/user/:id", apiHandler.HandleGetUserById)

	if err := srv.Start(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))); err != nil {
		slog.Error("Start", "error", err)
		os.Exit(1)
	}
}
