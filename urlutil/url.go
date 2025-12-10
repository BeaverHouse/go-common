package urlutil

import (
	"net/url"
	"strings"
)

// NormalizeURL adds https:// prefix if URL doesn't have a protocol scheme,
// and validates the final URL format
func NormalizeURL(inputURL string) (string, error) {
	normalizedURL := strings.TrimSpace(inputURL)
	if normalizedURL == "" {
		return "", nil
	}

	if !strings.HasPrefix(normalizedURL, "http://") && !strings.HasPrefix(normalizedURL, "https://") {
		normalizedURL = "https://" + normalizedURL
	}

	_, err := url.Parse(normalizedURL)
	if err != nil {
		return "", err
	}

	return normalizedURL, nil
}
