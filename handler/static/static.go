package static

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/ui"
)

func HandleStyle(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/css")
	return c.String(http.StatusOK, ui.CssString)
}
