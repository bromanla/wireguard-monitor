package common

import (
	"log"
	"os"
)

// GetEnv retrieves an environment variable or returns a fallback value.
// If neither is set, it logs a fatal error and exits.
func GetEnv(env string, fallback ...string) string {
	// Try to look up environment variable
	if val, ok := os.LookupEnv(env); ok {
		return val
	}

	// Return first fallback if provided
	if len(fallback) > 0 {
		return fallback[0]
	}

	// No value or fallback: fatal error
	log.Fatalf("required environment variable %q is not set", env)

	// Unreachable, but satisfy compiler
	return ""
}
