## STAGE 1 - MOD DOWNLOAD
FROM golang:1.14.2-alpine AS mod-download

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

# STAGE 3 - FINAL

FROM alpine:latest 

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/geoip-server /
RUN chmod a+x /geoip-server

ENV GEOIP_DB=/GeoLite2-City.mmdb

ENTRYPOINT ["/geoip-server"]
EXPOSE 8080
