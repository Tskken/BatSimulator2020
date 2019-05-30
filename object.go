package main

import (
	"BatSimulator2020/tileddecoder"
	"fmt"
	"github.com/faiface/pixel"
)

type Object struct {
	Rect   pixel.Rect
	Sprite *Sprite
	Action func()
}

func NewObject(obj *tileddecoder.Object, m *tileddecoder.Map) *Object {
	rec := pixel.R(obj.X*WorldScale, (-obj.Y-obj.Height)*WorldScale, (obj.X+obj.Width)*WorldScale, -obj.Y*WorldScale)
	rec = rec.Moved(pixel.V(-TileSize, float64(m.Height)*TileSize*WorldScale-TileSize))
	return &Object{
		//Vec:  rec.Min,
		Rect: rec,
	}
}

func (o *Object) String() string {
	return fmt.Sprintf("Rect{%v}", o.Rect)
}
