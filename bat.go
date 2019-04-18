package main

import "github.com/faiface/pixel"

type Bat struct {
	Sprite   *pixel.Sprite
	HitBox   pixel.Rect
	Position pixel.Vec
	Speed    float64
}

func NewBat(winCenter pixel.Vec) Bat {
	return Bat{
		HitBox: pixel.R(
			0,
			0,
			SpriteWidth*Scale,
			SpriteHeight*Scale,
		).Moved(
			winCenter.Sub(
				pixel.V(
					(SpriteWidth/2)*Scale,
					(SpriteHeight/2)*Scale,
				),
			),
		),
		Position: winCenter,
		Speed:    500.0,
	}
}

func (b *Bat) Moved(vec pixel.Vec, obj []Object) {
	newBox := b.HitBox.Moved(vec)

	for _, o := range obj {
		if o.Bounds.Intersect(newBox) != pixel.R(0,0,0,0){
			return
		}
	}

	b.Position = b.Position.Add(vec)
	b.HitBox = newBox
}

func (b *Bat) Draw(target pixel.Target) {
	b.Sprite.Draw(target, pixel.IM.Scaled(pixel.ZV, Scale).Moved(b.Position))
}
