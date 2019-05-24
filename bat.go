package main

import (
	"github.com/faiface/pixel"
)

const (
	BatScale = 2
)

type Bat struct {
	Sprite *pixel.Sprite
	HitBox pixel.Rect
	Matrix pixel.Matrix
	Speed  float64
}

func NewBat(winCenter pixel.Vec) *Bat {
	return &Bat{
		HitBox: pixel.R(
			0,
			0,
			(SpriteWidth/2)*Scale,
			(SpriteHeight/2)*Scale,
		).Moved(
			winCenter.Sub(
				pixel.V(
					(SpriteWidth/4)*Scale,
					(SpriteHeight/4)*Scale,
				),
			),
		),
		Matrix: pixel.IM.Scaled(pixel.ZV, BatScale).Moved(winCenter),
		Speed:  500.0,
	}
}

func (b *Bat) CollisionCheck(vec pixel.Vec) bool {
	return QGo.IsIntersect(ToBounds(b.HitBox.Moved(vec)))
}

func (b *Bat) Moved(vec pixel.Vec) {
	b.Matrix = b.Matrix.Moved(vec)
	b.HitBox = b.HitBox.Moved(vec)
}

func (b *Bat) Draw(target pixel.Target) {
	b.Sprite.Draw(target, b.Matrix)
}
