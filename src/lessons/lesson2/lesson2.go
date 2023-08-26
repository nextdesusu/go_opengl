package lesson2

import (
	"fmt"
	"learn_opengl/src/lib/helpers"

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

func createWindow() {
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

	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func Lesson2() {
	// verticies := [9]uint32{}
	// verticiesPointer := unsafe.Pointer(&verticies)

	gl.Init()
	vertexShader, err := helpers.NewShaderForLesson("v.vertex", 2, true)
	helpers.FinishOnError(err)
	fragmentShader, err := helpers.NewShaderForLesson("f.frag", 2, false)
	helpers.FinishOnError(err)

	fmt.Println("vertex", vertexShader)
	fmt.Println("fragment", fragmentShader)

	// var vbo uint32
	// gl.GenBuffers(1, &vbo)
	// gl.BufferData(gl.ARRAY_BUFFER, len(verticies), verticiesPointer, gl.STATIC_DRAW)
}

// func init() { runtime.LockOSThread() }
