package main

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/go-webserver"

	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/core"
	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/geoip"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = true

	server := webserver.New()
	server.Register(&core.Core{})
	server.Register(&geoip.GeoIP{})

	err := server.Start()
	log.CheckError(err)

}
