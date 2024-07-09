package huggingface

import (
	"errors"
	"fmt"
)

var (
	ErrExecutingRequest     = errors.New("falied to execute request")
	ErrUnmarshalingResponse = errors.New("failed to unmarshal response")
	ErrEncodingRequestBody  = errors.New("failed to encode request body")
	ErrCreatingRequest      = errors.New("failed to create new request")
	ErrReadingResponseBody  = errors.New("failed to read response body")
)

type HTTPError struct {
	StatusCode int
	Body       *string
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s - %s", e.StatusCode, e.Message, *e.Body)
}
