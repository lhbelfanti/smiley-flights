package flights

import "errors"

var (
	FailedToParseDepartureDate = errors.New("failed to parse departure date")
	FailedToParseReturnDate    = errors.New("failed to parse return date")
)
