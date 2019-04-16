package main

import (
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
		g.Bat.Position.Y += g.Bat.Speed * g.DT
		g.Animation.Update(Up)
	case g.Window.Pressed(pixelgl.KeyS):
		g.Bat.Position.Y -= g.Bat.Speed * g.DT
		g.Animation.Update(Down)
	case g.Window.Pressed(pixelgl.KeyD):
		g.Bat.Position.X += g.Bat.Speed * g.DT
		g.Animation.Update(Right)
	case g.Window.Pressed(pixelgl.KeyA):
		g.Bat.Position.X -= g.Bat.Speed * g.DT
		g.Animation.Update(Left)
	default:
		g.Animation.Update(Idle)
	}
}
