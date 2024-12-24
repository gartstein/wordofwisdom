package tcp

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github/wordofwisdom/pkg"
	"github/wordofwisdom/server/pow"
)

func TestHandleConnection(t *testing.T) {
	mockChallenge := "mockChallenge-12345"
	// Generate a valid solution
	validSolution := pow.SolveChallenge(mockChallenge)

	mockConn := pkg.NewMockConn(validSolution + "\n")
	handleConnection(mockConn, func() string {
		return mockChallenge
	})

	output := make([]byte, 1024)
	n, err := mockConn.Read(output)
	if err != nil && err != io.EOF {
		t.Fatalf("Unexpected error during read: %v", err)
	}
	output = output[:n]
	t.Logf("MockConn Output: %s", mockConn.Output())

	// Validate the output
	assert.Contains(t, mockConn.Output(), "Quote:", "Expected 'Quote:' in the output")
	foundQuote := false
	for _, quote := range quotes { // quotes is the array of available quotes
		if strings.Contains(mockConn.Output(), quote) {
			foundQuote = true
			break
		}
	}
	assert.True(t, foundQuote, "Expected one of the quotes in the output")
}
