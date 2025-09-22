# Use the official Golang image as a base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
