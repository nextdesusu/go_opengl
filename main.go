package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson1"
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
	lesson1.Lesson1()

	fmt.Println("Finish")
}
