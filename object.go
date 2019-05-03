package main

import (
	"BatSimulator2020/mapdecoder"
	"fmt"
	"github.com/dhconnelly/rtreego"
	"github.com/faiface/pixel"
)

type Object struct {
	Vec    pixel.Vec
	Rect   pixel.Rect
	Sprite *pixel.Sprite
}

func (w *World) NewObject(obj *mapdecoder.Object) *Object {
	rec := pixel.R(obj.X*WorldScale, (-obj.Y-obj.Height)*WorldScale, (obj.X+obj.Width)*WorldScale, -obj.Y*WorldScale)
	rec = rec.Moved(pixel.V(-TileSize, float64(w.Height)*TileSize*WorldScale-TileSize))
	return &Object{
		Vec:  rec.Min,
		Rect: rec,
	}
}

func (o *Object) Bounds() *rtreego.Rect {
	return ToRect(o.Rect.Moved(o.Vec))
}

func (o *Object) Intersects(rct pixel.Rect) bool {
	return len(RTree.SearchIntersect(ToRect(rct))) >= 1
}

func (o *Object) String() string {
	return fmt.Sprintf("Vec{%v}, Rect{%v}", o.Vec, o.Rect)
}
