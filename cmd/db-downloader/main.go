package main

import (
	"time"

	"github.com/pieterclaerhout/go-geoip"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	log.PrintTimestamp = true

	downloader := geoip.NewDatabaseDownloader("database.mmdb", 5*time.Second)

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
