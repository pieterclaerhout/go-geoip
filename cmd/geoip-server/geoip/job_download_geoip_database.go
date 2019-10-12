package geoip

import (
	"github.com/pieterclaerhout/go-geoip"
	"github.com/pieterclaerhout/go-log"
)

// DownloadGeoIPDatabaseJob is a background job for updating the GeoIP database
type DownloadGeoIPDatabaseJob struct {
	GeoDBDownloader *geoip.DatabaseDownloader
}

// Name returns the name of the job
func (job DownloadGeoIPDatabaseJob) Name() string {
	return "download-geoip-database"
}

// Run runs the job and logs errors
func (job DownloadGeoIPDatabaseJob) Run() {
	if err := job.downloadDBIfNeeded(); err != nil {
		log.Error(err)
	}
}

// downloadDBIfUpdated if the database isn't found locally or out-of-date
func (job DownloadGeoIPDatabaseJob) downloadDBIfNeeded() error {

	log.Debug(job.Name(), "| Checking if the database needs updating")

	localChecksum, err := job.GeoDBDownloader.LocalChecksum()
	if err != nil {
		return err
	}

	remoteChecksum, err := job.GeoDBDownloader.RemoteChecksum()
	if err != nil {
		return err
	}

	log.Debug(job.Name(), "| Local checksum: ", localChecksum)
	log.Debug(job.Name(), "| Remote checksum:", remoteChecksum)

	shouldDownload, err := job.GeoDBDownloader.ShouldDownload()
	if err != nil {
		return err
	}

	if !shouldDownload {
		log.Debug(job.Name(), "| Database is up-to-date, no download needed")
		return nil
	}

	log.Warn(job.Name(), "| Database not found or outdated, downloading")

	err = job.GeoDBDownloader.Download()
	if err != nil {
		return err
	}

	log.Info(job.Name(), "| Database downloaded succesfully")

	return nil

}
