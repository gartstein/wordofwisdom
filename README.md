# Word of Wisdom TCP Server

This repository contains a test task for a Server Engineer position. The goal is to design and implement a "Word of Wisdom" TCP server. The server provides quotes from a collection after verifying clients through a Proof of Work (PoW) mechanism, protecting the server against DDoS attacks.

## Features

1. **Proof of Work (PoW)**:
    - The server uses a PoW challenge-response protocol to authenticate clients and mitigate DDoS attacks.
    - The PoW algorithm is based on SHA-256, ensuring simplicity and efficiency.

2. **Word of Wisdom**:
    - After successful PoW verification, the server sends back a random quote from the "Word of Wisdom" book or another collection of quotes.

3. **Dockerized Solution**:
    - Dockerfiles are provided for both the server and client to simplify setup and deployment.

4. **Makefile**:
    - A `Makefile` is included to streamline building, running, stopping, and cleaning up server and client Docker containers.

---

## Repository Structure

```
.
├── build/
│   ├── server/             # Dockerfile for the server
│   ├── client/             # Dockerfile for the client
├── client/
│   ├── main.go             # Main client implementation
├── server/
│   ├── main.go             # Main server implementation
│   ├── pow                 # Proof of Work implementation
│   ├── tcp                 # TCP server implementation
├── pkg/
│   ├── mock.go             # Mock connection utilities for testing
│   ├── log.go              # Logging
│   ├── context.go          # Helpers for context
├── Makefile                # Build and run automation
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksum
├── LICENSE                 # License file
└── README.md               # Documentation
```

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) installed on your system.
- [Go 1.18+](https://golang.org/) (if you choose to run the code without Docker).

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/word-of-wisdom-server.git
   cd word-of-wisdom-server
   ```

2. Build the Docker images for the server and client:
   ```bash
   docker build -t wordofwisdom-server ./build/server
   docker build -t wordofwisdom-client ./build/client
   ```

---

### Usage

#### Using the Makefile

The `Makefile` simplifies the build and run process for both the server and client. Use the following commands:

- **Build the Docker images**:
  ```bash
  make build
  ```
  This builds both the server and client images.

- **Run the server and client**:
  ```bash
  make run
  ```
  This starts the server and runs the client, connecting it to the server.

- **Run only the server**:
  ```bash
  make run-server
  ```

- **Run only the client**:
  ```bash
  make run-client
  ```

- **Stop the running containers**:
  ```bash
  make stop
  ```

- **Clean up Docker images and containers**:
  ```bash
  make clean
  ```

#### Run Without Makefile

- **Run the Server**:
  ```bash
  docker run -d --name wordofwisdom-server -p 8080:8080 wordofwisdom-server
  ```

- **Run the Client**:
  ```bash
  docker run --rm --network="host" wordofwisdom-client
  ```

#### Run Directly Using Go

- **Server**:
  ```bash
  go run server/main.go
  ```

- **Client**:
  ```bash
  go run client/main.go
  ```

---

### Proof of Work

The PoW mechanism uses a computational puzzle to verify clients. The algorithm:
- Challenges the client to find a string that, when hashed with SHA-256, produces a hash with a specified number of leading zeros.
- Protects the server by ensuring the client performs sufficient computational work.

---

### Testing

You can test the server and client interaction using the client. Upon successful PoW verification, the server responds with a random quote.

---

### Configuration

- **Port**: The server listens on port `8080` by default. This can be configured in `server/main.go`.
- **Quotes**: Add or modify quotes in the `server/quotes.txt` file.
- **PoW Difficulty**: The difficulty level of the PoW challenge can be adjusted in `server/pow.go`.

---

## Design Choices

- **PoW Algorithm**: SHA-256-based PoW was chosen for its simplicity, efficiency, and wide adoption.
- **Random Quotes**: Quotes are fetched from a predefined text file, allowing for easy customization.
- **Dockerization**: Docker ensures platform independence and simplifies deployment.
- **Makefile**: The inclusion of a `Makefile` streamlines common tasks for development and deployment.

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

