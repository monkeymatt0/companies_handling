# Use the official Go 1.22 Alpine image as the base for building
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Use a smaller image for the final stage (distroless or Alpine)
FROM alpine:latest

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder image
COPY --from=builder /app/main .

# Copy the .env and config.yaml files into the container
COPY cmd/.env .env
COPY cmd/config.yaml config.yaml

# Expose the port your application listens on
EXPOSE 8080

# Define the entry point to run your application
CMD ["./main"]