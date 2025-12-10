package urlutil

import (
	"net/url"
	"strings"

	"github.com/BeaverHouse/go-common/errorhandle"
)

// NormalizeURL adds https:// prefix if URL doesn't have a protocol scheme,
// and validates the final URL format
func NormalizeURL(inputURL string) (string, error) {
	parsedURL := strings.TrimSpace(inputURL)
	if parsedURL == "" {
		return "", nil
	}

	if !strings.HasPrefix(parsedURL, "http://") && !strings.HasPrefix(parsedURL, "https://") {
		parsedURL = "https://" + parsedURL
	}

	_, err := url.Parse(parsedURL)
	if err != nil {
		return "", errorhandle.ErrValidationFailed("invalid URL format: " + parsedURL)
	}

	return parsedURL, nil
}
