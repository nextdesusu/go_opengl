package helpers

import (
	"fmt"

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

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := createStringBuffer(int(logLength))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	return &Program{
		Id: program,
	}, nil
}

func (program *Program) Use() {
	gl.UseProgram(program.Id)
}
