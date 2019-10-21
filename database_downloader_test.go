package geoip_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-geoip"
)

func TestDatabaseDownloaderNew(t *testing.T) {

	dbPath := "database.mmdb"
	timeout := 1 * time.Second

	actual := geoip.NewDatabaseDownloader(dbPath, timeout)

	assert.NotNil(t, actual, "actual")
	assert.Equal(t, dbPath, actual.TargetFilePath)
	assert.Equal(t, geoip.DefaultDownloadURL, actual.DownloadURL)
	assert.Equal(t, geoip.DefaultChecksumURL, actual.ChecksumURL)

}

func TestDatabaseDownloaderLocalChecksumNoDBFile(t *testing.T) {

	downloader := geoip.NewDatabaseDownloader("db.mmdb", 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Empty(t, actual, "actual")

}

func TestDatabaseDownloaderLocalChecksumNoChecksumFile(t *testing.T) {

	defer filet.CleanUp(t)

	dbPath := "db.mmdb"

	filet.File(t, dbPath, "")

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Empty(t, actual, "actual")

}

func TestDatabaseDownloaderLocalChecksum_Valid(t *testing.T) {

	defer filet.CleanUp(t)

	expected := "checksum"

	dbPath := "db.mmdb"
	filet.File(t, dbPath, expected)

	checksumPath := dbPath + geoip.DefaultChecksumExt
	filet.File(t, checksumPath, expected)

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)

	actual, err := downloader.LocalChecksum()
	assert.NoError(t, err, "error")
	assert.Equal(t, expected, actual, "actual")

}

func TestDatabaseDownloaderRemoteChecksumValid(t *testing.T) {

	expected := "checksum"

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

func TestDatabaseDownloaderRemoteChecksumInvalidURL(t *testing.T) {

	downloader := geoip.NewDatabaseDownloader("", 5*time.Second)
	downloader.ChecksumURL = "ht&@-tp://:aa"

	actual, err := downloader.RemoteChecksum()
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestDatabaseDownloaderRemoteChecksumTimeout(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

func TestDatabaseDownloaderRemoteChecksumReadBodyError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

func TestDatabaseDownloaderShouldDownloadFalse(t *testing.T) {

	defer filet.CleanUp(t)

	expected := "checksum"

	dbPath := "db.mmdb"
	filet.File(t, dbPath, expected)

	checksumPath := dbPath + geoip.DefaultChecksumExt
	filet.File(t, checksumPath, expected)

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte(expected))
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)
	downloader.ChecksumURL = s.URL

	actual, err := downloader.ShouldDownload()
	assert.NoError(t, err, "error")
	assert.False(t, actual, "actual")

}

func TestDatabaseDownloaderShouldDownloadTrue(t *testing.T) {

	defer filet.CleanUp(t)

	expected := "checksum"

	dbPath := "db.mmdb"
	filet.File(t, dbPath, expected)

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte(expected))
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)
	downloader.ChecksumURL = s.URL

	actual, err := downloader.ShouldDownload()
	assert.NoError(t, err, "error")
	assert.True(t, actual, "actual")

}

func TestDatabaseDownloaderDownloadValid(t *testing.T) {

	expected := "expected"
	checksum := "checksum"

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata", "validdb.tgz")
	tgzData, _ := ioutil.ReadFile(path)

	dbPath := "db.mmdb"
	os.Remove(dbPath)
	defer os.Remove(dbPath)

	checksumPath := dbPath + geoip.DefaultChecksumExt
	os.Remove(checksumPath)
	defer os.Remove(checksumPath)

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.RequestURI, geoip.DefaultChecksumExt) {
				w.Write([]byte(checksum))
			} else {
				w.Write(tgzData)
			}
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)
	downloader.DownloadURL = s.URL + "/validdb.tgz"
	downloader.ChecksumURL = downloader.DownloadURL + geoip.DefaultChecksumExt

	err := downloader.Download()

	assert.NoError(t, err, "error")
	assert.FileExists(t, dbPath)
	assert.FileExists(t, checksumPath)

	filet.FileSays(t, dbPath, []byte(expected+"\n"))
	filet.FileSays(t, checksumPath, []byte(checksum))

}

func TestDatabaseDownloaderDownloadInvalidDownload(t *testing.T) {

	checksum := "checksum"

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata", "invaliddb.tgz")
	tgzData, _ := ioutil.ReadFile(path)

	dbPath := "db.mmdb"
	os.Remove(dbPath)
	defer os.Remove(dbPath)

	checksumPath := dbPath + geoip.DefaultChecksumExt
	os.Remove(checksumPath)
	defer os.Remove(checksumPath)

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.RequestURI, geoip.DefaultChecksumExt) {
				w.Write([]byte(checksum))
			} else {
				w.Write(tgzData)
			}
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader(dbPath, 1*time.Second)
	downloader.DownloadURL = s.URL + "/invaliddb.tgz"
	downloader.ChecksumURL = downloader.DownloadURL + geoip.DefaultChecksumExt

	err := downloader.Download()

	assert.Error(t, err, "error")

}

func TestDatabaseDownloaderDownloadInvalidURL(t *testing.T) {

	downloader := geoip.NewDatabaseDownloader("", 5*time.Second)
	downloader.DownloadURL = "ht&@-tp://:aa"

	err := downloader.Download()
	assert.Error(t, err)

}

func TestDatabaseDownloaderDownloadTimeout(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(500 * time.Millisecond)
			w.Write([]byte("checksum"))
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader("", 250*time.Millisecond)
	downloader.DownloadURL = s.URL

	err := downloader.Download()
	assert.Error(t, err)

}

func TestDatabaseDownloaderDownloadReadBodyError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "1")
		}),
	)
	defer s.Close()

	downloader := geoip.NewDatabaseDownloader("", 1*time.Second)
	downloader.DownloadURL = s.URL

	err := downloader.Download()
	assert.Error(t, err)

}
