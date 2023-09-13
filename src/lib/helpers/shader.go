package helpers

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/goki/mat32"
)

type Shader struct {
	ID uint32
}

func NewShaderForLesson(lesson int, vp, fp string) (*Shader, error) {
	lessonFolder := fmt.Sprintf("lesson%d", lesson)
	vpPath := path.Join("assets", lessonFolder, vp)
	fpPath := path.Join("assets", lessonFolder, fp)

	return NewShader(vpPath, fpPath)
}

func NewShader(vertexPath, fragmentPath string) (*Shader, error) {
	vertexSrc, err := readShaderPath(vertexPath)
	if err != nil {
		return nil, err
	}
	fragmentSrc, err := readShaderPath(fragmentPath)
	if err != nil {
		return nil, err
	}

	vertexShaderPart, err := compileShaderPart(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	fragmentShaderPart, err := compileShaderPart(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	go gl.DeleteShader(vertexShaderPart)
	go gl.DeleteShader(fragmentShaderPart)

	ID := gl.CreateProgram()
	gl.AttachShader(ID, vertexShaderPart)
	gl.AttachShader(ID, fragmentShaderPart)
	gl.LinkProgram(ID)
	err = CheckGlError(ID, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Program linking failed")
	if err != nil {
		return nil, err
	}

	return &Shader{
		ID: ID,
	}, nil
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.ID)
}

func (shader *Shader) SetBool(name CString, value bool) {
	flag := gl.FALSE
	if value {
		flag = gl.TRUE
	}
	gl.Uniform1i(shader.getUniformLocation(name), int32(flag))
}

func (shader *Shader) SetInt(name CString, value int32) {
	gl.Uniform1i(shader.getUniformLocation(name), value)
}

func (shader *Shader) SetUniform4f(name CString, v0 float32, v1 float32, v2 float32, v3 float32) {
	gl.Uniform4f(shader.getUniformLocation(name), v0, v1, v2, v3)
}

func (shader *Shader) SetFloat(name CString, value float32) {
	gl.Uniform1f(shader.getUniformLocation(name), value)
}

func (shader *Shader) SetMat4(name CString, value *mat32.Mat4) {
	gl.UniformMatrix4fv(shader.getUniformLocation(name), 1, false, &value[0])
}

func (shader *Shader) getUniformLocation(name CString) int32 {
	return gl.GetUniformLocation(shader.ID, name.Ptr)
}

func readShaderPath(path string) (string, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	res := bytes.NewBuffer(src).String()
	return EnsureNullTerminated(res), nil
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
