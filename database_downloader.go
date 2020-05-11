package geoip

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DefaultDownloadURL is the URL where to download the tgz file
const DefaultDownloadURL = "https://download.maxmind.com/app/geoip_download?suffix=tar.gz"

// DefaultChecksumURL is the URL where to download the checksum file
const DefaultChecksumURL = "https://download.maxmind.com/app/geoip_download?suffix=tar.gz.sha256"

// DefaultChecksumExt is the default extension for the local checksum
const DefaultChecksumExt = ".sha256"

// DatabaseDownloader is a struct used to download the GeoLite database from maxmind
type DatabaseDownloader struct {
	LicenseKey        string       // The license key to use for the downloads
	TargetFilePath    string       // The path where to store the database
	localChecksumPath string       // The path where the local checksum is stored, defaults to target path + ".sha256"
	DownloadURL       string       // The URL where to download the tgz file
	ChecksumURL       string       // The URL where to download the remote checksum
	httpClient        *http.Client // The HTTP client which is used to do the downloading
}

// NewDatabaseDownloader returns a new DatabaseDownloader instance configured with the default URLs
func NewDatabaseDownloader(licenseKey string, targetFilePath string, timeout time.Duration) *DatabaseDownloader {
	return &DatabaseDownloader{
		LicenseKey:        licenseKey,
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

	foundFile := false

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

		targetFileDir := filepath.Dir(downloader.TargetFilePath)
		if !downloader.fileExists(targetFileDir) {
			if err := os.MkdirAll(targetFileDir, 0755); err != nil {
				return err
			}
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

		remoteChecksum, err := downloader.RemoteChecksum()
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(downloader.localChecksumPath, []byte(remoteChecksum), 0666); err != nil {
			return err
		}

		foundFile = true

	}

	if foundFile {
		return nil
	}

	return errors.New("Invalid download, tgz doesn't contain a .mmdb file")

}

// doRequest is a helper function to perform a GET request and return the HTTP response
func (downloader *DatabaseDownloader) doGETRequest(urlString string) (*http.Response, error) {

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	q := parsedURL.Query()
	q.Set("edition_id", "GeoLite2-City")
	q.Set("license_key", downloader.LicenseKey)
	parsedURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
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

	if resp.StatusCode == 401 {
		return nil, errors.New("Invalid license key")
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
