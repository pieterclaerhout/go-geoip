## STAGE 1 - MOD DOWNLOAD
FROM golang:1.14.3-alpine AS mod-download

RUN mkdir -p /app

ADD go.mod /app
ADD go.sum /app

WORKDIR /app
RUN go mod download


## STAGE 2 - PREBUILD
FROM mod-download AS builder

ADD . /app

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -a --ldflags '-extldflags -static -s -w' -o geoip-server github.com/pieterclaerhout/go-geoip/v2/cmd/geoip-server


# STAGE 4 - FINAL
FROM alpine:3.11 

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/geoip-server /
RUN chmod a+x /geoip-server

ENV GEOIP_DB=/GeoLite2-City.mmdb

ENTRYPOINT ["/geoip-server"]
EXPOSE 8080
