package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API key from the headers of an HTTP req
// Example: Authorization: ApiKey{insert API key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Incorrect auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}

	return vals[1], nil
}
