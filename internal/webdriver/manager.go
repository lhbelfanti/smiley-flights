package webdriver

import (
	"context"

	"smiley-flights/internal/log"
	"smiley-flights/internal/setup"
)

// NewManager initializes a new WebDriver with all its elements
type NewManager func(ctx context.Context) Manager

// MakeNewManager creates a new NewManager
func MakeNewManager() NewManager {
	return func(ctx context.Context) Manager {
		log.Info(ctx, "Initializing WebDriver...")
		var manager Manager = &LocalManager{}
		setup.Must(manager.InitWebDriverService(ctx))
		setup.Must(manager.InitWebDriver(ctx))
		log.Info(ctx, "WebDriver initialized!")

		return manager
	}
}
