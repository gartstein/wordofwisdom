package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close connection: %v\n", err)
		}
	}()

	scanner := bufio.NewReader(conn)
	challenge, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read challenge: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Received challenge: %s\n", challenge)

	// Generate a solution for the challenge
	solution := generateSolution(strings.TrimSpace(challenge))
	fmt.Printf("Generated solution: %s\n", solution)

	// Send the solution to the server
	_, err = fmt.Fprintf(conn, "%s\n", solution)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send solution to server: %v\n", err)
		os.Exit(1)
	}

	// Read the response from the server
	resp, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read response from server: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Server response: %s\n", resp)
}

// generateSolution generates a solution for the provided challenge.
func generateSolution(challenge string) string {
	var solution string
	for i := 0; ; i++ {
		solution = strconv.Itoa(i)
		hash := sha256.Sum256([]byte(challenge + solution))
		hashHex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashHex, "0000") {
			return solution
		}
	}
}
