package client

import (
	"fmt"
	"strings"
)

// HTTPError is a type that implements the error interface.
// It contains relevant request and response fields for debugging.
type HTTPError struct {
	URL        string
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("request to %s failed with status code %d and body %s", e.URL, e.StatusCode, e.Body)
}

// Is is used for comparing with errors.Is
func (e *HTTPError) Is(target error) bool {
	t, ok := target.(*HTTPError)
	if !ok {
		return false
	}
	return (e.StatusCode == t.StatusCode || t.StatusCode == 0) &&
		(strings.Contains(e.Body, t.Body) || t.Body == "")
}

// FundAccountError is a type that implements the error interface.
// It contains relevant network response code for debugging.
type FundAccountError struct {
	Err       error
	NetworkRC string
}

func (e *FundAccountError) Error() string {
	return fmt.Sprintf("network response code %s: err %v", e.NetworkRC, e.Err)
}

// Unwrap returns the underlying error.
func (e *FundAccountError) Unwrap() error {
	return e.Err
}
