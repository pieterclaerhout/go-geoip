REVISION := $(shell git rev-parse --short HEAD)

init:
	@mkdir -p build

build-server: init
	@go build -trimpath -ldflags "-s -w" -o build/geoip-server github.com/pieterclaerhout/go-geoip/v2/cmd/geoip-server

build-db-downloader: init
	@go build -trimpath -ldflags "-s -w" -o build/db-downloader github.com/pieterclaerhout/go-geoip/v2/cmd/db-downloader

build-docker-image:
	docker build -t geoip-server .
	# go-james docker-image

publish-docker-image: build-docker-image
	docker tag geoip-server pieterclaerhout/geoip-server:$(REVISION)
	docker push pieterclaerhout/geoip-server:$(REVISION)
	
run-docker-image: build-docker-image
	docker run --rm -p 8080:8080 geoip-server
