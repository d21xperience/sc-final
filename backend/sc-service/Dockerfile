# Gunakan image resmi Go
FROM golang:1.22.3 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o main .

# Image runtime
FROM debian:bullseye-slim

# Install dependencies runtime (jika diperlukan)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary dari builder stage
COPY --from=builder /app/main /usr/local/bin/

# Set entrypoint
ENTRYPOINT ["main"]

# Gunakan port environment variable (opsional)
EXPOSE 8080
