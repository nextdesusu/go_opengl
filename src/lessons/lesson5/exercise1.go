package lesson5

import (
	"learn_opengl/src/lib/helpers"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/mat32"
)

func exercise1() {
	window := createWindow()
	shader, err := helpers.NewShaderForLesson(5, "v.vert", "f.frag")
	helpers.FinishOnError(err)
	vertices := []float32{
		0.5, 0.5, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, // bottom let
		-0.5, 0.5, 0.0, 0.0, 1.0, // top let
	}

	indices := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	const FLOAT_32_SIZE = 4
	VAO, VBO, _ := createTriangleBuffersWithIndices(vertices, indices, func() {
		// position
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*FLOAT_32_SIZE, 0)
		gl.EnableVertexAttribArray(0)

		// texture coords
		gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*FLOAT_32_SIZE, 3*FLOAT_32_SIZE)
		gl.EnableVertexAttribArray(1)
	})

	containerTexture, err := helpers.NewTextureForLesson(5, "container.jpg")
	helpers.FinishOnError(err)
	faceTexture, err := helpers.NewTextureForLesson(5, "awesomeface.png")
	helpers.FinishOnError(err)

	containerTextureId := helpers.NewCstring("texture1")
	faceTextureId := helpers.NewCstring("texture2")
	transformId := helpers.NewCstring("transform")

	shader.Use()
	shader.SetInt(containerTextureId, 0)
	shader.SetInt(faceTextureId, 1)

	gl.ClearColor(0.2, 0.3, 0.3, 1)

	transform := mat32.NewMat4()
	for !window.ShouldClose() {
		processInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		containerTexture.Bind(gl.TEXTURE0)
		faceTexture.Bind(gl.TEXTURE1)

		transform.SetIdentity()
		transform.SetRotationZ(float32(glfw.GetTime()))
		transform.SetPos(mat32.Vec3{X: 0.5, Y: -0.5, Z: 0})

		shader.Use()
		shader.SetMat4(transformId, transform)

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
}
