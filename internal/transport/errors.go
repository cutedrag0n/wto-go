package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIError represents a non-success HTTP response. Body preserves unrecognized error formats.
type APIError struct {
	StatusCode int
	Header     http.Header
	Body       json.RawMessage
	Type       string                     `json:"type,omitempty"`
	Title      string                     `json:"title,omitempty"`
	Detail     string                     `json:"detail,omitempty"`
	Message    string                     `json:"message,omitempty"`
	Status     int                        `json:"status,omitempty"`
	Instance   string                     `json:"instance,omitempty"`
	Errors     map[string][]string        `json:"errors,omitempty"`
	Extensions map[string]json.RawMessage `json:"extensions,omitempty"`
}

// Error deliberately omits server text, which can reflect credentials or request data.
func (e *APIError) Error() string {
	return fmt.Sprintf("wto: HTTP %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}
