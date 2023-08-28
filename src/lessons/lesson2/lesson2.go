package lesson2

import (
	"fmt"
	"learn_opengl/src/lib/helpers"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func setViewport(w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}

func processInput(w *glfw.Window) {
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		w.SetShouldClose(true)
	}
}

func createWindow() *glfw.Window {
	var width, height = 800, 600
	window, err := glfw.CreateWindow(int(width), int(height), "Learn opengl go", nil, nil)
	helpers.FinishOnError(err)
	if window == nil {
		helpers.FinishOnError(fmt.Errorf("window is nil"))
	}

	window.MakeContextCurrent()

	gl.Init()
	setViewport(width, height)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		setViewport(width, height)
	})

	return window
}

/*
 * Creates the Vertex Array Object for a triangle.
 */
func createTriangleBuffers(vertices []float32) (VAO uint32, VBO uint32) {
	gl.GenVertexArrays(1, &VAO)

	gl.GenBuffers(1, &VBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// specify the format of our vertex input
	// (shader) input 0
	// vertex has size 3
	// vertex items are of type FLOAT
	// do not normalize (already done)
	// stride of 3 * sizeof(float) (separation of vertices)
	// offset of where the position data starts (0 for the beginning)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO, VBO
}

func createTriangleBuffersWithIndices(vertices []float32, indices []uint32) (VAO uint32, VBO uint32, EBO uint32) {
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.GenBuffers(1, &EBO)
	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	const FLOAT_32_SIZE = 4
	const UINT_32_SIZE = 4

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*FLOAT_32_SIZE, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*UINT_32_SIZE, gl.Ptr(indices), gl.STATIC_DRAW)

	// specify the format of our vertex input
	// (shader) input 0
	// vertex has size 3
	// vertex items are of type FLOAT
	// do not normalize (already done)
	// stride of 3 * sizeof(float) (separation of vertices)
	// offset of where the position data starts (0 for the beginning)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*FLOAT_32_SIZE, 0)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO, VBO, EBO
}

func Lesson2() {
	window := createWindow()
	vertexShader, err := helpers.NewShaderForLesson("v.vertex", 2, helpers.WithShaderType(gl.VERTEX_SHADER))
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderForLesson("f.frag", 2, helpers.WithShaderType(gl.FRAGMENT_SHADER))
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

		program.Use()
		// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
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
