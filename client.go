package geoip

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client is used to get the results from the server API
type Client struct {
	url         string
	timeout     time.Duration
	lookupCache map[string]*IPLocation
}

// NewClient returns a new Client instance with the given URL
func NewClient(url string, timeout time.Duration) *Client {
	return &Client{
		url:         url,
		timeout:     timeout,
		lookupCache: map[string]*IPLocation{},
	}
}

// ClearCache clears the cache for the lookups
func (client *Client) ClearCache() {
	client.lookupCache = map[string]*IPLocation{}
}

// Lookup returns the full country information for a specific IP address
func (client *Client) Lookup(ipaddress string) (*IPLocation, error) {

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	if location, cached := client.lookupCache[ipaddress]; cached {
		location.IsCached = true
		return location, nil
	}

	params := url.Values{}
	params.Add("ip", ipaddress)

	fullURL := client.url + "?" + params.Encode()

	httpClient := &http.Client{}
	httpClient.Timeout = client.timeout

	resp, err := httpClient.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errResponse ErrorResponse
	if err := json.Unmarshal(rawData, &errResponse); err != nil {
		return nil, err
	}

	if errResponse.Error != "" {
		return nil, errors.New(errResponse.Error)
	}

	var location *IPLocation
	if err := json.Unmarshal(rawData, &location); err != nil {
		return nil, err
	}

	location.IsCached = false

	client.lookupCache[ipaddress] = location

	return location, nil

}

// CountryCode returns the country code for a specific IP address
func (client *Client) CountryCode(ipaddress string) (string, error) {
	location, err := client.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	return location.CountryCode(), nil
}

// CountryName returns the country name for a specific IP address
func (client *Client) CountryName(ipaddress string) (string, error) {
	location, err := client.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	return location.CountryName(), nil
}

// RegionName returns the region name for a specific IP address
//
// Region can be:
// - west-us
// - south-brazil
// - japan-east
// - southeast-asia
// - west-europe (the default)
func (client *Client) RegionName(ipaddress string) (string, error) {
	location, err := client.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	return location.RegionName(), nil
}

// TimeZone returns the timezone for a specific IP address
func (client *Client) TimeZone(ipaddress string) (string, error) {
	location, err := client.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	return location.TimeZone(), nil
}
