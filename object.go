package main

import (
	"BatSimulator2020/mapdecoder"
	"fmt"
	"github.com/faiface/pixel"
)

type Object struct {
	Vec    pixel.Vec
	Rect   pixel.Rect
	Sprite *pixel.Sprite
	Action func()
}

func (w *World) NewObject(obj *mapdecoder.Object) *Object {
	rec := pixel.R(obj.X*WorldScale, (-obj.Y-obj.Height)*WorldScale, (obj.X+obj.Width)*WorldScale, -obj.Y*WorldScale)
	rec = rec.Moved(pixel.V(-TileSize, float64(w.Height)*TileSize*WorldScale-TileSize))
	return &Object{
		Vec:  rec.Min,
		Rect: rec,
	}
}

func (o *Object) String() string {
	return fmt.Sprintf("Vec{%v}, Rect{%v}", o.Vec, o.Rect)
}
