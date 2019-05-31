package main

import (
	"BatSimulator2020/core"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	g, err := core.NewGame()
	if err != nil {
		panic(err)
	}

	g.MainGameLoop()
}

func main() {
	pixelgl.Run(run)
}
