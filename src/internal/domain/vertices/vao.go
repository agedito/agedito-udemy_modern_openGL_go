package vertices

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type VAO struct {
	id uint32
}

func (_vao *VAO) Create() {
	gl.GenVertexArrays(1, &_vao.id)
}

func (_vao *VAO) Activate() {
	gl.BindVertexArray(_vao.id)
}

func (_vao *VAO) Deactivate() {
}
