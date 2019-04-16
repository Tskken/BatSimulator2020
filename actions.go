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

func (g *Game) ActionHandler() {
	switch {
	case g.Window.Pressed(pixelgl.KeyW):
		g.Bat.Position.Y += g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position) {
			g.Camera.Position.Y += g.Camera.Speed * g.DT
			g.Camera.Bounds.Max.Y += g.Camera.Speed * g.DT
			g.Camera.Bounds.Min.Y += g.Camera.Speed * g.DT
		}

		g.Animation.Update(Up)
	case g.Window.Pressed(pixelgl.KeyS):
		g.Bat.Position.Y -= g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position){
			g.Camera.Position.Y -= g.Camera.Speed * g.DT
			g.Camera.Bounds.Max.Y -= g.Camera.Speed * g.DT
			g.Camera.Bounds.Min.Y -= g.Camera.Speed * g.DT
		}

		g.Animation.Update(Down)
	case g.Window.Pressed(pixelgl.KeyD):
		g.Bat.Position.X += g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position){
			g.Camera.Position.X += g.Camera.Speed * g.DT
			g.Camera.Bounds.Max.X += g.Camera.Speed * g.DT
			g.Camera.Bounds.Min.X += g.Camera.Speed * g.DT
		}

		g.Animation.Update(Right)
	case g.Window.Pressed(pixelgl.KeyA):
		g.Bat.Position.X -= g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position) {
			g.Camera.Position.X -= g.Camera.Speed * g.DT
			g.Camera.Bounds.Max.X -= g.Camera.Speed * g.DT
			g.Camera.Bounds.Min.X -= g.Camera.Speed * g.DT
		}

		g.Animation.Update(Left)
	default:
		g.Animation.Update(Idle)
	}
}
