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

type DatabaseDownloader struct {
	TargetFilePath    string
	localChecksumPath string
	downloadURL       string
	checksumURL       string
	httpClient        *http.Client
}

func NewDatabaseDownloader(targetFilePath string, timeout time.Duration) *DatabaseDownloader {
	return &DatabaseDownloader{
		TargetFilePath:    targetFilePath,
		localChecksumPath: targetFilePath + ".md5",
		downloadURL:       "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz",
		checksumURL:       "https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5",
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (downloader *DatabaseDownloader) LocalChecksum() (string, error) {

	if !downloader.fileExists(downloader.TargetFilePath) {
		return "", nil
	}

	if !downloader.fileExists(downloader.localChecksumPath) {
		return "", nil
	}

	// file, err := os.Open(downloader.TargetFilePath)
	// if err != nil {
	// 	return "", err
	// }
	// defer file.Close()

	localChecksum, err := ioutil.ReadFile(downloader.localChecksumPath)
	if err != nil {
		return "", err
	}

	return string(localChecksum), nil

}

func (downloader *DatabaseDownloader) RemoteChecksum() (string, error) {

	resp, err := downloader.doRequest(downloader.checksumURL)
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

func (downloader *DatabaseDownloader) Download() error {

	resp, err := downloader.doRequest(downloader.downloadURL)
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

func (downloader *DatabaseDownloader) doRequest(urlString string) (*http.Response, error) {

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

func (downloader *DatabaseDownloader) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
