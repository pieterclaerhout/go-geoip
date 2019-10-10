package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/pieterclaerhout/go-log"

	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/server"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.PrintTimestamp = true

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	debug := os.Getenv("DEBUG")
	if debug == "1" {
		log.DebugMode = true
	}

	engine := server.New()
	err := engine.Start(port)
	log.CheckError(err)

	os.Exit(0)

}
