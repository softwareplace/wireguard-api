# Use the official Golang image
FROM golang:1.22.5-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire source code
COPY . .

# Build the Go app
RUN go build -o wireguard-api cmd/server/main.go

# Start with a minimal image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built app from the builder stage
COPY --from=builder /app/wireguard-api /bin/wireguard-api

# Expose the port the app will run on
EXPOSE 8080

# Command to run the app
CMD ["/bin/wireguard-api"]
