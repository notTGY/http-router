FROM golang:1.13 AS builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=mipsle go install github.com/tarm/serial
RUN GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags="-extldflags=-static" -tags netgo -o http_server main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/http_server .
