package canvas

import (
	"agedito/udemy/modernOpenGL/internal/domain/shapes/triangle"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var (
	_triangle = triangle.Triangle{}
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

	_triangle.Create()

	return nil
}

func (_canvas *Canvas) GetGlVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func (_canvas *Canvas) Draw() {
	gl.ClearColor(_canvas.config.ClearColor.Red, _canvas.config.ClearColor.Green, _canvas.config.ClearColor.Blue, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	_triangle.Draw()
}

func (_canvas *Canvas) Terminate() {
}
