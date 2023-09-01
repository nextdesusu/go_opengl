package lesson3

import (
	"learn_opengl/src/lib/helpers"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func exercise3() {
	window := createWindow()
	vertexShader, err := NewShaderForLesson(3, "v_exercise3.vert", "f_exercise3.frag")
	helpers.FinishOnError(err)

	vertices := []float32{
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		0.0, 0.5, 0.0, // top
	}

	VAO, VBO := createTriangleBuffers(vertices)

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		vertexShader.Use()
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
		gl.EnableVertexAttribArray(0)

		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
}
