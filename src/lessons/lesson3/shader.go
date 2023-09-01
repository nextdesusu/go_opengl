package lesson3

import (
	"bytes"
	"fmt"
	"learn_opengl/src/lib/helpers"
	"os"
	"path"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	ID               uint32
	locationKeysPool []string
	locationKeys     map[string]*uint8
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
	err = helpers.CheckGlError(ID, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Program linking failed")
	if err != nil {
		return nil, err
	}

	return &Shader{
		ID:           ID,
		locationKeys: make(map[string]*uint8),
	}, nil
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.ID)
}

func (shader *Shader) SetBool(name string, value bool) {
	flag := gl.FALSE
	if value {
		flag = gl.TRUE
	}
	gl.Uniform1i(shader.getUniformLocation(name), int32(flag))
}

func (shader *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(shader.getUniformLocation(name), value)
}

func (shader *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(shader.getUniformLocation(name), value)
}

func (shader *Shader) getUniformLocation(name string) int32 {
	return gl.GetUniformLocation(shader.ID, shader.createOrGetKey(name))
}

func (shader *Shader) createOrGetKey(key string) *uint8 {
	att, ok := shader.locationKeys[key]
	if ok {
		return att
	}

	handled := helpers.EnsureNullTerminated(key)

	shader.locationKeysPool = append(shader.locationKeysPool, handled)
	ptr := gl.Str(handled)
	shader.locationKeys[key] = ptr
	return ptr
}

func readShaderPath(path string) (string, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	res := bytes.NewBuffer(src).String()
	return helpers.EnsureNullTerminated(res), nil
}

func compileShaderPart(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	err := helpers.CheckGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Shader compilation failed")
	if err != nil {
		return 0, err
	}

	return shader, nil
}
