package main

import (
	"github.com/faiface/pixel"
	"math"
)

type Camera struct {
	Matrix   pixel.Matrix
	Bounds   pixel.Rect
	Position pixel.Vec
}

func (c *Camera) UpdateCamera(vec pixel.Vec, dt float64) {
	vec.X -= c.Bounds.W() / 2
	vec.Y -= c.Bounds.H() / 2
	vec = pixel.Lerp(c.Position, vec, 1-math.Pow(1.0/128, dt))
	c.Matrix = pixel.IM.Moved(vec.Scaled(-1 / 1))
	c.Matrix = c.Matrix.Scaled(vec, 1)
	c.Position = vec
}
