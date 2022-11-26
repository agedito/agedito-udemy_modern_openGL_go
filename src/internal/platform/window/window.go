package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	title  string
	window *glfw.Window
}

func (_window *Window) Create(_title string, _width int, _height int) error {
	_window.title = _title

	initErr := glfw.Init()
	if initErr != nil {
		return initErr
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	_glfwWindow, creationErr := glfw.CreateWindow(_width, _height, _title, nil, nil)
	if creationErr != nil {
		return creationErr
	}
	_glfwWindow.MakeContextCurrent()
	_window.window = _glfwWindow

	return nil
}

func (_window *Window) IsRunning() bool {
	return !_window.window.ShouldClose()
}

func (_window *Window) Terminate() {
	glfw.Terminate()
}

func (_window *Window) Swap() {
	glfw.PollEvents()
	_window.window.SwapBuffers()
}
