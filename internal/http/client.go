package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"smiley-flights/internal/log"
)

type (
	// Client is an abstraction of the CustomClient methods
	Client interface {
		// NewRequest creates a new HTTP Requests and executes it
		NewRequest(ctx context.Context, method, url string, body any, headers map[string]string) (Response, error)
	}

	// CustomClient represent a custom http.Client wrapper
	CustomClient struct {
		HTTPClient *http.Client
	}

	// Response represent the necessary data of the request response
	Response struct {
		Body   string
		Status string
	}
)

const maxRetries int = 3

// NewClient create a new CustomClient
func NewClient() *CustomClient {
	jar, _ := cookiejar.New(nil)
	return &CustomClient{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Jar:     jar,
		},
	}
}

// NewRequest creates a new HTTP Requests and executes it
func (c *CustomClient) NewRequest(ctx context.Context, method, url string, body any, headers map[string]string) (Response, error) {
	var jsonData []byte
	var err error

	if body != nil {
		// Marshal body to JSON
		switch v := body.(type) {
		case []byte:
			jsonData = v
		default:
			jsonData, err = json.Marshal(body)
			if err != nil {
				log.Error(ctx, err.Error())
				return Response{}, FailedToMarshalBody
			}
		}
	}

	for attempt := range maxRetries {
		// Create new HTTP request
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Error(ctx, err.Error())
			return Response{}, FailedToCreateRequest
		}

		// Set default headers
		req.Header.Set("Content-Type", "application/json")

		// Set custom headers
		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := c.HTTPClient.Do(req)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// Successful response: returning the data
			defer func(body io.ReadCloser) {
				err = body.Close()
				if err != nil {
					log.Error(ctx, err.Error())
				}
			}(resp.Body)

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Error(ctx, err.Error())
				return Response{}, FailedToReadResponse
			}

			return Response{
				Body:   string(respBody),
				Status: resp.Status,
			}, nil
		}

		errMsg := "nil"
		if err != nil {
			errMsg = err.Error()
		} else if resp != nil {
			respBody, _ := io.ReadAll(resp.Body)
			errMsg = fmt.Sprintf("status %d | body: %s", resp.StatusCode, string(respBody))
			resp.Body.Close()
		}

		log.Error(ctx, fmt.Sprintf("Request to %s failed with error %s. Retrying... (Attempt %d/%d)", url, errMsg, attempt+1, maxRetries))

		// Delay before retrying
		time.Sleep(500 * time.Millisecond)
	}

	log.Error(ctx, FailedToExecuteRequest.Error())
	return Response{}, FailedToExecuteRequest
}
