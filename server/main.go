package main

import (
	"context"
	"fmt"
	"github/wordofwisdom/pkg"
	"github/wordofwisdom/server/tcp"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const address = ":8080"

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Create a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal channel to catch termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle termination signals
	go func() {
		sig := <-signalChan
		fmt.Printf("\nReceived signal: %v. Shutting down server...\n", sig)
		cancel()
	}()

	err := tcp.StartServer(ctx, address)
	if err != nil {
		pkg.LogError("server encountered an error: ", err)
		os.Exit(1)
	}
	fmt.Println("Server is stopped")
}
