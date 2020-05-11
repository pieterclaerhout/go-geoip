package geoip

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, ".env")
	godotenv.Load(envPath)
}

func Test_DatabaseDownloader_New(t *testing.T) {

	dbPath := "database.mmdb"
	timeout := 1 * time.Second

	actual := NewDatabaseDownloader(os.Getenv("LICENSE_KEY"), dbPath, timeout)

	assert.NotNil(t, actual, "actual")
	assert.Equal(t, dbPath, actual.TargetFilePath)
	assert.Equal(t, dbPath+DefaultChecksumExt, actual.localChecksumPath)
	assert.Equal(t, timeout, actual.httpClient.Timeout)

}
