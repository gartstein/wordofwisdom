# Use the official Golang image to build the server
FROM golang:1.18 AS builder

# Set the working directory
WORKDIR /app

# Copy the server code
COPY . .

# Build the server binary
RUN go build -o bin ./client/main.go

# Use a lightweight image for running the server
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the server binary from the builder stage
COPY --from=builder /app/bin /app/bin

# Expose the server port
EXPOSE 8080

# Set entrypoint to run the server binary
ENTRYPOINT ["./bin"]
