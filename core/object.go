package core

import (
	"BatSimulator2020/tileddecoder"
	"fmt"
	"github.com/faiface/pixel"
)

const (
	TileSize = 16
)

type Object struct {
	Rect   pixel.Rect
	Sprite *Sprite
	Action func()
}

func NewObject(obj *tileddecoder.Object, m *tileddecoder.Map) *Object {
	rec := pixel.R(obj.X*GlobalConfig.WorldScale, (-obj.Y-obj.Height)*GlobalConfig.WorldScale, (obj.X+obj.Width)*GlobalConfig.WorldScale, -obj.Y*GlobalConfig.WorldScale)
	rec = rec.Moved(pixel.V(-TileSize, float64(m.Height)*TileSize*GlobalConfig.WorldScale-TileSize))
	return &Object{
		//Vec:  rec.Min,
		Rect: rec,
	}
}

func (o *Object) String() string {
	return fmt.Sprintf("Rect{%v}", o.Rect)
}
