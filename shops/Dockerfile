# ===== STAGE 1: Build the Go binary =====
FROM golang:1.23-alpine AS builder

# Install git (nếu dùng go get lấy module từ github)
RUN apk add --no-cache git 

# Set working directory
WORKDIR /app

# Copy go mod và sum trước để optimize cache Docker layer
COPY go.mod go.sum  ./

# Download dependencies
RUN go mod download

# Copy toàn bộ code
COPY . .

# Build app
RUN go build -o main .

# ===== STAGE 2: Build a minimal image =====
FROM alpine:latest AS release

# Tạo thư mục app
WORKDIR /root/

# Copy binary từ stage build
COPY --from=builder /app/main .

EXPOSE 8203

# Lệnh chạy app
CMD ["./main"]
