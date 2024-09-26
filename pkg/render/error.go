package render

import "net/http"

// Error represents a json-encoded API error.
type Error struct {
	Err error `json:"error"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// NewError returns a new json-encoded API error from an error
func NewError(err error) error {
	return NewErrorStr(err.Error())
}

// NewErrorStr returns a new json error message.
func NewErrorStr(message string) error {
	return &Error{Err: errorStr(message)}
}

type errorStr string

func (e errorStr) Error() string {
	return string(e)
}

var (
	// ErrJSONDecode is returned when the json body is invalid.
	ErrJSONDecode = NewErrorStr("error decoding json")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = NewErrorStr(http.StatusText(http.StatusUnauthorized))

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = NewErrorStr(http.StatusText(http.StatusForbidden))

	// ErrInternal is returned when user hits an internal server error.
	ErrInternal = NewErrorStr(http.StatusText(http.StatusInternalServerError))

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = NewErrorStr(http.StatusText(http.StatusNotFound))

	// ErrBadRequest is returned when a request is not valid.
	ErrBadRequest = NewErrorStr(http.StatusText(http.StatusBadRequest))
)
