package square

import (
	"agedito/udemy/modernOpenGL/internal/domain/vertices"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Square struct {
	vao vertices.VAO
	vbo vertices.VBO
	veo vertices.EBO
}

func (_square *Square) Create() {
	_square.vbo.Create(_vertices)
	_square.vao.Create()
	_square.veo.Create(_indices)

	_square.vbo.Activate()
	_square.vao.Activate()
	_square.veo.Activate()

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func (_square *Square) Draw() {
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
}

func (_square *Square) Activate() {
	_square.vao.Activate()
}
