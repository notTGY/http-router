FROM golang:1.13 AS builder

WORKDIR /app
COPY http_server.go .

RUN GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags="-extldflags=-static" -tags netgo -o http_server http_server.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/http_server .
