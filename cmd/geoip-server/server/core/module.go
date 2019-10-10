package core

import (
	"github.com/labstack/echo"
)

// Core defines the core module
type Core struct{}

// Register the endpoints on the router
func (module *Core) Register(router *echo.Echo) {
	g := router.Group("/")
	g.GET("", module.handlerRoot)
	g.Any("status", module.handlerStatus)
}
