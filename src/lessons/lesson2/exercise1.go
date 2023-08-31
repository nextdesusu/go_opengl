package lesson2

import (
	"learn_opengl/src/lib/helpers"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func exercise1() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderPartForLesson("v.vert", 2, helpers.WithShaderPartType(gl.VERTEX_SHADER))
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderPartForLesson("f.frag", 2, helpers.WithShaderPartType(gl.FRAGMENT_SHADER))
	helpers.FinishOnError(err)

	program, err := helpers.NewProgram(vertexShader, fragmentShader)
	helpers.FinishOnError(err)

	vertexShader.Dispose()
	fragmentShader.Dispose()

	t1 := createTriangle2D(0, 0, 100, 100)
	t2 := createTriangle2D(100, 0, 100, 100)
	var vertices []float32
	vertices = append(vertices, t1...)
	vertices = append(vertices, t2...)

	VAO, VBO := createTriangleBuffers(vertices)

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()

		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)

	program.Delete()
}
