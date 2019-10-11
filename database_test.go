package geoip_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-geoip"
)

func Test_Database_Lookup(t *testing.T) {

	type test struct {
		input       string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", false},
		{"168.63.67.114", false},
		{"52.174.239.61", false},
		{"1.1.1.1", false},
		{"127.0.0.1", true},
		{"153.232.156.201", false},
		{"138.197.69.54", false},
		{"5.145.169.67", false},
		{"170.84.87.234", false},
		{"116.209.59.102", false},
		{"invalid", true},
		{"", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			actual, err := db.Lookup(tc.input)

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.NoErrorf(t, err, tc.input)
				// assert.NotNil(t, actual, tc.input)
				// assert.Equalf(t, tc.input, actual.IPAddress, tc.input)
			}

		})

	}

}

func Test_Database_CountryCode(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "BE", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "BE", false},
		{"168.63.67.114", "IE", false},
		{"52.174.239.61", "NL", false},
		{"1.1.1.1", "AU", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "JP", false},
		{"138.197.69.54", "US", false},
		{"5.145.169.67", "ES", false},
		{"170.84.87.234", "BR", false},
		{"116.209.59.102", "CN", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			actual, err := db.CountryCode(tc.input)

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.Equalf(t, tc.expected, actual, tc.input)
				assert.NoErrorf(t, err, tc.input)
			}

		})

	}
}

func Test_Database_CountryName(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "Belgium", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "Belgium", false},
		{"168.63.67.114", "Ireland", false},
		{"52.174.239.61", "Netherlands", false},
		{"1.1.1.1", "Australia", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "Japan", false},
		{"138.197.69.54", "United States", false},
		{"5.145.169.67", "Spain", false},
		{"170.84.87.234", "Brazil", false},
		{"116.209.59.102", "China", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			actual, err := db.CountryName(tc.input)

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.Equalf(t, tc.expected, actual, tc.input)
				assert.NoErrorf(t, err, tc.input)
			}

		})

	}
}

func Test_Database_RegionName(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "west-europe", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "west-europe", false},
		{"168.63.67.114", "west-europe", false},
		{"52.174.239.61", "west-europe", false},
		{"1.1.1.1", "southeast-asia", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "japan-east", false},
		{"138.197.69.54", "west-us", false},
		{"5.145.169.67", "west-europe", false},
		{"170.84.87.234", "south-brazil", false},
		{"116.209.59.102", "southeast-asia", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			actual, err := db.RegionName(tc.input)

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.Equalf(t, tc.expected, actual, tc.input)
				assert.NoErrorf(t, err, tc.input)
			}

		})

	}

}

func Test_Database_TimeZone(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "Europe/Brussels", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "Europe/Brussels", false},
		{"168.63.67.114", "Europe/Dublin", false},
		{"52.174.239.61", "Europe/Amsterdam", false},
		{"1.1.1.1", "Australia/Sydney", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "Asia/Tokyo", false},
		{"138.197.69.54", "America/New_York", false},
		{"5.145.169.67", "Europe/Madrid", false},
		{"170.84.87.234", "America/Recife", false},
		{"116.209.59.102", "Asia/Shanghai", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			actual, err := db.TimeZone(tc.input)

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.Equalf(t, tc.expected, actual, tc.input)
				assert.NoErrorf(t, err, tc.input)
			}

		})

	}

}

func Test_Database_DBWithPath_Invalid(t *testing.T) {

	db := geoip.NewDatabase("invalid-path")

	_, err := db.Lookup("213.118.8.79")

	assert.Error(t, err)

}

func Test_Database_DefaultDatabase_Cache(t *testing.T) {

	db := openTestDatabase(t)
	db.ClearCache()

	input := "213.118.8.79"

	nonCached, err := db.Lookup(input)
	assert.NotNil(t, nonCached, "nonCached")
	assert.False(t, nonCached.IsCached, "nonCached.IsCached")
	assert.NoErrorf(t, err, "nonCached")

	cached, err := db.Lookup(input)
	assert.NotNil(t, cached, "cached")
	assert.True(t, cached.IsCached, "cached.IsCached")
	assert.NoErrorf(t, err, "cached")

	db.ClearCache()

	clearCached, err := db.Lookup(input)
	assert.NotNil(t, clearCached, "clearCached")
	assert.False(t, clearCached.IsCached, "clearCached.IsCached")
	assert.NoErrorf(t, err, "clearCached")

}

func openTestDatabase(t *testing.T) *geoip.Database {
	t.Helper()
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "testdata", "GeoLite2-City.mmdb")
	db := geoip.NewDatabase(path)
	return db
}
