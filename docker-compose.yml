# Base image
FROM golang:1.21-alpine

# Set environment
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install dependencies
RUN apk update && apk add --no-cache git

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy project files
COPY . .

# Install air (for auto reload inside container)
RUN go install github.com/air-verse/air@latest

# Command to run app using air
CMD ["air"]
