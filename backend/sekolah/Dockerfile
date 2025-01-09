# Base image
FROM golang:1.22.3 AS builder

# Set environment variables
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Set working directory
WORKDIR /app

# Copy go.mod, go.sum, dan dependensi lokal
COPY auth_service/go.mod auth_service/go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh source code
COPY auth_service/. .

# Expose port gRPC
EXPOSE 50051

# Jalankan aplikasi dengan Air untuk hot reload
CMD ["air"]