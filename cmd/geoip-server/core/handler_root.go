package core

import (
	"net/http"

	"github.com/labstack/echo"
)

func (module *Core) handlerRoot(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotFound, "Move on, nothing here!")
}
