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
RUN go build -o api-key-generator cmd/generator/api_secret/main.go

# Start with a minimal image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

ARG PORT

ENV PORT=$PORT

# Copy the built app from the builder stage
COPY --from=builder /app/wireguard-api /bin/wireguard-api
COPY --from=builder /app/api-key-generator /bin/api-key-generator

# Expose the port the app will run on
EXPOSE $PORT

# Command to run the app
CMD ["/bin/wireguard-api"]
