package errorhandle

import (
	"regexp"
	"strconv"
)

// ExtractHTTPStatusFromError extracts HTTP status code from error message.
// If it fails to extract, it returns 500.
//
//   - Error format: [module prefix][http status code]-[error code]: [error message]
//   - Example: COM401-01: Invalid username or password
//   - Allowed module prefix: AU (Auth Module), COM (Common Module), and optional user input.
func ExtractHTTPStatusFromError(errorMessage string, modulePrefix string) int {
	var re *regexp.Regexp
	if modulePrefix == "" {
		re = regexp.MustCompile(`(AU|COM)(\d{3})-\d{2}:`)
	} else {
		re = regexp.MustCompile(`(AU|COM|` + modulePrefix + `)(\d{3})-\d{2}:`)
	}
	matches := re.FindStringSubmatch(errorMessage)
	if len(matches) >= 3 {
		if code, err := strconv.Atoi(matches[2]); err == nil {
			return code
		}
	}
	return 500
}
