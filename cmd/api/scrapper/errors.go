package scrapper

import "errors"

var (
	FailedToGetSessionData = errors.New("failed to get session data")
	FailedToLoadSearchPage = errors.New("failed to load search page")
)
