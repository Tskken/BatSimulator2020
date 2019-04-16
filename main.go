package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	g, err := NewGame()
	if err != nil {
		panic(err)
	}

	g.MainGameLoop()
}

func main() {
	pixelgl.Run(run)
}
