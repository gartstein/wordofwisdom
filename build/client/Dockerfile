# Use the official Golang image to build the client
FROM golang:1.18 as builder

# Set the working directory
WORKDIR /app

# Copy the client code
COPY . .

# Build the client binary
RUN go build -o client ./client/main.go

# Use a lightweight image for running the client
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the client binary from the builder stage
COPY --from=builder /app/client /app/client

# Set entrypoint to run the client
ENTRYPOINT ["./client"]
