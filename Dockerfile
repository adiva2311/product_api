# Gunakan base image golang
FROM golang:1.23-alpine

# Set environment variable untuk Go
ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct

# Buat direktori kerja
WORKDIR /app

# Copy semua file ke container
COPY .env .env
COPY . .

# Download dependency
RUN go mod tidy

# Build aplikasi
RUN go build -o product_api

# Expose port
EXPOSE 8080

# Jalankan aplikasi
CMD ["./product_api"]
