package helpers

import "strings"

func IsNullTerminated(s string) bool {
	return strings.HasSuffix(s, "\x00")
}
