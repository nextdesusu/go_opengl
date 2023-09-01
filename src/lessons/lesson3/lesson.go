package lesson3

import (
	"learn_opengl/src/lib/helpers"
	"math"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func uniform() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderPartForLesson("v.vert", 3, helpers.WithShaderPartType(gl.VERTEX_SHADER))
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderPartForLesson("f.frag", 3, helpers.WithShaderPartType(gl.FRAGMENT_SHADER))
	helpers.FinishOnError(err)

	program, err := helpers.NewProgram(vertexShader, fragmentShader)
	helpers.FinishOnError(err)

	vertexShader.Dispose()
	fragmentShader.Dispose()

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
		greenValue := (math.Sin(time) / 2) + 0.5

		program.Use()
		u_ourColor := program.GetUniformLocation("ourColor")
		gl.Uniform4f(u_ourColor, 0.0, float32(greenValue), 0.0, 1)
		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteBuffers(1, &EBO)

	program.Delete()
}

func attributes() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderPartForLesson("v_attrib.vert", 3, helpers.WithShaderPartType(gl.VERTEX_SHADER))
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderPartForLesson("f_attrib.frag", 3, helpers.WithShaderPartType(gl.FRAGMENT_SHADER))
	helpers.FinishOnError(err)

	program, err := helpers.NewProgram(vertexShader, fragmentShader)
	helpers.FinishOnError(err)

	vertexShader.Dispose()
	fragmentShader.Dispose()

	vertices := []float32{
		// positions         // colors
		0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom left
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // top
	}

	VAO, VBO := createTriangleBuffers(vertices)

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
		gl.EnableVertexAttribArray(0)

		gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
		gl.EnableVertexAttribArray(1)

		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)

	program.Delete()
}

func lesson3() {
	exercise2()
}
