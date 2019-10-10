package server

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// ErrorResponse defines a generic error response for a web request
type ErrorResponse struct {
	XMLName xml.Name `json:"-"`
	Error   string   `json:"error" xml:"message"`
}

// handleError handles the error response based on the request format
func (server *Server) handleError(err error, c echo.Context) {

	code := http.StatusInternalServerError
	message := err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = fmt.Sprintf("%v", he.Message)
	}

	requestedContentType := c.Request().Header.Get("Content-Type")

	// See: https://stackoverflow.com/a/46205879/118188
	response := ErrorResponse{
		xml.Name{Local: "error"},
		message,
	}

	switch requestedContentType {
	case "application/xml":
		c.XML(code, response)
	case "text/xml":
		c.XML(code, response)
	case "text/plain":
		c.String(code, response.Error)
	default:
		c.JSON(code, response)
	}

}
