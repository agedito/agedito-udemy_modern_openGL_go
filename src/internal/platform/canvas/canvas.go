package canvas

import (
	"agedito/udemy/modernOpenGL/internal/domain/vertexs"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var (
	vertices = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}

	vaoId uint32
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

	_vbo := vertexs.VBO{}
	_vbo.Create(vertices)

	gl.GenVertexArrays(1, &vaoId)
	gl.BindVertexArray(vaoId)
	gl.EnableVertexAttribArray(0)

	_vbo.Activate()
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return nil
}

func (_canvas *Canvas) GetGlVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func (_canvas *Canvas) Draw() {
	gl.ClearColor(_canvas.config.ClearColor.Red, _canvas.config.ClearColor.Green, _canvas.config.ClearColor.Blue, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
}

func (_canvas *Canvas) Terminate() {
}
