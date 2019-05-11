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

func (g *Game) ActionHandler() {
	v := pixel.ZV
	var a Action
	moved := false
	if g.Window.Pressed(pixelgl.KeyW) {
		tV := v.Add(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Up
			moved = true
		}
	}

	if g.Window.Pressed(pixelgl.KeyS) {
		tV := v.Sub(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Down
			moved = true
		}
	}

	if g.Window.Pressed(pixelgl.KeyD) {
		tV := v.Add(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Right
			moved = true
		}
	}

	if g.Window.Pressed(pixelgl.KeyA) {
		tV := v.Sub(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Left
			moved = true
		}
	}

	if !moved {
		g.Animation.Update(Idle)
	} else {
		g.Bat.Moved(v)
		g.Animation.Update(a)
	}
}
