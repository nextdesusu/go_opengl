package helpers

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	Id uint32
}

func NewProgram(vertexShader, fragmentShader *Shader) (*Program, error) {
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
