# Use the official Golang image to build the server
FROM golang:1.18 as builder

# Set the working directory
WORKDIR /app

# Copy the server code
COPY . .

# Build the server binary
RUN go build -o server ./server/main.go

# Use a lightweight image for running the server
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the server binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose the server port
EXPOSE 8080

# Set entrypoint to run the server
ENTRYPOINT ["./server"]
