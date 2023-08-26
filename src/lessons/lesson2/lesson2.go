package lesson2

import (
	"fmt"
	"learn_opengl/src/lib/helpers"
	"os"
	"path"
)

func readShaderSrc(name string) string {
	p := path.Join("assets", "lesson2", name)
	file, err := os.ReadFile(p)
	helpers.FinishOnError(err)
	return string(file)
}

func Lesson2() {
	// verticies := [9]uint32{}
	// verticiesPointer := unsafe.Pointer(&verticies)

	vertexShader := readShaderSrc("v.vertex")
	fragmentShader := readShaderSrc("f.frag")

	fmt.Println("vertex", vertexShader)
	fmt.Println("fragment", fragmentShader)
	// var vbo uint32
	// gl.GenBuffers(1, &vbo)
	// gl.BufferData(gl.ARRAY_BUFFER, len(verticies), verticiesPointer, gl.STATIC_DRAW)
}

// func init() { runtime.LockOSThread() }
