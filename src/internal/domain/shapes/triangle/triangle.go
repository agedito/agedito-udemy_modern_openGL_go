package triangle

import (
	"agedito/udemy/modernOpenGL/internal/domain/vertices"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Triangle struct {
	vao vertices.VAO
	vbo vertices.VBO
}

func (_triangle *Triangle) Create() {
	_triangle.vbo.Create(_vertices)
	_triangle.vao.Create()

	_triangle.vbo.Activate()
	_triangle.vao.Activate()

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func (_triangle *Triangle) Draw() {
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
