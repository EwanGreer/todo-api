package pages

import (
	"embed"
	"log/slog"

	"github.com/labstack/echo/v4"
)

const pagesDir = "public/html"

//go:embed public
var publicFS embed.FS

type PagesHandler struct {
	// TODO: database sqlite3
}

func NewPagesHandler() *PagesHandler {
	return &PagesHandler{}
}

func (h PagesHandler) HandleRootPage(c echo.Context) error {
	b, err := publicFS.ReadFile(pagesDir + "/index.html")
	if err != nil {
		slog.Error("ReadFile", "error", err)
		return c.HTML(500, "<p>Something went wrong</p>")
	}

	return c.HTMLBlob(200, b)
}
