FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

ARG PORT
ENV PORT=$PORT

# Copy the built app from the builder stage
COPY ./.temp/wireguard-api /bin/wireguard-api
COPY ./.temp/api-key-generator /bin/api-key-generator

# Expose the port the app will run on
EXPOSE $PORT

# Command to run the app
CMD ["/bin/wireguard-api"]
