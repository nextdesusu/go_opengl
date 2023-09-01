package helpers

import "strings"

func IsNullTerminated(s string) bool {
	return strings.HasSuffix(s, "\x00")
}

func EnsureNullTerminated(s string) string {
	if IsNullTerminated(s) {
		return s
	}

	return s + "\x00"
}
