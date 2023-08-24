package jobqueue

import "runtime"

// Arrange that main.main runs on main thread.
func InitJobQueue() {
	runtime.LockOSThread()
}

// Main runs the main SDL service loop.
// The binary's main.main must call sdl.Main() to run this loop.
// Main does not return. If the binary needs to do other work, it
// must do it in separate goroutines.
func ProcessJobs() {
	for f := range mainfunc {
		f()
	}
}

// queue of work to run in main thread.
var mainfunc = make(chan func())

// do runs f on the main thread.
func DoJob(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}
