package tcp

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github/wordofwisdom/pkg"
	"github/wordofwisdom/server/pow"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today.",
	"Do not wait to strike till the iron is hot, but make it hot by striking.",
	"Great minds discuss ideas; average minds discuss events; small minds discuss people.",
}

// StartServer starts the TCP server on the specified address.
func StartServer(ctx context.Context, address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer shutdownServer(listener)

	fmt.Printf("Server is running on %s...\n", address)

	// Channel to signal server stop
	done := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("Context canceled. Shutting down listener...")
		listener.Close()
		close(done)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// Check if the listener is closed and suppress specific errors
			if isClosedConnectionError(err) {
				fmt.Println("Listener closed. Stopping server...")
				return nil
			}

			select {
			case <-done:
				// Exit loop if the context is canceled
				fmt.Println("Server stopped accepting new connections.")
				return nil
			default:
				// Log errors for unexpected issues
				pkg.LogError("unexpected error on listener", err)
				continue
			}
		}

		// Handle each connection in a separate goroutine
		go handleConnectionWithContext(ctx, conn)
	}
}

// isClosedConnectionError checks if the error is due to a closed listener.
func isClosedConnectionError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, net.ErrClosed)
}

func shutdownServer(listener net.Listener) {
	fmt.Println("Shutting down server...")
	listener.Close()
}

func handleConnectionWithContext(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	if pkg.IsContextDone(ctx) {
		fmt.Println("Context canceled. Stopping connection handler.")
		return
	}

	if err := safeHandleConnection(conn); err != nil {
		fmt.Fprintf(os.Stderr, "Error handling connection: %v\n", err)
	}
}

func handleConnection(conn net.Conn, generateChallenge func() string) {
	// Generate a POW challenge
	challenge := generateChallenge()
	_, write := conn.Write([]byte(fmt.Sprintf("%s\n", challenge)))
	if write != nil {
		return
	}
	s := pow.SolveChallenge(challenge)
	fmt.Println("challenge:", challenge)
	fmt.Println("Solution:", s, pow.ValidateSolution(challenge, s))

	// Read the client's solution
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	solution := strings.TrimSpace(string(buffer[:n]))
	if !pow.ValidateSolution(challenge, solution) {
		_, err := conn.Write([]byte("Invalid solution. Connection closed.\n"))
		if err != nil {
			return
		}
		return
	}

	// Send a random quote upon successful POW validation
	quote := quotes[rand.Intn(len(quotes))]
	_, err = conn.Write([]byte(fmt.Sprintf("Quote: %s\n", quote)))
	if err != nil {
		return
	}
}

// safeHandleConnection wraps handleConnection and handles potential errors
func safeHandleConnection(conn net.Conn) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "[%s] Recovered from panic in connection handler: %v\n", time.Now().Format(time.RFC3339), r)
			if conn != nil {
				conn.Close()
			}
		}
	}()

	if conn == nil {
		err := errors.New("nil connection received")
		fmt.Fprintf(os.Stderr, "[%s] Error: %v\n", time.Now().Format(time.RFC3339), err)
		return err
	}
	// Attempting to handle connection with error capturing
	handleConnection(conn, pow.GenerateChallenge)

	return nil
}
