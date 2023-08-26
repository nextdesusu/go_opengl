package helpers

import (
	"fmt"
	"os"
	"path"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	Id uint32
}

var shaderSources map[uint32]string = make(map[uint32]string)

func NewShaderForLesson(shaderName string, lesson int, isVertexShader bool) (*Shader, error) {
	lessonFolder := fmt.Sprintf("lesson%d", lesson)
	p := path.Join("assets", lessonFolder, shaderName)
	file, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	src := string(file)

	return NewShader(src, isVertexShader)
}

func NewShader(src string, isVertexShader bool) (shader *Shader, err error) {
	var shaderId uint32
	if isVertexShader {
		shaderId = gl.CreateShader(gl.VERTEX_SHADER)
	} else {
		shaderId = gl.CreateShader(gl.FRAGMENT_SHADER)
	}

	processedSrc := src

	shaderSources[shaderId] = processedSrc

	handledSrc := gl.Str(processedSrc)

	gl.ShaderSource(shaderId, 1, &handledSrc, nil)
	gl.CompileShader(shaderId)
	err = checkShader(shaderId)
	if err != nil {
		return nil, err
	}

	return &Shader{
		Id: shaderId,
	}, nil
}

func (shader *Shader) Destroy() {
	// Free source from memory
	delete(shaderSources, shader.Id)

}

func checkShader(shaderId uint32) error {
	var success int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &success)

	if success == 0 {
		const buffSize int32 = 512
		var infoLog string
		gl.GetShaderInfoLog(shaderId, buffSize, nil, gl.Str(infoLog))
		return fmt.Errorf("shader compilation failed: %s", infoLog)
	}

	return nil
}
