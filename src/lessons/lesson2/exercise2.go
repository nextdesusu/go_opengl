package lesson2

import (
	"learn_opengl/src/lib/helpers"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func exercise2() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderForLesson("v.vertex", 2, helpers.WithShaderType(gl.VERTEX_SHADER))
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderForLesson("f.frag", 2, helpers.WithShaderType(gl.FRAGMENT_SHADER))
	helpers.FinishOnError(err)

	program, err := helpers.NewProgram(vertexShader, fragmentShader)
	helpers.FinishOnError(err)

	vertexShader.Dispose()
	fragmentShader.Dispose()

	t1 := createTriangle2D(0, 0, 100, 100)
	t2 := createTriangle2D(0, 100, 200, 100)

	VAO1, VBO1 := createTriangleBuffers(t1)
	VAO2, VBO2 := createTriangleBuffers(t2)

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()

		gl.BindVertexArray(VAO1)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(t1)/3))

		gl.BindVertexArray(VAO2)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(t2)/3))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO1)
	gl.DeleteBuffers(1, &VBO1)

	gl.DeleteVertexArrays(1, &VAO2)
	gl.DeleteBuffers(1, &VBO2)

	program.Delete()
}
