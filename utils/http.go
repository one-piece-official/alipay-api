package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeHTTPClientGet(httpClient *http.Client, url string, body io.Reader) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), responseTimeoutMillisecond*time.Millisecond)
	defer cancel()

	return getWithContext(ctx, httpClient, url, body)
}

// getWithContext, Sending an HTTP request and accepting context.
func getWithContext(ctx context.Context, httpClient *http.Client, url string, body io.Reader) (res *http.Response, err error) {
	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, body)
	if err != nil {
		return nil, fmt.Errorf("http client request with context failed: %w", err)
	}

	if res, err = httpClient.Do(req); err != nil {
		return nil, fmt.Errorf("http client do failed: %w", err)
	}

	return res, nil
}
