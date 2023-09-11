package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson5"
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
	lesson5.Lesson5()

	fmt.Println("Finish")
}
