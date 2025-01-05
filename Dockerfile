# Use a multi-stage build to minimize the final image size

# Stage 1: Build the Go HTTP server
FROM golang:1.23.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY http_server.go .

# Set the environment variables for cross-compilation
ENV GOOS=linux
ENV GOARCH=mips
ENV CGO_ENABLED=0

# Build the Go application
RUN go build -o http_server http_server.go

# Stage 2: Create a minimal image to hold the compiled binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/http_server .

# Expose port 80
EXPOSE 80

# Command to run the HTTP server
CMD ["./http_server"]
