package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateChallenge creates a random challenge string for POW.
func GenerateChallenge() string {
	return fmt.Sprintf("challenge-%d", rand.Intn(1000000))
}

// ValidateSolution checks if the given solution matches the POW challenge requirements.
func ValidateSolution(challenge, solution string) bool {
	hash := sha256.Sum256([]byte(challenge + solution))
	hashHex := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashHex, "0000")
}

// SolveChallenge generates a solution for a given challenge.
func SolveChallenge(challenge string) string {
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
