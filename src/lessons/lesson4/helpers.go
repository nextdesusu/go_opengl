package lesson4

import (
	"fmt"
	"learn_opengl/src/lib/app_math"
	"learn_opengl/src/lib/helpers"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func checkArg(value, max int, label string) {
	if value < 0 || value > max {
		panic(fmt.Errorf("incorrect %s: should be less than %d or bigger than %d", label, max, 0))
	}
}

const SCREEN_WIDTH, SCREEN_HEIGHT = 800, 600

func setViewport(w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}

func processInput(w *glfw.Window) {
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		w.SetShouldClose(true)
	}
}

func createWindow() *glfw.Window {
	window, err := glfw.CreateWindow(int(SCREEN_WIDTH), int(SCREEN_HEIGHT), "Learn opengl go", nil, nil)
	helpers.FinishOnError(err)
	if window == nil {
		helpers.FinishOnError(fmt.Errorf("window is nil"))
	}

	window.MakeContextCurrent()

	gl.Init()
	setViewport(SCREEN_WIDTH, SCREEN_HEIGHT)

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

func createTriangleBuffersWithIndices(vertices []float32, indices []uint32, bindAttrs func()) (VAO uint32, VBO uint32, EBO uint32) {
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

	bindAttrs()

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO, VBO, EBO
}

const ogl_top float32 = 1
const ogl_bottom float32 = -1
const ogl_left = -1
const ogl_right = 1

type point struct {
	X, Y float32
}

func normalizedPoint(x, y int) point {
	X := app_math.NumberInRange[float32](0, float32(SCREEN_WIDTH), ogl_left, ogl_right, float32(x))
	Y := app_math.NumberInRange[float32](0, float32(SCREEN_HEIGHT), ogl_top, ogl_bottom, float32(y))
	return point{
		X: X,
		Y: Y,
	}
}

func createTriangle2D(x, y, w, h int) []float32 {
	checkArg(x, SCREEN_WIDTH, "x")
	checkArg(y, SCREEN_HEIGHT, "y")
	checkArg(w, SCREEN_WIDTH, "width")
	checkArg(h, SCREEN_HEIGHT, "height")

	tr := normalizedPoint(x, y)
	br := normalizedPoint(x+w, y+h)
	bl := normalizedPoint(x, y+h)

	return []float32{
		// x, y
		tr.X, tr.Y, 0.0, // top right
		br.X, br.Y, 0.0, // bottom right
		bl.X, bl.Y, 0.0, // bottom left
	}
}
