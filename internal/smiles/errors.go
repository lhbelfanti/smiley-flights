package smiles

import "errors"

var (
	FailedToUnmarshalResponse = errors.New("failed to unmarshal response body")
	FailedToExecuteRequest    = errors.New("failed to execute request")
)
