package geoip

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DatabaseDownloader_New(t *testing.T) {

	dbPath := "database.mmdb"
	timeout := 1 * time.Second

	actual := NewDatabaseDownloader(dbPath, timeout)

	assert.NotNil(t, actual, "actual")
	assert.Equal(t, dbPath, actual.TargetFilePath)
	assert.Equal(t, dbPath+DefaultChecksumExt, actual.localChecksumPath)
	assert.Equal(t, timeout, actual.httpClient.Timeout)

}
