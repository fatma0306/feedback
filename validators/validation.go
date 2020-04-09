package validators

import (
	"strings"
)

// IsEmpty cheks a string is empty or not
func IsEmpty(input string) bool {
	if len(strings.TrimSpace(input)) == 0 {
		return true
	}
	return false
}
