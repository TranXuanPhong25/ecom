# ===== STAGE 1: Build the Go binary =====
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# ===== STAGE 2: Build a minimal image =====
FROM alpine:latest AS release

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
