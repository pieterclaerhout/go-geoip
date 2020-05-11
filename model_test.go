package geoip_test

import (
	"testing"

	"github.com/pieterclaerhout/go-geoip"
	"github.com/stretchr/testify/assert"
)

func TestModelContinentCode(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "EU", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "EU", false},
		{"168.63.67.114", "EU", false},
		{"52.174.239.61", "EU", false},
		{"1.1.1.1", "OC", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "AS", false},
		{"138.197.69.54", "NA", false},
		{"5.145.169.67", "EU", false},
		{"170.84.87.234", "SA", false},
		{"116.209.59.102", "AS", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			info, err := db.Lookup(tc.input)

			actual := info.ContinentCode()

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

func TestModelContinentName(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", "Europe", false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", "Europe", false},
		{"168.63.67.114", "Europe", false},
		{"52.174.239.61", "Europe", false},
		{"1.1.1.1", "Oceania", false},
		{"127.0.0.1", "", true},
		{"153.232.156.201", "Asia", false},
		{"138.197.69.54", "North America", false},
		{"5.145.169.67", "Europe", false},
		{"170.84.87.234", "South America", false},
		{"116.209.59.102", "Asia", false},
		{"invalid", "", true},
		{"", "", true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			info, err := db.Lookup(tc.input)
			actual := info.ContinentName()

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

func TestModelSubdivisionCodes(t *testing.T) {

	type test struct {
		input       string
		expected    []string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", []string{"VLG", "VWV"}, false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", []string{"VLG", "VWV"}, false},
		{"168.63.67.114", []string{"L"}, false},
		{"52.174.239.61", []string{"NH"}, false},
		{"1.1.1.1", []string{}, false},
		{"127.0.0.1", []string{}, true},
		{"153.232.156.201", []string{"12"}, false},
		{"138.197.69.54", []string{"NJ"}, false},
		{"5.145.169.67", []string{}, false},
		{"170.84.87.234", []string{"PE"}, false},
		{"invalid", []string{}, true},
		{"", []string{}, true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			info, err := db.Lookup(tc.input)
			actual := info.SubdivisionCodes()

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

func TestModelSubdivisionNames(t *testing.T) {

	type test struct {
		input       string
		expected    []string
		shouldError bool
	}

	var tests = []test{
		{"213.118.8.79", []string{"Flanders", "West Flanders Province"}, false},
		{"2a02:1811:b212:3500:f938:7fca:5c2b:6595", []string{"Flanders", "West Flanders Province"}, false},
		{"168.63.67.114", []string{"Leinster"}, false},
		{"52.174.239.61", []string{"North Holland"}, false},
		{"1.1.1.1", []string{}, false},
		{"127.0.0.1", []string{}, true},
		{"153.232.156.201", []string{"Chiba"}, false},
		{"138.197.69.54", []string{"New Jersey"}, false},
		{"5.145.169.67", []string{}, false},
		{"170.84.87.234", []string{"Pernambuco"}, false},
		{"invalid", []string{}, true},
		{"", []string{}, true},
	}

	db := openTestDatabase(t)

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			db.ClearCache()

			info, err := db.Lookup(tc.input)
			actual := info.SubdivisionNames()

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

func TestModelApproximateGPSCoordinate(t *testing.T) {

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

			info, err := db.Lookup(tc.input)
			actual := info.ApproximateGPSCoordinate()

			if tc.shouldError {
				assert.Emptyf(t, actual, tc.input)
				assert.Errorf(t, err, tc.input)
			} else {
				assert.NotNil(t, actual, tc.input)
				assert.NotZero(t, actual.Latitude, tc.input)
				assert.NotZero(t, actual.Longitude, tc.input)
				assert.NoErrorf(t, err, tc.input)
			}

		})

	}

}

func TestModelDefaultCountryCode(t *testing.T) {
	location := &geoip.IPLocation{}
	assert.Equal(t, "BE", location.CountryCode())
}

func TestModelDefaultCountryName(t *testing.T) {
	location := &geoip.IPLocation{}
	assert.Equal(t, "Belgium", location.CountryName())
}

func TestModelDefaultTimeZone(t *testing.T) {
	location := &geoip.IPLocation{}
	assert.Equal(t, "Europe/Brussels", location.TimeZone())
}
