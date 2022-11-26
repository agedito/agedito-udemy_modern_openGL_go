package vertexs

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type VBO struct {
	id uint32
}

func (_vbo *VBO) Create(_vertices []float32) {
	gl.GenBuffers(1, &_vbo.id)
	gl.BindBuffer(gl.ARRAY_BUFFER, _vbo.id)
	// 4 is because float32 is 4 bytes long
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(_vertices), gl.Ptr(_vertices), gl.STATIC_DRAW)
}

func (_vbo *VBO) Activate() {
	gl.BindBuffer(gl.ARRAY_BUFFER, _vbo.id)
}

func (_vbo *VBO) Deactivate() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
