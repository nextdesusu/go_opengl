package main

import (
	"fmt"
	"learn_opengl/src/lessons/lesson1"
	"learn_opengl/src/lib/helpers"
)

func main() {
	fmt.Print("Start")

	helpers.InitGLFW()
	defer helpers.DestroyGLFW()
	lesson1.Lesson1()

	fmt.Print("Finish")
}
