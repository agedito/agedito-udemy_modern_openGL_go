package main

import (
	"agedito/udemy/modernOpenGL/bootstrap"
	"agedito/udemy/modernOpenGL/internal/application/config"
)

func main() {
	_windowConfig := config.WindowConfig{
		Title:  "My first window",
		Width:  800,
		Height: 800,
	}

	bootstrap.Run(_windowConfig)
}
