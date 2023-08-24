package main

import (
	"fmt"
	"learn_opengl/src/lib/helpers"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// runtime.LockOSThread()
}

func setViewport(w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}

func main() {
	fmt.Println("b4 all")
	fmt.Print("Start")
	defer helpers.DestroyOpenGL()
	helpers.InitOpenGL()
	var width, height = 800, 600
	window, err := glfw.CreateWindow(int(width), int(height), "Learn opengl go", nil, nil)
	helpers.FinishOnError(err)
	if window == nil {
		helpers.FinishOnError(fmt.Errorf("Window is nil"))
	}

	window.MakeContextCurrent()

	setViewport(width, height)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		setViewport(width, height)
	})

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}

	fmt.Print("Finish")
}
