package serverapp

import (
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/pieterclaerhout/go-geoip/v2"
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/go-webserver/v2"
)

type serverApp struct {
	db *geoip.Database
}

// New returns a ServerApp instance
func New() webserver.App {

	dbPath := getenv("GEOIP_DB")
	log.Info("Using GeoIP db:", dbPath)

	licenseKey := getenv("LICENSE_KEY")
	log.Info("Using license key:", licenseKey)

	db := geoip.NewDatabase(dbPath)

	interval := 1 * time.Hour
	if os.Getenv("DEBUG") == "1" {
		interval = 1 * time.Minute
	}

	updaterJob := dbUpdaterJob{
		db:         db,
		downloader: geoip.NewDatabaseDownloader(licenseKey, dbPath, 30*time.Second),
		interval:   interval,
	}

	err := updaterJob.downloadDBIfNeeded()
	log.CheckError(err)

	go updaterJob.run()

	return &serverApp{
		db: db,
	}

}

// Name returns the name of this app
func (a *serverApp) Name() string {
	return "geoip"
}

// Register registers the routes for this app
func (a *serverApp) Register(r *chi.Mux) {

	r.Get("/lookup", a.handleLookup())
	r.Post("/lookup", a.handleLookup())

	r.NotFound(a.handleNotFound())
	r.MethodNotAllowed(a.handleMethodNotAllowed())

}
