package square

import (
	"agedito/udemy/modernOpenGL/internal/domain/vertices"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Square struct {
	vao vertices.VAO
	vbo vertices.VBO
}

func (_square *Square) Create() {
	_square.vbo.Create(_vertices)
	_square.vao.Create()

	_square.vbo.Activate()
	_square.vao.Activate()

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func (_square *Square) Draw() {
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

func (_square *Square) Activate() {
	_square.vao.Activate()
}
