# 🏗️ Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS builder

# Install required tools
RUN apk add --no-cache git && \
    apk add --no-cache libc6-compat

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy full project
COPY . .

# Build statically-linked Go binary (no CGO, stripped binary)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# 🪶 Optional check size
# RUN du -h main

# 🧼 Stage 2: Tiny scratch image
FROM scratch

# Copy binary only, nothing else
COPY --from=builder /app/main /main

# Expose API port
EXPOSE 8080

# Run binary
ENTRYPOINT ["/main"]
