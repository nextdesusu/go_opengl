package lesson4

import (
	"learn_opengl/src/lib/helpers"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type controls struct {
	mixValue float32
}

func exercise4() {
	window := createWindow()
	shader, err := helpers.NewShaderForLesson(4, "v.vert", "f_uniform.frag")
	helpers.FinishOnError(err)
	vertices := []float32{
		// positions          // colors           // texture coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // bottom let
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // top let
	}

	indices := []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}

	const FLOAT_32_SIZE = 4
	VAO, VBO, _ := createTriangleBuffersWithIndices(vertices, indices, func() {
		// position
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 8*FLOAT_32_SIZE, 0)
		gl.EnableVertexAttribArray(0)

		// color
		gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 8*FLOAT_32_SIZE, 3*FLOAT_32_SIZE)
		gl.EnableVertexAttribArray(1)

		// texture
		gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 8*FLOAT_32_SIZE, 6*FLOAT_32_SIZE)
		gl.EnableVertexAttribArray(2)
	})

	containerTexture, err := helpers.NewTextureForLesson(4, "container.jpg")
	helpers.FinishOnError(err)
	faceTexture, err := helpers.NewTextureForLesson(4, "awesomeface.png")
	helpers.FinishOnError(err)

	containerTextureId := helpers.NewCstring("texture1")
	faceTextureId := helpers.NewCstring("texture2")
	mixValue := helpers.NewCstring("mixValue")

	shader.Use()
	shader.SetInt(containerTextureId, 0)
	shader.SetInt(faceTextureId, 1)

	c := &controls{
		mixValue: 0.5,
	}

	localProcessInput := func(w *glfw.Window) {
		if w.GetKey(glfw.KeyEscape) == glfw.Press {
			w.SetShouldClose(true)
		}
		const step float32 = 0.0001
		if w.GetKey(glfw.KeyUp) == glfw.Press {
			c.mixValue += step
			if c.mixValue >= 1.0 {
				c.mixValue = 1.0
			}
		}
		if w.GetKey(glfw.KeyDown) == glfw.Press {
			c.mixValue -= step
			if c.mixValue <= 0.0 {
				c.mixValue = 0.0
			}
		}
	}

	gl.ClearColor(0.2, 0.3, 0.3, 1)
	for !window.ShouldClose() {
		localProcessInput(window)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		containerTexture.Bind(gl.TEXTURE0)
		faceTexture.Bind(gl.TEXTURE1)

		shader.Use()
		shader.SetFloat(mixValue, c.mixValue)

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
}
