package pkg

import (
	"os"
)

// Use this function to support Railway's dynamic port assignment
func GetPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	// fallback for local development
	return "8080"
}
