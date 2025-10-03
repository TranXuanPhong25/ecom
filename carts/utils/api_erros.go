// utils/api_errors.go
package utils

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("API error (status %d): %s - %v", e.StatusCode, e.Message, e.Err)
	}
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

func NewAPIError(statusCode int, message string, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
