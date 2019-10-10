package core

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo"
)

func (module *Core) handlerStatus(c echo.Context) error {

	type response struct {
		Hostname string `json:"hostname"`
		Runtime  string `json:"runtime"`
	}

	res := response{
		Hostname: c.Request().Host,
		Runtime:  runtime.Version() + " on " + runtime.GOOS + "/" + runtime.GOARCH,
	}

	return c.JSON(http.StatusOK, res)

}
