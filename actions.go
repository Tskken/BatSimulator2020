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
	a := Idle

	if g.Window.Pressed(pixelgl.KeyW) {
		tV := pixel.ZV.Add(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		if !g.Bat.CollisionCheck(v) {
			v = tV
			a = Up
		}
	}

	if g.Window.Pressed(pixelgl.KeyS) {
		tV := v.Sub(pixel.V(0, g.Bat.Speed*g.DeltaTime.DT))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Down
		}
	}

	if g.Window.Pressed(pixelgl.KeyD) {
		tV := v.Add(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Right
		}
	}

	if g.Window.Pressed(pixelgl.KeyA) {
		tV := v.Sub(pixel.V(g.Bat.Speed*g.DeltaTime.DT, 0))
		if !g.Bat.CollisionCheck(tV) {
			v = tV
			a = Left
		}
	}

	g.Bat.Moved(v)
	g.Animation.Update(a)
}
