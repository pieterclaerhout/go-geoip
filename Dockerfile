FROM golang:1.13.1-alpine AS builder

ADD . /app/backend
WORKDIR /app/backend

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /geoip-server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /geoip-server ./
RUN chmod +x ./geoip-server

ENTRYPOINT ["./geoip-server"]
EXPOSE 8080