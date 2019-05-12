package main

import (
	"github.com/Tskken/QuadGo"
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
		Matrix: pixel.IM.Scaled(pixel.ZV, BatScale).Moved(winCenter),
		Speed:  500.0,
	}
}

func (b *Bat) CollisionCheck(vec pixel.Vec) bool {
	bounds := b.HitBox.Moved(vec)
	return QGo.IsIntersect(QuadGo.NewBounds(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
}

func (b *Bat) Moved(vec pixel.Vec) {
	b.Matrix = b.Matrix.Moved(vec)
	b.HitBox = b.HitBox.Moved(vec)
}

func (b *Bat) Draw(target pixel.Target) {
	b.Sprite.Draw(target, b.Matrix)
}
