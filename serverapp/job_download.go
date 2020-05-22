package serverapp

import (
	"time"

	"github.com/pieterclaerhout/go-geoip"
	"github.com/pieterclaerhout/go-log"
)

type dbUpdaterJob struct {
	downloader *geoip.DatabaseDownloader
	interval   time.Duration
}

func (j dbUpdaterJob) run() {
	log.Info("Starting automatic database update, interval:", j.interval)
	for {
		if err := j.downloadDBIfNeeded(); err != nil {
			log.Error(err)
		}
		time.Sleep(j.interval)
	}
}

func (j dbUpdaterJob) downloadDBIfNeeded() error {

	log.Debug("Checking if the database needs updating")

	localChecksum, err := j.downloader.LocalChecksum()
	if err != nil {
		return err
	}

	remoteChecksum, err := j.downloader.RemoteChecksum()
	if err != nil {
		return err
	}

	log.Debug("Local checksum: ", localChecksum)
	log.Debug("Remote checksum:", remoteChecksum)

	shouldDownload, err := j.downloader.ShouldDownload()
	if err != nil {
		return err
	}

	if !shouldDownload {
		log.Debug("Database is up-to-date, no download needed")
		return nil
	}

	log.Warn("Database not found or outdated, downloading")

	if err := j.downloader.Download(); err != nil {
		return err
	}

	log.Info("Database downloaded succesfully")

	return nil

}
