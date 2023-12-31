package helpers

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	Id uint32
}

func NewProgramForLesson(vertexName, fragmentName string, lesson int) (*Program, error) {
	vertexShader, err := NewShaderPartForLesson(vertexName, lesson, WithShaderPartType(gl.VERTEX_SHADER))
	if err != nil {
		return nil, err
	}
	defer vertexShader.Dispose()

	fragmentShader, err := NewShaderPartForLesson(fragmentName, lesson, WithShaderPartType(gl.FRAGMENT_SHADER))
	if err != nil {
		return nil, err
	}
	defer fragmentShader.Dispose()

	program, err := NewProgram(vertexShader, fragmentShader)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func NewProgram(vertexShader, fragmentShader *ShaderPart) (*Program, error) {
	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader.Id)
	gl.AttachShader(program, fragmentShader.Id)
	gl.LinkProgram(program)

	err := CheckGlError(program, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Program linking failed")
	if err != nil {
		return nil, err
	}
	return &Program{
		Id: program,
	}, nil
}

func (program *Program) Use() {
	gl.UseProgram(program.Id)
}

func (program *Program) Delete() {
	gl.DeleteProgram(program.Id)
}

func (program *Program) GetUniformLocation(uniformName string) int32 {
	uniform := uniformName
	if !IsNullTerminated(uniform) {
		uniform += "\x00"
	}

	uniformPtr := gl.Str(uniform)
	val := gl.GetUniformLocation(program.Id, uniformPtr)
	if val == -1 {
		panic(fmt.Errorf("not found uniform location for: %s", uniformName))
	}
	return val
}
