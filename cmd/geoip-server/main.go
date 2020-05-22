package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pieterclaerhout/go-geoip/serverapp"
	"github.com/pieterclaerhout/go-log"
	webserver "github.com/pieterclaerhout/go-webserver/v2"
)

func main() {

	// Setup logging
	log.PrintColors = true
	log.PrintTimestamp = true
	log.DebugMode = (os.Getenv("DEBUG") == "1")

	// Load the .env file
	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, ".env")
	godotenv.Load(envPath)

	// Run the app with the server
	err := webserver.New().RunWithApps(
		serverapp.New(),
	)
	log.CheckError(err)

}
