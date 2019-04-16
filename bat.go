package main

import "github.com/faiface/pixel"

type Bat struct {
	Sprite   *pixel.Sprite
	HitBox   pixel.Rect
	Position pixel.Vec
	Speed    float64
}
