# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o eas-inbound-filter .

# Runtime stage - Kong image
FROM kong:3.9-ubuntu

# Copy the compiled plugin
COPY --from=builder /build/eas-inbound-filter /usr/local/bin/eas-inbound-filter
USER root
# Make it executable
RUN chmod +x /usr/local/bin/eas-inbound-filter
RUN chown -R kong:kong /usr/local/kong
RUN chmod -R 755 /usr/local/kong
