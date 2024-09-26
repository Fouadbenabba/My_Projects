# Start from the official Golang image for building the application
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/myapp/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set up certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the app runs
EXPOSE 8080

# Command to run the application
CMD ["./main"]
