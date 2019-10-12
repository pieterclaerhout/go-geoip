package core

import (
	"github.com/labstack/echo"

	"github.com/pieterclaerhout/go-geoip"
)

// Core defines the core module
type Core struct {
	GeoDB *geoip.Database
}

// Register the endpoints on the router
func (module *Core) Register(router *echo.Echo) {
	g := router.Group("/")
	g.GET("", module.handlerRoot)
	g.Any("status", module.handlerStatus)
}

// Start is executed when the server starts
func (module *Core) Start() {
}

// Stop is executed when the server stop
func (module *Core) Stop() {
}
