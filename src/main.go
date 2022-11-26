package main

import (
	"agedito/udemy/modernOpenGL/internal/platform/canvas"
	"agedito/udemy/modernOpenGL/internal/platform/window"
)

func main() {
	_window := window.Window{}
	_canvas := canvas.Canvas{}

	_ = _window.Create("My first window", 800, 800)
	_ = _canvas.Create()

	for _window.IsRunning() {
		_canvas.Draw()
		_window.Swap()
	}

	_canvas.Terminate()
	_window.Terminate()
}
