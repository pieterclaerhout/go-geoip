package main

import (
	"os"
	"runtime"

	"github.com/pieterclaerhout/go-log"

	"github.com/pieterclaerhout/go-geoip/cmd/geoip-server/server"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.PrintTimestamp = true

	engine := server.New()
	err := engine.Start()
	log.CheckError(err)

	os.Exit(0)

}
