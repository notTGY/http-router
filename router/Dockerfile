FROM golang:1.17 AS builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=mipsle go mod tidy
RUN GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags="-extldflags=-static" -tags netgo -o http_server main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/http_server .
