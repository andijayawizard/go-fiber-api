# Gunakan base image Go 1.24.2 (fix masalah versi)
FROM golang:1.24.2-alpine

# Install dependencies
RUN apk update && apk add --no-cache git

# Set workdir
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu (optimize cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Buat folder tmp/ untuk binary
RUN mkdir -p tmp

# Install air
RUN go install github.com/air-verse/air@latest

# Build awal dummy (supaya tmp/main ada)
RUN go build -o ./tmp/main

# Jalankan Air (auto reload)
CMD ["air"]
# CMD ["air", "-c", ".air.dev.toml"]
