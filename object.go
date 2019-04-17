package main

import (
	"github.com/faiface/pixel"
)

type Collideable uint8

type ObjectCollection map[Collideable][]Object

const (
	CollideTrue Collideable = iota
	CollideFalse
)

type Object struct {
	Position   pixel.Vec
	Bounds     pixel.Rect
	Sprite     *pixel.Sprite
}

func NewObject(w, h float64, pos pixel.Vec, sprite *pixel.Sprite) Object {
	return Object{
		Position:pos,
		Bounds:pixel.R(0,0, w, h).Moved(pixel.V(w/2,h/2)),
		Sprite:sprite,
	}
}

func (o *Object) Draw(target pixel.Target) {
	o.Sprite.Draw(target, pixel.IM.Moved(o.Position))
}

func (obj ObjectCollection) Draw(target pixel.Target) {
	for _, obs := range obj {
		for _, o := range obs {
			o.Draw(target)
		}
	}
}