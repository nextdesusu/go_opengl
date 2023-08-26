package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson2"
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
	lesson2.Lesson2()

	fmt.Println("Finish")
}
