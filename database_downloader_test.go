package geoip_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-geoip"
)

func Test_DatabaseDownloader_New(t *testing.T) {

	dbPath := "database.mmdb"
	timeout := 1 * time.Second

	actual := geoip.NewDatabaseDownloader(dbPath, timeout)

	assert.NotNil(t, actual, "actual")
	assert.Equal(t, dbPath, actual.TargetFilePath)
	assert.Equal(t, geoip.DefaultDownloadURL, actual.DownloadURL)
	assert.Equal(t, geoip.DefaultChecksumURL, actual.ChecksumURL)

}

func Test_DatabaseDownloader_LocalChecksum_NoDBFile(t *testing.T) {

	downloader := geoip.NewDatabaseDownloader("db.mmdb", 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Empty(t, actual, "actual")

}

func Test_DatabaseDownloader_LocalChecksum_NoChecksumFile(t *testing.T) {

	dbPath := "db.mmdb"
	defer os.Remove(dbPath)

	ioutil.WriteFile(dbPath, []byte(""), 0666)

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Empty(t, actual, "actual")

}

func Test_DatabaseDownloader_LocalChecksum_Valid(t *testing.T) {

	expected := "checksum"

	dbPath := "db.mmdb"
	defer os.Remove(dbPath)
	ioutil.WriteFile(dbPath, []byte(expected), 0666)

	checksumPath := dbPath + geoip.DefaultChecksumExt
	defer os.Remove(checksumPath)
	ioutil.WriteFile(checksumPath, []byte(expected), 0666)

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Equal(t, expected, actual, "actual")

}

func Test_DatabaseDownloader_RemoteChecksum_Valid(t *testing.T) {

	expected := "checksum"

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(expected))
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader("", 1*time.Second)
	downloader.ChecksumURL = s.URL

	actual, err := downloader.RemoteChecksum()
	assert.NoError(t, err, "error")
	assert.Equal(t, expected, actual, "actual")

}

func Test_DatabaseDownloader_RemoteChecksum_InvalidURL(t *testing.T) {

	downloader := geoip.NewDatabaseDownloader("", 5*time.Second)
	downloader.ChecksumURL = "ht&@-tp://:aa"

	actual, err := downloader.RemoteChecksum()
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_DatabaseDownloader_RemoteChecksum_Timeout(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(500 * time.Millisecond)
			w.Write([]byte("checksum"))
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader("", 250*time.Millisecond)
	downloader.ChecksumURL = s.URL

	actual, err := downloader.RemoteChecksum()
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_DatabaseDownloader_RemoteChecksum_ReadBodyError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "1")
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader("", 1*time.Second)
	downloader.ChecksumURL = s.URL

	actual, err := downloader.RemoteChecksum()
	assert.Error(t, err)
	assert.Empty(t, actual)

}

// func testDatabaseDownloaderDBPath(t *testing.T) string {
// 	t.Helper()
// 	return "database.mmdb"
// }

// func testDatabaseDownloaderCleanup(t *testing.T, dbPath string) {
// 	t.Helper()
// 	os.Remove(dbPath)
// }
