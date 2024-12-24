package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChallenge(t *testing.T) {
	challenge := GenerateChallenge()
	assert.NotEmpty(t, challenge, "Expected a non-empty challenge")
}

func TestValidateSolution(t *testing.T) {
	challenge := "challenge-12345"
	validSolution := SolveChallenge(challenge)
	invalidSolution := "wrong"
	assert.True(t, ValidateSolution(challenge, validSolution), "Expected valid solution for challenge")
	assert.False(t, ValidateSolution(challenge, invalidSolution), "Expected invalid solution for challenge")
}
