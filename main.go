package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson3"
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
	lesson3.Lesson3()

	fmt.Println("Finish")
}
