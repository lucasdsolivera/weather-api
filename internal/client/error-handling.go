package client

import "fmt"

type HTTPError struct {
	StatusCode int
	Err        error
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func NewHTTPError(status int, msg string) *HTTPError {
	return &HTTPError{
		StatusCode: status,
		Err:        fmt.Errorf("%v", msg),
	}
}
