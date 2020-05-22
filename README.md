# go-geoip

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-geoip)](https://goreportcard.com/report/github.com/pieterclaerhout/go-geoip)
[![Documentation](https://godoc.org/github.com/pieterclaerhout/go-geoip?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-geoip)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-geoip/raw/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-geoip.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-geoip)
[![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-geoip.svg)](https://github.com/pieterclaerhout/go-geoip/issues)

This is a wrapper around the GeoIP databases from [MaxMind](https://www.maxmind.com/en/home).

It also has a web API which automatically keeps the GeoIP database up-to-date. This is a very handy tool when you use e.g. a microservice approach and you don't want to keep a copy of the database for each microservice which needs GeoIP capabilities. By using the server approach, you can keep this functionality in a central place.

## Building

```
make build-server
```

## Building docker image

```
make build-docker-image
```

## Running manually

```
PORT=8080 GEOIP_DB=testdata/GeoLite2-City.mmdb LICENSE_KEY=license ./geoip-server
```

## Running via docker

The [docker image](https://hub.docker.com/r/pieterclaerhout/geoip-server) automatically includes a copy of the [GeoLite2-City.mmdb](https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz) database.

```
docker run --rm pieterclaerhout/geoip-server
```

## Using the server

```bash
$ curl "http://localhost:8080/lookup?ip=1.1.1.1"
{
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
}
```

## Using the client

```go
package main

import (
    "fmt"
    "os"
    "time"
    
    geoip "github.com/pieterclaerhout/go-geoip/v2"
)

func main() {

    client := geoip.NewClient("http://localhost:8080/  lookup", 5*time.Second)
  
    actual, err := client.Lookup("1.1.1.1")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
  
    fmt.Printf("%v\n", actual)

}
```

## About the license key

To run the GeoIP server and use the automatic database download, you need to obtain a license key from MaxMind as explained [here](https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases/).

It's an easy and straightforward process:

1. [Sign up for a MaxMind account](https://www.maxmind.com/en/geolite2/signup) (no purchase required)

2. Set your password and create a [license key](https://www.maxmind.com/en/accounts/current/license-key)

3. Use the environment variable `LICENSE_KEY` to set the license key which needs to be used.