package helpers

import "github.com/go-gl/gl/v3.3-core/gl"

type CString struct {
	Value string
	Ptr   *uint8
}

func NewCstring(src string) CString {
	handled := EnsureNullTerminated(src)

	return CString{
		Value: handled,
		Ptr:   gl.Str(handled),
	}
}
