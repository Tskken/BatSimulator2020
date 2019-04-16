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

func (g *Game) ShiftX() pixel.Vec {
	return pixel.V(g.Camera.Speed*g.DT, 0)
}

func (g *Game) ShiftY() pixel.Vec {
	return pixel.V(0, g.Camera.Speed*g.DT)
}

func (g *Game) ActionHandler() {
	switch {
	case g.Window.Pressed(pixelgl.KeyW):
		g.Bat.Position.Y += g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position.Add(pixel.V(0, 50))) {
			g.Camera.UpdateCamera(g.ShiftY(), Positive)
		}

		g.Animation.Update(Up)
	case g.Window.Pressed(pixelgl.KeyS):
		g.Bat.Position.Y -= g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position.Sub(pixel.V(0, 50))) {
			g.Camera.UpdateCamera(g.ShiftY(), Negative)
		}

		g.Animation.Update(Down)
	case g.Window.Pressed(pixelgl.KeyD):
		g.Bat.Position.X += g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position.Add(pixel.V(50, 0))) {
			g.Camera.UpdateCamera(g.ShiftX(), Positive)
		}

		g.Animation.Update(Right)
	case g.Window.Pressed(pixelgl.KeyA):
		g.Bat.Position.X -= g.Bat.Speed * g.DT

		if !g.Camera.Bounds.Contains(g.Bat.Position.Sub(pixel.V(50, 0))) {
			g.Camera.UpdateCamera(g.ShiftX(), Negative)
		}

		g.Animation.Update(Left)
	default:
		g.Animation.Update(Idle)
	}
}
