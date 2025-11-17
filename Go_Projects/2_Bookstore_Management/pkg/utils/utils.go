// Package utils provides utility functions for HTTP request handling
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the JSON body from an HTTP request and unmarshals it into the provided interface.
// It returns an error if reading or unmarshalling fails.
func ParseBody(r *http.Request, x interface{}) error {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // Return error if reading fails
	}

	// Unmarshal the JSON body into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		return err // Return error if unmarshalling fails
	}

	return nil // Return nil if everything succeeds
}