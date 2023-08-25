package helpers

import "github.com/go-gl/glfw/v3.3/glfw"

func InitGLFW() {
	FinishOnError(glfw.Init())
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
}

func DestroyGLFW() {
	glfw.Terminate()
}
