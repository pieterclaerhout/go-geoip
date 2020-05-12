package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/core"
	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/geoip"
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/go-webserver"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = true

	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, ".env")
	godotenv.Load(envPath)

	server := webserver.New()
	server.Register(&core.Core{})
	server.Register(&geoip.GeoIP{})

	err := server.Start()
	log.CheckError(err)

}
