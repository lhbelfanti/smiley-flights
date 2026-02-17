package elements

import (
	"context"
	"fmt"
	"time"

	"smiley-flights/internal/log"
)

// RetrieveAndClickButton retrieves a button element and clicks on it
type RetrieveAndClickButton func(ctx context.Context, by, value, element string, timeout time.Duration) error

// MakeRetrieveAndClickButton creates a new RetrieveAndClickButton
func MakeRetrieveAndClickButton(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndClickButton {
	return func(ctx context.Context, by, value, element string, timeout time.Duration) error {
		button, err := waitAndRetrieveElement(ctx, by, value, timeout)
		if err != nil {
			log.Error(ctx, fmt.Sprintf("failed to retrieve button: %s. error: %v", element, err))
			return FailedToRetrieveButton
		}

		err = button.Click()
		if err != nil {
			log.Error(ctx, fmt.Sprintf("failed to click button: %s. error: %v", element, err))
			return FailedToClickButton
		}

		return nil
	}
}
