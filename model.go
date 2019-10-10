package geoip

// Continent defines the continent information for an IP location
type Continent struct {
	ISOCode string            `maxminddb:"code"`  // The continent code
	Names   map[string]string `maxminddb:"names"` // The names for this continent
}

// Country defines the country information for an IP location
type Country struct {
	ISOCode string            `maxminddb:"iso_code"` // The ISO country code
	Names   map[string]string `maxminddb:"names"`    // The names for this country
}

// GPSCoordinate defines a GPS coordinate
type GPSCoordinate struct {
	Latitude  float64 `maxminddb:"latitude"`  // The latitude for this location
	Longitude float64 `maxminddb:"longitude"` // The longitude for this location
}

// Location defines all the location information we know such as timezone, lat and lon
type Location struct {
	GPSCoordinate
	TimeZone string `maxminddb:"time_zone"` // The timezone for this location
}

// Subdivision defines a subdivision of the country (state, province)
type Subdivision struct {
	ISOCode string            `maxminddb:"iso_code"` // The ISO subdivision code
	Names   map[string]string `maxminddb:"names"`    // The names for this subdivision
}

// IPLocation defines all info we know about a location based on it's IP address
type IPLocation struct {
	IPAddress    string
	Continent    *Continent     `maxminddb:"continent"`
	Country      *Country       `maxminddb:"country"`
	Location     *Location      `maxminddb:"location"`
	Subdivisions []*Subdivision `maxminddb:"subdivisions"`
	IsCached     bool
}

// CountryCode returns the ISO country code for the location
func (ipLocation *IPLocation) CountryCode() string {
	if ipLocation == nil || ipLocation.Country == nil {
		return defaultCountryCode
	}
	return ipLocation.Country.ISOCode
}

// CountryName returns the country name for the location in English
func (ipLocation *IPLocation) CountryName() string {
	if ipLocation == nil || ipLocation.Country == nil {
		return defaultCountryName
	}
	return ipLocation.Country.Names["en"]
}

// ContinentCode returns the ISO country code for the location
func (ipLocation *IPLocation) ContinentCode() string {
	if ipLocation == nil || ipLocation.Continent == nil {
		return ""
	}
	return ipLocation.Continent.ISOCode
}

// ContinentName returns the country name for the location in English
func (ipLocation *IPLocation) ContinentName() string {
	if ipLocation == nil || ipLocation.Continent == nil {
		return ""
	}
	return ipLocation.Continent.Names["en"]
}

// RegionName returns the region name for the location in English
func (ipLocation *IPLocation) RegionName() string {
	countryCode := ipLocation.CountryCode()
	return CountryCodeToRegion(countryCode)
}

// TimeZone returns the timezone for the location
func (ipLocation *IPLocation) TimeZone() string {
	if ipLocation == nil || ipLocation.Location == nil {
		return defaultTimeZone
	}
	return ipLocation.Location.TimeZone
}

// SubdivisionCodes returns the codes for the subdivisions
func (ipLocation *IPLocation) SubdivisionCodes() []string {
	if ipLocation == nil || ipLocation.Subdivisions == nil {
		return []string{}
	}
	result := []string{}
	for _, item := range ipLocation.Subdivisions {
		result = append(result, item.ISOCode)
	}
	return result
}

// SubdivisionNames returns the names for the subdivisions
func (ipLocation *IPLocation) SubdivisionNames() []string {
	if ipLocation == nil || ipLocation.Subdivisions == nil {
		return []string{}
	}
	result := []string{}
	for _, item := range ipLocation.Subdivisions {
		result = append(result, item.Names["en"])
	}
	return result
}

// ApproximateGPSCoordinate returns the GPS coordinates for the location
func (ipLocation *IPLocation) ApproximateGPSCoordinate() *GPSCoordinate {
	if ipLocation == nil || ipLocation.Location == nil {
		return nil
	}
	return &ipLocation.Location.GPSCoordinate
}
