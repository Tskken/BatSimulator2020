package core

import (
	"github.com/Tskken/QuadGo"
	"github.com/faiface/pixel"
)

const (
	MaxEntities = 25
)

var QGo *QuadGo.QuadGo

func LoadQuadGo() {
	QGo, _ = QuadGo.NewQuadGo(MaxEntities, GlobalConfig.WindowWidth, GlobalConfig.WindowHeight)
}

func ToBounds(bounds pixel.Rect) QuadGo.Bounds {
	return QuadGo.NewBounds(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)
}
