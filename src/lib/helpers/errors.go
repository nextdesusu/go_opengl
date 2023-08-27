package helpers

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type getGlParam func(uint32, uint32, *int32)
type getInfoLog func(uint32, int32, *int32, *uint8)

func CheckGlError(glObject uint32, errorParam uint32, getParamFn getGlParam,
	getInfoLogFn getInfoLog, failMsg string) error {

	var success int32
	getParamFn(glObject, errorParam, &success)
	if success == gl.FALSE {
		var infoLog [512 * 10]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		return fmt.Errorf(failMsg + "\n" + string(infoLog[:512*10]))
	}

	return nil
}
