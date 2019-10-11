package core

import (
	"net/http"

	"github.com/labstack/echo"
)

func (module *Core) handlerLookup(c echo.Context) error {

	type request struct {
		IPAddress string `json:"ip" form:"ip" query:"ip"`
	}

	var r request
	if err := c.Bind(&r); err != nil {
		return err
	}

	result, err := module.GeoDB.Lookup(r.IPAddress)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)

}
