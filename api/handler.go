package api

import (
	"github.com/labstack/echo/v4"
)

type ApiHandler struct{}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

func (h ApiHandler) HandleGetUserById(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"msg": map[string]any{
			"username": "user1",
			"email":    "dummy@emails.com",
		},
	})
}
