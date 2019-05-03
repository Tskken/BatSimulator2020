package main

import (
	"github.com/dhconnelly/rtreego"
	"github.com/faiface/pixel"
)

const (
	BatScale = 2
)

type Bat struct {
	Sprite   *pixel.Sprite
	HitBox   pixel.Rect
	Position pixel.Vec
	Matrix   pixel.Matrix
	Speed    float64
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
		Position: winCenter,
		Matrix:   pixel.IM.Scaled(pixel.ZV, BatScale).Moved(winCenter),
		Speed:    500.0,
	}
}

func CollisionCheck(rec pixel.Rect) bool {
	treeObjects := RTree.NearestNeighbors(25, rtreego.Point{rec.Center().X, rec.Center().Y})
	for _, t := range treeObjects {
		if t.(*Object).Rect.Intersect(rec) != pixel.R(0, 0, 0, 0) {
			return true
		}
	}
	return false
}

func (b *Bat) Moved(vec pixel.Vec) {
	newBox := b.HitBox.Moved(vec)
	if CollisionCheck(newBox) {
		return
	}

	b.Matrix = b.Matrix.Moved(vec)
	b.Position = b.Position.Add(vec)
	b.HitBox = newBox

}

func (b *Bat) Draw(target pixel.Target) {
	b.Sprite.Draw(target, b.Matrix)
}
