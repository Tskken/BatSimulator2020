package main

import "github.com/faiface/pixel"

type Camera struct {
	Matrix   pixel.Matrix
	Bounds pixel.Rect
	Position pixel.Vec
	Speed    float64
}
