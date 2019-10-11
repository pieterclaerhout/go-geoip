download-testdata:
	rm -rf testdata
	mkdir -p testdata
	mkdir -p tmp
	curl https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz > tmp/GeoLite2-City.tar.gz
	tar xzf tmp/GeoLite2-City.tar.gz -C tmp --strip-components 1
	mv tmp/GeoLite2-City.mmdb testdata/
	rm -rf tmp/

build-server:
	@go build -trimpath -ldflags "-s -w" -o geoip-server github.com/pieterclaerhout/go-geoip/cmd/geoip-server

run-server: build-server
	@PORT=:8080 GEOIP_DB=testdata/GeoLite2-City.mmdb ./geoip-server

build-docker-image:
	docker build -t geoip-server .
	
run-docker-image: build-docker-image
	docker run --rm -p 8080:8080 geoip-server
