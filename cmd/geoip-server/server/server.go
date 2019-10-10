package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pieterclaerhout/go-log"

	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/server/core"
)

// Server is an abstraction of a webserver
type Server struct {
	engine *echo.Echo
}

// New returns a new Server instacce
func New() *Server {
	return &Server{}
}

// Start starts the webserver on the indicated port
func (server *Server) Start(port string) error {

	server.engine = echo.New()
	server.engine.HideBanner = true
	server.engine.Debug = log.DebugMode
	server.engine.HTTPErrorHandler = server.handleError

	server.registerMiddlewares()

	server.Register(&core.Core{})

	return server.engine.Start(port)

}

// Register registers the module on the main router
func (server *Server) Register(module Module) {
	module.Register(server.engine)
}

// registerMiddlewares registers the middleware which is going to be used
func (server *Server) registerMiddlewares() {

	server.engine.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	server.engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${status} ${uri}\n",
	}))

	server.engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

}
