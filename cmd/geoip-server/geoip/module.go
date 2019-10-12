package geoip

import (
	"os"
	"time"

	"github.com/labstack/echo"

	"github.com/pieterclaerhout/go-geoip"
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/go-webserver/jobqueue"
)

// GeoIP defines the geoip module
type GeoIP struct {
	GeoDB *geoip.Database
}

// Register the endpoints on the router
func (module *GeoIP) Register(router *echo.Echo) {
	g := router.Group("/lookup")
	g.Any("", module.handlerLookup)
}

// Start is executed when the server starts
func (module *GeoIP) Start() {

	dbPath := os.Getenv("GEOIP_DB")
	if dbPath == "" {
		log.Fatal("GEOIP_DB env var not set")
	}

	module.GeoDB = geoip.NewDatabase(dbPath)
	log.Info("Using GeoIP db:", dbPath)

	job := &DownloadGeoIPDatabaseJob{
		GeoDBDownloader: geoip.NewDatabaseDownloader(dbPath, 5*time.Second),
	}

	interval := 1 * time.Hour
	if os.Getenv("DEBUG") == "1" {
		interval = 1 * time.Minute
	}

	jobqueue.Default().Every(interval, job)

	job.Run()

}

// Stop is executed when the server stop
func (module *GeoIP) Stop() {
}
