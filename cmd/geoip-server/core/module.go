package core

import (
	"os"

	"github.com/labstack/echo"
	"github.com/pieterclaerhout/go-geoip"
	"github.com/pieterclaerhout/go-log"
)

// Core defines the core module
type Core struct {
	GeoipDB *geoip.Database
}

// Register the endpoints on the router
func (module *Core) Register(router *echo.Echo) {
	g := router.Group("/")
	g.GET("", module.handlerRoot)
	g.Any("lookup", module.handlerLookup)
	g.Any("status", module.handlerStatus)
}

// Start is executed when the server starts
func (module *Core) Start() {

	dbPath := os.Getenv("GEOIP_DB")
	if dbPath == "" {
		log.Fatal("GEOIP_DB env var not set")
	}

	module.GeoipDB = geoip.NewDatabase(dbPath)
	log.Info("Using GeoIP db:", dbPath)

}

// Stop is executed when the server stop
func (module *Core) Stop() {
}
