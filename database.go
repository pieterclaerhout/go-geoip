package geoip

import (
	"sync"

	"github.com/oschwald/maxminddb-golang"
	"github.com/pkg/errors"
)

// Database is the main wrapper around the MaxMind GeoIP database
type Database struct {
	path             string
	lookupCache      map[string]*IPLocation
	lookupCacheMutex sync.RWMutex
}

// NewDatabase returns a new Database instance with the given database path
func NewDatabase(path string) *Database {
	return &Database{
		path:        path,
		lookupCache: map[string]*IPLocation{},
	}
}

// ClearCache clears the cache for the lookups
func (database *Database) ClearCache() {
	database.lookupCacheMutex.Lock()
	defer database.lookupCacheMutex.Unlock()
	database.lookupCache = map[string]*IPLocation{}
}

// Lookup returns the full country information for a specific IP address
func (database *Database) Lookup(ipaddress string) (*IPLocation, error) {

	if ipaddress == "" {
		return nil, nil
	}

	location, err := database.lookupFromCache(ipaddress)
	if err != nil {
		return nil, err
	}
	if location != nil {
		return location, nil
	}

	var record interface{}

	db, err := maxminddb.Open(database.path)
	if err != nil {
		return location, err
	}
	defer db.Close()

	ip, private, err := isPrivateIP(ipaddress)
	if err != nil {
		return nil, err
	}
	if private {
		return nil, nil
	}

	if err := db.Lookup(ip, &location); err != nil {
		return location, err
	}

	_ = db.Lookup(ip, &record)

	if location == nil {
		return location, errors.New("No info for: " + ipaddress)
	}

	location.IPAddress = ipaddress
	location.IsCached = false

	database.lookupCacheMutex.Lock()
	defer database.lookupCacheMutex.Unlock()
	database.lookupCache[ipaddress] = location

	return location, nil

}

func (database *Database) lookupFromCache(ipaddress string) (*IPLocation, error) {

	database.lookupCacheMutex.RLock()
	defer database.lookupCacheMutex.RUnlock()

	if location, cached := database.lookupCache[ipaddress]; cached {
		location.IsCached = true
		return location, nil
	}

	return nil, nil

}

// CountryCode returns the country code for a specific IP address
func (database *Database) CountryCode(ipaddress string) (string, error) {
	location, err := database.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	if location == nil {
		return "", nil
	}
	return location.CountryCode(), nil
}

// CountryName returns the country name for a specific IP address
func (database *Database) CountryName(ipaddress string) (string, error) {
	location, err := database.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	if location == nil {
		return "", nil
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
func (database *Database) RegionName(ipaddress string) (string, error) {
	location, err := database.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	if location == nil {
		return "", nil
	}
	return location.RegionName(), nil
}

// TimeZone returns the timezone for a specific IP address
func (database *Database) TimeZone(ipaddress string) (string, error) {
	location, err := database.Lookup(ipaddress)
	if err != nil {
		return "", err
	}
	if location == nil {
		return "", nil
	}
	return location.TimeZone(), nil
}
