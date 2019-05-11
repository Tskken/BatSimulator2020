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

func (o *Object) Center() (x, y float64) {
	return o.Rect.Center().X, o.Rect.Center().Y
}

func (o *Object) Bounds() (minX, minY, maxX, maxY float64) {
	return o.Rect.Min.X, o.Rect.Min.Y, o.Rect.Max.X, o.Rect.Max.Y
}

func (o *Object) W() float64 {
	return o.Rect.W()
}

func (o *Object) H() float64 {
	return o.Rect.H()
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
