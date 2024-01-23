package stopthecap

import "strings"

// contains checks if an element exists in a slice
func contains(slice []string, element string) bool {
	for _, item := range slice {
		normalizedItem := strings.ToLower(item)
		if normalizedItem == element {
			return true
		}
	}
	return false
}
