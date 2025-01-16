# Gunakan base image Golang
FROM golang:1.23.4-alpine

# Set Working Directory di dalam Container
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk caching layer dependencies
COPY go.mod go.sum ./

# Unduh dependencies agar tetap cache-efficient
RUN go mod tidy

# Copy seluruh kode sumber ke dalam container
COPY . .

# Install Air untuk hot reload
RUN go install github.com/cosmtrek/air@v1.45.0

# Expose port gRPC
EXPOSE 50051

# Jalankan aplikasi dengan Air untuk hot reload
CMD ["air"]
