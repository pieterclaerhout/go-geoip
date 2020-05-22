package geoip_test

import (
	"testing"

	"github.com/pieterclaerhout/go-geoip/v2"
	"github.com/stretchr/testify/assert"
)

func TestCountryCodeToName(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"BE", "Belgium", false},
		{"be", "Belgium", false},
		{"XX", "", true},
	}

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			actual, err := geoip.CountryCodeToName(tc.input)

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

func TestCountryNameToCode(t *testing.T) {

	type test struct {
		input       string
		expected    string
		shouldError bool
	}

	var tests = []test{
		{"Belgium", "BE", false},
		{"belgium", "BE", false},
		{"Non-existing country", "", true},
	}

	for _, tc := range tests {

		t.Run(tc.input, func(t *testing.T) {

			actual, err := geoip.CountryNameToCode(tc.input)

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
