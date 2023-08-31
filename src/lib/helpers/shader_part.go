package helpers

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderPart struct {
	Id uint32
}

type ShaderPartCreationOpts struct {
	ShaderType uint32
	Src        string
}

type ShaderCreationConfigOverride func(*ShaderPartCreationOpts)

func WithSrc(src string) ShaderCreationConfigOverride {
	return func(opts *ShaderPartCreationOpts) {
		// string have to be null terminated
		// if !strings.HasSuffix(src, "\x00") {
		// 	src += "\x00"
		// }
		opts.Src = src
	}
}

func WithShaderPartType(t uint32) ShaderCreationConfigOverride {
	return func(opts *ShaderPartCreationOpts) {
		opts.ShaderType = t
	}
}

func NewShaderPartForLesson(shaderName string, lesson int, overrides ...ShaderCreationConfigOverride) (*ShaderPart, error) {
	lessonFolder := fmt.Sprintf("lesson%d", lesson)
	p := path.Join("assets", lessonFolder, shaderName)
	file, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	src := bytes.NewBuffer(file).String()
	args := make([]ShaderCreationConfigOverride, 0)
	args = append(args, WithSrc(src))
	args = append(args, overrides...)
	return NewShaderPart(args...)
}

func NewShaderPart(overrides ...ShaderCreationConfigOverride) (shader *ShaderPart, err error) {
	opts := ShaderPartCreationOpts{}
	for _, override := range overrides {
		override(&opts)
	}
	if opts.ShaderType == 0 {
		return nil, fmt.Errorf("missing shader type")
	}

	if !strings.HasSuffix(opts.Src, "\x00") {
		opts.Src += "\x00"
	}
	shaderId, err := compileShaderPart(opts.Src, opts.ShaderType)

	// gl.ShaderSource(shaderId, 1, srcStringData, nil)
	// gl.CompileShader(shaderId)
	// err = checkShader(shaderId)
	if err != nil {
		return nil, err
	}

	return &ShaderPart{
		Id: shaderId,
	}, nil
}

func (shader *ShaderPart) Dispose() {
	// Free source from memory
	gl.DeleteShader(shader.Id)
}

func compileShaderPart(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	err := CheckGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Shader compilation failed")
	if err != nil {
		return 0, err
	}

	return shader, nil
}
