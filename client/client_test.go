package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

// Mock connection for testing
type mockConn struct {
	readBuffer  *bytes.Buffer
	writeBuffer *bytes.Buffer
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	return m.readBuffer.Read(b)
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	return m.writeBuffer.Write(b)
}

func (m *mockConn) Close() error {
	return nil
}

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// Mock function for client to allow injection of the server address
func mainWithConnection(serverAddr string) {
	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return
	}
	defer conn.Close()

	// Read the challenge from the server
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		fmt.Printf("Failed to read challenge from server: %v\n", scanner.Err())
		return
	}
	challenge := scanner.Text()
	fmt.Printf("Received challenge: %s\n", challenge)

	// Generate a solution for the challenge
	solution := generateSolution(challenge)
	fmt.Printf("Generated solution: %s\n", solution)

	// Send the solution to the server
	_, err = fmt.Fprintf(conn, "%s\n", solution)
	if err != nil {
		fmt.Printf("Failed to send solution to server: %v\n", err)
		return
	}

	// Read the response from the server
	if !scanner.Scan() {
		fmt.Printf("Failed to read response from server: %v\n", scanner.Err())
		return
	}
	response := scanner.Text()
	fmt.Printf("Server response: %s\n", response)
}

// Test for the client
func TestClient(t *testing.T) {
	// Create a mock server listener
	listener, err := net.Listen("tcp", ":0") // Bind to any available port
	if err != nil {
		t.Fatalf("Failed to create mock server: %v", err)
	}
	defer listener.Close()

	mockServerAddr := listener.Addr().String()

	// WaitGroup to ensure synchronization
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the mock server
	go func() {
		defer wg.Done()
		conn, err := listener.Accept()
		if err != nil {
			t.Fatalf("Failed to accept connection: %v", err)
		}
		defer conn.Close()

		// Mock server sends a challenge
		challenge := "challenge-12345"
		_, _ = conn.Write([]byte(challenge + "\n"))

		// Read the client's solution
		scanner := bufio.NewScanner(conn)
		if !scanner.Scan() {
			t.Fatalf("Failed to read solution from client: %v", scanner.Err())
		}
		solution := scanner.Text()
		fmt.Printf("Received solution from client: %s\n", solution)

		// Validate the solution and send a response
		if ValidateSolution(challenge, solution) {
			_, _ = conn.Write([]byte("Quote: Success!\n"))
		} else {
			_, _ = conn.Write([]byte("Invalid solution. Connection closed.\n"))
		}
	}()

	// Run the client logic
	go func() {
		mainWithConnection(mockServerAddr)
	}()

	// Wait for the client to finish
	wg.Wait()
}

// Dummy implementation of ValidateSolution for the test
func ValidateSolution(challenge, solution string) bool {
	hash := sha256.Sum256([]byte(challenge + solution))
	hashHex := fmt.Sprintf("%x", hash)
	return strings.HasPrefix(hashHex, "0000")
}
