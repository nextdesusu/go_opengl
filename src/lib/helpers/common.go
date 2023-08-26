package helpers

import "strings"

func createStringBuffer(size int) string {
	return strings.Repeat("\x00", int(size+1))
}
