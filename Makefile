init:
	@mkdir -p build

build-server: init
	@go build -trimpath -ldflags "-s -w" -o build/geoip-server github.com/pieterclaerhout/go-geoip/v2/cmd/geoip-server

build-db-downloader: init
	@go build -trimpath -ldflags "-s -w" -o build/db-downloader github.com/pieterclaerhout/go-geoip/v2/cmd/db-downloader

publish-docker-image:
	@go-james docker-image
	
run-docker: publish-docker-image
	@docker run --rm -p 8080:8080 pieterclaerhout/geoip-server
