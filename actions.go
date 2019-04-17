package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Action uint8

const (
	Up Action = iota
	Down
	Left
	Right
	Idle
)

/*
	TODO: Update for future
		- Add hit box detection
*/
func (g *Game) ActionHandler() {
	switch {
	case g.Window.Pressed(pixelgl.KeyW):
		g.Bat.Moved(pixel.V(0, g.Bat.Speed*g.DT), g.Objects[CollideTrue])
		g.Animation.Update(Up)
	case g.Window.Pressed(pixelgl.KeyS):
		g.Bat.Moved(pixel.V(0, -g.Bat.Speed*g.DT), g.Objects[CollideTrue])
		g.Animation.Update(Down)
	case g.Window.Pressed(pixelgl.KeyD):
		g.Bat.Moved(pixel.V(g.Bat.Speed*g.DT, 0), g.Objects[CollideTrue])
		g.Animation.Update(Right)
	case g.Window.Pressed(pixelgl.KeyA):
		g.Bat.Moved(pixel.V(-g.Bat.Speed*g.DT, 0), g.Objects[CollideTrue])
		g.Animation.Update(Left)
	default:
		g.Animation.Update(Idle)
	}
}
