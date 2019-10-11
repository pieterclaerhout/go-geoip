package geoip

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// DefaultDownloadURL is the URL where to download the tgz file
const DefaultDownloadURL = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"

// DefaultChecksumURL is the URL where to download the checksum file
const DefaultChecksumURL = "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5"

// DefaultChecksumExt is the default extension for the local checksum
const DefaultChecksumExt = ".md5"

// DatabaseDownloader is a struct used to download the GeoLite database from maxmind
type DatabaseDownloader struct {
	TargetFilePath    string       // The path where to store the database
	localChecksumPath string       // The path where the local checksum is stored, defaults to target path + ".md5"
	DownloadURL       string       // The URL where to download the tgz file
	ChecksumURL       string       // The URL where to download the remote checksum
	httpClient        *http.Client // The HTTP client which is used to do the downloading
}

// NewDatabaseDownloader returns a new DatabaseDownloader instance configured with the default URLs
//
// The default URLs download the latest GeoLite2-City database from:
// https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5
func NewDatabaseDownloader(targetFilePath string, timeout time.Duration) *DatabaseDownloader {
	return &DatabaseDownloader{
		TargetFilePath:    targetFilePath,
		localChecksumPath: targetFilePath + DefaultChecksumExt,
		DownloadURL:       DefaultDownloadURL,
		ChecksumURL:       DefaultChecksumURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// LocalChecksum returns the local checksum if any
func (downloader *DatabaseDownloader) LocalChecksum() (string, error) {

	if !downloader.fileExists(downloader.TargetFilePath) {
		return "", nil
	}

	if !downloader.fileExists(downloader.localChecksumPath) {
		return "", nil
	}

	localChecksum, err := ioutil.ReadFile(downloader.localChecksumPath)
	if err != nil {
		return "", err
	}

	return string(localChecksum), nil

}

// RemoteChecksum returns the remote checksum
func (downloader *DatabaseDownloader) RemoteChecksum() (string, error) {

	resp, err := downloader.doGETRequest(downloader.ChecksumURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil

}

// ShouldDownload checks if a download is needed or not
//
// It does this by comparing the local and remote checksum.
// If they are different, a download is needed
func (downloader *DatabaseDownloader) ShouldDownload() (bool, error) {

	localChecksum, err := downloader.LocalChecksum()
	if err != nil {
		return false, err
	}

	remoteChecksum, err := downloader.RemoteChecksum()
	if err != nil {
		return false, err
	}

	return remoteChecksum != localChecksum, nil

}

// Download performs the actual download of the database
func (downloader *DatabaseDownloader) Download() error {

	resp, err := downloader.doGETRequest(downloader.DownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	uncompressedStream, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {

		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if !strings.HasSuffix(header.Name, ".mmdb") {
			continue
		}

		outFile, err := os.Create(downloader.TargetFilePath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			return err
		}

	}

	remoteChecksum, err := downloader.RemoteChecksum()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(downloader.localChecksumPath, []byte(remoteChecksum), 0666); err != nil {
		return err
	}

	return nil

}

// doRequest is a helper function to perform a GET request and return the HTTP response
func (downloader *DatabaseDownloader) doGETRequest(urlString string) (*http.Response, error) {

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Encoding", "")
	req.Header.Set("Connection", "close")
	req.Header.Set("Accept-Encoding", "deflate, identity")

	resp, err := downloader.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// fileExists is a helper function to check if a file exists or not
func (downloader *DatabaseDownloader) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
