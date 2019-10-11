REVISION := $(shell git rev-parse --short HEAD)

build-server:
	@go build -trimpath -ldflags "-s -w" -o geoip-server github.com/pieterclaerhout/go-geoip/cmd/geoip-server

build-db-downloader:
	@go build -trimpath -ldflags "-s -w" -o db-downloader github.com/pieterclaerhout/go-geoip/cmd/db-downloader

run-server: build-server
	@PORT=8080 GEOIP_DB=GeoLite2-City.mmdb DEBUG=1 ./geoip-server

run-db-downloader: build-db-downloader
	@DEBUG=1 ./db-downloader

build-docker-image:
	docker build -t geoip-server .

publish-docker-image: build-docker-image
	docker tag geoip-server pieterclaerhout/geoip-server:$(REVISION)
	docker tag geoip-server pieterclaerhout/geoip-server:latest
	docker push pieterclaerhout/geoip-server:$(REVISION)
	docker push pieterclaerhout/geoip-server:latest
	
run-docker-image: build-docker-image
	docker run --rm -p 8080:8080 geoip-server
