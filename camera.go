package main

import "github.com/faiface/pixel"

type Camera struct {
	Matrix   pixel.Matrix
	Bounds   pixel.Rect
	Position pixel.Vec
	Speed    float64
}

type CamDirection uint8

const (
	Positive CamDirection = iota
	Negative
)

func (c *Camera) UpdateCamera(vec pixel.Vec, direction CamDirection) {
	switch direction {
	case Positive:
		c.Bounds.Max = c.Bounds.Max.Add(vec)
		c.Bounds.Min = c.Bounds.Min.Add(vec)
	case Negative:
		c.Bounds.Max = c.Bounds.Max.Sub(vec)
		c.Bounds.Min = c.Bounds.Min.Sub(vec)
	}

	c.Position = c.Bounds.Center()

}
