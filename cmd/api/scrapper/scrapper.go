package scrapper

import (
	"smiley-flights/cmd/api/page"

	"github.com/tebeka/selenium"
)

type New func(webDriver selenium.WebDriver) Execute

// MakeNew creates a new New
func MakeNew() New {
	return func(webDriver selenium.WebDriver) Execute {
		// Helpers
		loadPage := page.MakeLoad(webDriver)

		return MakeExecute(loadPage)
	}
}
