package lesson4

import (
	"learn_opengl/src/lib/helpers"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func lesson() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderForLesson(4, "v.vert", "f_with_color.frag")
	helpers.FinishOnError(err)

	// textCoords := []float32{
	// 	0.0, 0.0, // lower-let corner
	// 	1.0, 0.0, // lower-right corner
	// 	0.5, 1.0, // top-center corner
	// }

	vertices := []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}

	// texCoords := []float32{
	// 	0.0, 0.0, // lower-let corner
	// 	1.0, 0.0, // lower-right corner
	// 	0.5, 1.0, // top-center corner
	// }

	VAO, VBO := createTriangleBuffers(vertices)

	ourColor := helpers.NewCstring("ourColor")

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		vertexShader.Use()
		vertexShader.SetUniform4f(ourColor, 1.0, 0.0, 0.0, 1)

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)

}
