# go-geoip

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-geoip)](https://goreportcard.com/report/github.com/pieterclaerhout/go-geoip)
[![Documentation](https://godoc.org/github.com/pieterclaerhout/go-geoip?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-geoip)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-geoip/raw/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-geoip.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-geoip)
[![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-geoip.svg)](https://github.com/pieterclaerhout/go-geoip/issues)

This is a wrapper around the GeoIP databases from [MaxMind](https://www.maxmind.com/en/home).

It also has a web API which automatically keeps the GeoIP database up-to-date.

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
PORT=8080 GEOIP_DB=testdata/GeoLite2-City.mmdb ./geoip-server
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
    "github.com/pieterclaerhout/go-geoip"
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