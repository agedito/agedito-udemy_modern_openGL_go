package bootstrap

import (
	"agedito/udemy/modernOpenGL/internal/platform/canvas"
	"agedito/udemy/modernOpenGL/internal/platform/window"
)

func Run(_windowConfig window.Config, _canvasConfig canvas.Config) {
	_window := window.Window{}
	_canvas := canvas.Canvas{}

	_ = _window.Create(_windowConfig)
	_ = _canvas.Create(_canvasConfig)

	for _window.IsRunning() {
		_canvas.Draw()
		_window.Swap()
	}

	_canvas.Terminate()
	_window.Terminate()
}
