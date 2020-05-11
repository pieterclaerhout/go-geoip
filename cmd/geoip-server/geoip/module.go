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
	GeoDB           *geoip.Database
	GeoDBDownloader *geoip.DatabaseDownloader
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

	licenseKey := os.Getenv("LICENSE_KEY")
	if licenseKey == "" {
		log.Fatal("LICENSE_KEY env var not set")
	}

	module.GeoDB = geoip.NewDatabase(dbPath)
	module.GeoDBDownloader = geoip.NewDatabaseDownloader(licenseKey, dbPath, 1*time.Minute)
	log.Info("Using GeoIP db:", dbPath)
	log.Info("Using license key:", licenseKey)

	job := &DownloadGeoIPDatabaseJob{
		GeoDBDownloader: module.GeoDBDownloader,
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
