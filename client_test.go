package geoip_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pieterclaerhout/go-geoip/v2"
	"github.com/stretchr/testify/assert"
)

func TestClientLookup(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.Lookup("1.1.1.1")

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, "1.1.1.1", actual.IPAddress, "ipaddress")
	assert.Equal(t, false, actual.IsCached, "is-cached")

}

func TestClientLookupInvalid(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.Lookup("invalid")

	assert.Error(t, err)
	assert.Nil(t, actual)

}

func TestClientLookupInvalidURL(t *testing.T) {

	client := geoip.NewClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := client.Lookup("1.1.1.1")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientLookupTimeout(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(500 * time.Millisecond)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("hello"))
		}),
	)
	defer s.Close()

	client := geoip.NewClient(s.URL, 250*time.Millisecond)

	actual, err := client.Lookup("1.1.1.1")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientLookupReadBodyError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "1")
		}),
	)
	defer s.Close()

	client := geoip.NewClient(s.URL, 250*time.Millisecond)

	actual, err := client.Lookup("1.1.1.1")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientLookupInvalidResponse(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"hello"`))
		}),
	)
	defer s.Close()

	client := geoip.NewClient(s.URL, 250*time.Millisecond)

	actual, err := client.Lookup("1.1.1.1")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientLookupCache(t *testing.T) {

	s := testServer(t)

	client := geoip.NewClient(s.URL, 250*time.Millisecond)
	client.ClearCache()

	input := "213.118.8.79"

	nonCached, err := client.Lookup(input)
	assert.NotNil(t, nonCached, "nonCached")
	assert.False(t, nonCached.IsCached, "nonCached.IsCached")
	assert.NoErrorf(t, err, "nonCached")

	cached, err := client.Lookup(input)
	assert.NotNil(t, cached, "cached")
	assert.True(t, cached.IsCached, "cached.IsCached")
	assert.NoErrorf(t, err, "cached")

	client.ClearCache()

	clearCached, err := client.Lookup(input)
	assert.NotNil(t, clearCached, "clearCached")
	assert.False(t, clearCached.IsCached, "clearCached.IsCached")
	assert.NoErrorf(t, err, "clearCached")

}

func TestClientCountryCodeValid(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.CountryCode("1.1.1.1")

	assert.NoError(t, err)
	assert.Equal(t, "AU", actual)

}

func TestClientCountryCodeInvalid(t *testing.T) {

	client := geoip.NewClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := client.CountryCode("1.1.1.1")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientCountryNameValid(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.CountryName("1.1.1.1")

	assert.NoError(t, err)
	assert.Equal(t, "Australia", actual)

}

func TestClientCountryNameInvalid(t *testing.T) {

	client := geoip.NewClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := client.CountryName("1.1.1.1")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientRegionNameValid(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.RegionName("1.1.1.1")

	assert.NoError(t, err)
	assert.Equal(t, "southeast-asia", actual)

}

func TestClientRegionNameInvalid(t *testing.T) {

	client := geoip.NewClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := client.RegionName("1.1.1.1")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestClientTimeZoneValid(t *testing.T) {

	s := testServer(t)
	defer s.Close()

	client := geoip.NewClient(s.URL, 5*time.Second)

	actual, err := client.TimeZone("1.1.1.1")

	assert.NoError(t, err)
	assert.Equal(t, "Australia/Sydney", actual)

}

func TestClientTimeZoneInvalid(t *testing.T) {

	client := geoip.NewClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := client.TimeZone("1.1.1.1")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func testServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if r.URL.Query().Get("ip") == "invalid" {
				w.Write([]byte(`{"error": "invalid ip address"}`))
				return
			}
			w.Write([]byte(`{
				"IPAddress": "1.1.1.1",
				"Continent": {
				  "ISOCode": "OC",
				  "Names": {
					"de": "Ozeanien",
					"en": "Oceania",
					"es": "Oceanía",
					"fr": "Océanie",
					"ja": "オセアニア",
					"pt-BR": "Oceania",
					"ru": "Океания",
					"zh-CN": "大洋洲"
				  }
				},
				"Country": {
				  "ISOCode": "AU",
				  "Names": {
					"de": "Australien",
					"en": "Australia",
					"es": "Australia",
					"fr": "Australie",
					"ja": "オーストラリア",
					"pt-BR": "Austrália",
					"ru": "Австралия",
					"zh-CN": "澳大利亚"
				  }
				},
				"Location": {
				  "Latitude": -33.494,
				  "Longitude": 143.2104,
				  "TimeZone": "Australia/Sydney"
				},
				"Subdivisions": null,
				"IsCached": false
			  }`))
		}),
	)
}
