package core

import (
	"fmt"
	"net/http"
)

// GetHeaderValueString doc ...
func GetHeaderValueString(key string, r *http.Request) (string, error) {
	value := r.Header.Get(key)
	if value == "" {
		return value, fmt.Errorf("The %v key header has not been obtained", key)
	}
	return value, nil
}

type contextKey int

const (
	varsKey contextKey = iota
)

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}
