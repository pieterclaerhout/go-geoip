## STAGE 1 - MOD DOWNLOAD
FROM golang:1.13.1-alpine AS mod-download

RUN mkdir -p /app

ADD go.mod /app
ADD go.sum /app

WORKDIR /app

RUN go mod download

## STAGE 2 - BUILD
FROM mod-download AS builder

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -a --ldflags '-extldflags -static' -o geoip-server github.com/pieterclaerhout/go-geoip/cmd/geoip-server

# STAGE 3 - GEOIP DB

FROM alpine:latest AS geoip-db

RUN apk --no-cache add curl

RUN mkdir -p ./tmp
RUN curl -o ./tmp/download.tgz https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz
RUN tar xzf tmp/download.tgz -C tmp --strip-components 1
RUN mv tmp/GeoLite2-City.mmdb ./
RUN rm -rf tmp

# STAGE 4 - FINAL

FROM alpine:latest 

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/geoip-server /
RUN chmod a+x /geoip-server

COPY --from=geoip-db /GeoLite2-City.mmdb /

ENV GEOIP_DB=/GeoLite2-City.mmdb

ENTRYPOINT ["/geoip-server"]
EXPOSE 8080
