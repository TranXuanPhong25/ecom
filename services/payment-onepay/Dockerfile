# ===== STAGE 1: Build the Go binary =====
FROM golang:1.24-alpine AS builder

# Install git (needed for go get module from github)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum first to optimize Docker layer cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build app
RUN go build -o main .

# ===== STAGE 2: Build a minimal image =====
FROM alpine:latest AS release

WORKDIR /root/

# Copy binary from build stage
COPY --from=builder /app/main .

# Run the app
CMD ["./main"]
