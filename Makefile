# Makefile for managing the Word of Wisdom server and client

# Variables
SERVER_IMAGE_NAME = wordofwisdom-server
CLIENT_IMAGE_NAME = wordofwisdom-client
SERVER_CONTAINER_NAME = wordofwisdom-server
CLIENT_CONTAINER_NAME = wordofwisdom-client
SERVER_DOCKERFILE_PATH = ./build/server/Dockerfile
CLIENT_DOCKERFILE_PATH = ./build/client/Dockerfile
SERVER_PORT = 8080

.PHONY: all build run stop clean run-client

# Default target: Build and run both server and client
all: build run

# Build Docker images
build: build-server build-client

build-server:
	@echo "Building the server Docker image..."
	docker build -t $(SERVER_IMAGE_NAME) -f $(SERVER_DOCKERFILE_PATH) .

build-client:
	@echo "Building the client Docker image..."
	docker build -t $(CLIENT_IMAGE_NAME) -f $(CLIENT_DOCKERFILE_PATH) .

# Run Docker containers
run: run-server run-client

run-server:
	@echo "Running the server Docker container..."
	docker run -d --name $(SERVER_CONTAINER_NAME) -p $(SERVER_PORT):$(SERVER_PORT) $(SERVER_IMAGE_NAME)

run-client:
	@echo "Running the client Docker container..."
	docker run --rm --name $(CLIENT_CONTAINER_NAME) --network="host" $(CLIENT_IMAGE_NAME)

# Stop Docker containers
stop: stop-server stop-client

stop-server:
	@echo "Stopping the server Docker container..."
	docker stop $(SERVER_CONTAINER_NAME) || true
	docker rm $(SERVER_CONTAINER_NAME) || true

stop-client:
	@echo "Stopping the client Docker container..."
	# Since the client runs with --rm, no explicit stop is required.

# Clean up Docker images and containers
clean: stop
	@echo "Removing Docker images..."
	docker rmi $(SERVER_IMAGE_NAME) || true
	docker rmi $(CLIENT_IMAGE_NAME) || true

# Print help
help:
	@echo "Usage:"
	@echo "  make build         - Build both server and client Docker images"
	@echo "  make build-server  - Build the server Docker image"
	@echo "  make build-client  - Build the client Docker image"
	@echo "  make run           - Run both server and client Docker containers"
	@echo "  make run-server    - Run the server Docker container"
	@echo "  make run-client    - Run the client Docker container"
	@echo "  make stop          - Stop both server and client containers"
	@echo "  make stop-server   - Stop the server container"
	@echo "  make stop-client   - Stop the client container"
	@echo "  make clean         - Remove all Docker images and containers"
	@echo "  make all           - Build and run both server and client"
	@echo "  make help          - Print this help message"
