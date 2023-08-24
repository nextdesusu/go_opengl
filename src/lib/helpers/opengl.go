package helpers

import "github.com/go-gl/glfw/v3.3/glfw"

func InitOpenGL() {
	FinishOnError(glfw.Init())
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
}

func DestroyOpenGL() {
	glfw.Terminate()
}
