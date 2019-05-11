package main

import "github.com/Tskken/QuadGo"

const (
	MaxEntities = 25
)

var QGo = QuadGo.NewQuadGo(MaxEntities, QuadGo.NewBounds(0,0,WindowWidth, WindowHeight))