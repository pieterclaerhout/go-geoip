package serverapp

import (
	"os"

	"github.com/pieterclaerhout/go-log"
)

func getenv(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatal(name + " env var not set")
	}
	return val
}
