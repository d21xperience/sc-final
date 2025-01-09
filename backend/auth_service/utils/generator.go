package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// GenerateUsername generates a unique username based on name and ID.
func GenerateUsername(name string, id int, isUsernameTaken func(string) bool) string {
	// Normalize the name: lowercase, remove special characters, and trim
	name = strings.ToLower(name)
	name = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(name, "")
	name = strings.TrimSpace(name)

	// Fallback if the name becomes empty after normalization
	if len(name) == 0 {
		name = "user"
	}

	// Base username
	username := fmt.Sprintf("%s%d", name, id)

	// Ensure uniqueness
	counter := 0
	for isUsernameTaken(username) {
		counter++
		username = fmt.Sprintf("%s%d", name, id+counter)
	}

	return username
}
