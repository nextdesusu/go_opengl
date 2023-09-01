package lesson3

import (
	"learn_opengl/src/lib/helpers"
	"math"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func exercise2() {
	window := createWindow()
	shader, err := NewShaderForLesson(3, "v_offset.vert", "f.frag")

	helpers.FinishOnError(err)

	vertices := []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}
	indicies := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	VAO, VBO, EBO := createTriangleBuffersWithIndices(vertices, indicies)

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		time := glfw.GetTime()
		moveValue := float32((math.Sin(time) / 2) + 0.5)

		shader.Use()
		shader.SetFloat("offset", moveValue)
		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteBuffers(1, &EBO)
}
