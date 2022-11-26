package bootstrap

import (
	"agedito/udemy/modernOpenGL/internal/application/config"
	"agedito/udemy/modernOpenGL/internal/platform/canvas"
	"agedito/udemy/modernOpenGL/internal/platform/window"
)

func Run(_windowConfig config.WindowConfig, _canvasConfig config.CanvasConfig) {
	_window := window.Window{}
	_canvas := canvas.Canvas{}

	_ = _window.Create(_windowConfig.Title, _windowConfig.Width, _windowConfig.Height)
	_ = _canvas.Create(_canvasConfig)

	for _window.IsRunning() {
		_canvas.Draw()
		_window.Swap()
	}

	_canvas.Terminate()
	_window.Terminate()
}
