package vertices

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type EBO struct {
	id uint32
}

func (_ebo *EBO) Create(_indices []int32) {
	gl.GenBuffers(1, &_ebo.id)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, _ebo.id)
	// 4 is because float32 is 4 bytes long
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(_indices), gl.Ptr(_indices), gl.STATIC_DRAW)
}

func (_ebo *EBO) Activate() {
	//gl.BindBuffer(gl.ARRAY_BUFFER, _ebo.id)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, _ebo.id)
}

func (_ebo *EBO) Deactivate() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
