package pkg

import "context"

// IsContextDone checks if context is done and returns a boolean
func IsContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
