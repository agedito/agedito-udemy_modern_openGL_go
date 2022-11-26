package canvas

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"unsafe"
)

var (
	vertices = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}

	vboId uint32
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

	gl.GenBuffers(1, &vboId)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboId)
	fmt.Println(len(vertices), unsafe.Sizeof(vertices), unsafe.Sizeof(vertices[0]))
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW) // 4 is because float32 is 4 bytes long
	gl.GenVertexArrays(1, &vaoId)
	gl.BindVertexArray(vaoId)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboId)
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
