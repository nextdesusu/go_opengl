package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson4"
	"learn_opengl/src/lib/helpers"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Start")

	helpers.InitGLFW()
	defer helpers.DestroyGLFW()
	lesson4.Lesson4()

	fmt.Println("Finish")
}
