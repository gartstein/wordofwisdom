# Word of Wisdom TCP Server

This repository contains a test task for a Server Engineer position. The goal is to design and implement a "Word of Wisdom" TCP server. The server provides quotes from a collection after verifying clients through a Proof of Work (PoW) mechanism, protecting the server against DDoS attacks. Below are the details of the implementation.

## Features

1. **Proof of Work (PoW)**:
    - The server uses a PoW challenge-response protocol to authenticate clients and mitigate DDoS attacks.
    - The choice of the PoW algorithm is explained in the implementation documentation.

2. **Word of Wisdom**:
    - Once the client successfully completes the PoW challenge, the server sends back a random quote from the "Word of Wisdom" book or another collection of quotes.

3. **Dockerized Solution**:
    - Dockerfiles are provided for both the server and client to simplify setup and deployment.

## Repository Structure

```
.
├── server/
│   ├── Dockerfile
│   ├── main.go            # Main server implementation
│   ├── pow.go             # Proof of Work implementation
│   ├── quotes.txt         # Collection of quotes
├── client/
│   ├── Dockerfile
│   ├── main.go            # Main client implementation
│   ├── pow_solver.go      # PoW solver implementation
├── README.md              # Documentation
└── go.mod                 # Go module dependencies
```

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) installed on your system.
- [Go 1.18+](https://golang.org/) (if you choose to run the code without Docker).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/word-of-wisdom-server.git
   cd word-of-wisdom-server
   ```

2. Build the Docker images for the server and client:
   ```bash
   docker build -t word-of-wisdom-server ./server
   docker build -t word-of-wisdom-client ./client
   ```

### Usage

#### Run the Server

Start the server using Docker:
```bash
docker run -p 12345:12345 word-of-wisdom-server
```

#### Run the Client

Start the client using Docker:
```bash
docker run --network="host" word-of-wisdom-client
```

Alternatively, you can run the server and client directly using Go:

- **Server**:
  ```bash
  go run server/main.go
  ```

- **Client**:
  ```bash
  go run client/main.go
  ```

### Proof of Work

The PoW mechanism uses a computational puzzle to verify clients. The algorithm:
- Challenges the client to find a string that, when hashed with SHA-256, produces a hash with a specified number of leading zeros.
- Protects the server by ensuring the client performs sufficient computational work.

### Testing

You can test the server and client interaction by running the client to connect to the server. Upon successful PoW verification, the server responds with a random quote.

### Configuration

- **Port**: The server listens on port `12345` by default. This can be configured in `server/main.go`.
- **Quotes**: Add or modify quotes in the `server/quotes.txt` file.
- **PoW Difficulty**: The difficulty level of the PoW challenge can be adjusted in `server/pow.go`.

## Design Choices

- **PoW Algorithm**: SHA-256-based PoW was chosen due to its simplicity, efficiency, and wide adoption. It strikes a balance between computational cost and ease of implementation.
- **Random Quotes**: Quotes are fetched from a predefined text file, allowing for easy customization.
- **Dockerization**: Docker ensures platform independence and simplifies the deployment process.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

