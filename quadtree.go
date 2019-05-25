package main

import (
	"github.com/Tskken/QuadGo"
	"github.com/faiface/pixel"
)

const (
	MaxEntities = 25
)

var QGo, _ = QuadGo.NewQuadGo(MaxEntities, WindowWidth, WindowHeight)

func ToBounds(bounds pixel.Rect) QuadGo.Bounds {
	return QuadGo.NewBounds(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)
}
