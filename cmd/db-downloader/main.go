package main

import (
	"os"
	"time"

	"github.com/pieterclaerhout/go-geoip/v2"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = true

	licenseKey := os.Getenv("LICENSE_KEY")
	if licenseKey == "" {
		log.Fatal("LICENSE_KEY env var not set")
	}

	downloader := geoip.NewDatabaseDownloader(licenseKey, "data/database.mmdb", 1*time.Minute)

	localChecksum, err := downloader.LocalChecksum()
	log.CheckError(err)
	log.Info("Local checksum:", localChecksum)

	remoteChecksum, err := downloader.RemoteChecksum()
	log.CheckError(err)
	log.Info("Remote checksum:", remoteChecksum)

	shouldDownload, err := downloader.ShouldDownload()
	log.CheckError(err)
	log.Info("Should download:", shouldDownload)

	err = downloader.Download()
	log.CheckError(err)

	localChecksum, err = downloader.LocalChecksum()
	log.CheckError(err)
	log.Info("Local checksum:", localChecksum)

}
