package http

import "errors"

var (
	FailedToMarshalBody    = errors.New("failed to marshal body to JSON")
	FailedToCreateRequest  = errors.New("failed to create a new HTTP request")
	FailedToReadResponse   = errors.New("failed to read response body")
	FailedToExecuteRequest = errors.New("failed to execute HTTP request after retries")
)
