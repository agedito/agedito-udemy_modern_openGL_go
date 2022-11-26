package canvas

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Canvas struct {
	config Config
}

func (_canvas *Canvas) Create(_config Config) error {
	_canvas.config = _config

	_openGlInitErr := gl.Init()
	if _openGlInitErr != nil {
		return _openGlInitErr
	}

	return nil
}

func (_canvas *Canvas) GetGlVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func (_canvas *Canvas) Draw() {
	gl.ClearColor(_canvas.config.ClearColor.Red, _canvas.config.ClearColor.Green, _canvas.config.ClearColor.Blue, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (_canvas *Canvas) Terminate() {
}
