package main

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/go-webserver"

	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/core"
)

func main() {

	server := webserver.New()
	server.Register(&core.Core{})

	err := server.Start()
	log.CheckError(err)

}
