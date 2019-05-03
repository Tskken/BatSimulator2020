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
	v := pixel.V(0, 0)
	var a Action
	moved := false
	if g.Window.Pressed(pixelgl.KeyW) {
		v = v.Add(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		a = Up
		moved = true
	}

	if g.Window.Pressed(pixelgl.KeyS) {
		v = v.Sub(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		a = Down
		moved = true
	}

	if g.Window.Pressed(pixelgl.KeyD) {
		v = v.Add(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		a = Right
		moved = true
	}

	if g.Window.Pressed(pixelgl.KeyA) {
		v = v.Sub(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		a = Left
		moved = true
	}

	if !moved {
		g.Animation.Update(Idle)
	} else {
		g.Bat.Moved(v)
		g.Animation.Update(a)
	}
}
