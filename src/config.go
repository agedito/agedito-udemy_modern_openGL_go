package main

import (
	"agedito/udemy/modernOpenGL/internal/application/config"
	"agedito/udemy/modernOpenGL/internal/domain/color"
	"agedito/udemy/modernOpenGL/internal/platform/window"
)

var _windowConfig = window.Config{
	Title:  "My first window",
	Width:  800,
	Height: 800,
}

var _canvasConfig = config.CanvasConfig{
	ClearColor: color.Color{Red: 0.1, Green: 0.2, Blue: 0.3},
}
