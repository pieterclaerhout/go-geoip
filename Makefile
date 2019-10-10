download-testdata:
	rm -rf testdata
	mkdir -p testdata
	mkdir -p tmp
	curl --silent https://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.tar.gz > tmp/GeoLite2-Country.tar.gz
	tar xzf tmp/GeoLite2-Country.tar.gz -C tmp --strip-components 1
	mv tmp/GeoLite2-Country.mmdb testdata/
	rm -rf tmp/
