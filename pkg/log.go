package pkg

import (
	"fmt"
	"os"
)

// LogError logs connection errors to stderr.
func LogError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
}
