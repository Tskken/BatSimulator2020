package core

import (
	"github.com/faiface/pixel"
	"math"
	"time"
)

type Bat struct {
	Sprite *Sprite
	HitBox pixel.Rect

	State      Action
	FrameIndex float64

	AnimationTimer time.Time
}

func NewBat(winCenter pixel.Vec) *Bat {
	return &Bat{
		Sprite: &Sprite{
			Sprite: EntitySprites.BatSprites[Idle][0],
			Matrix: pixel.IM.Scaled(pixel.ZV, GlobalConfig.BatScale).Moved(winCenter),
		},
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
	}
}

func (b *Bat) CollisionCheck(vec pixel.Vec) bool {
	return QGo.IsIntersect(ToBounds(b.HitBox.Moved(vec)))
}

func (b *Bat) Update(vec pixel.Vec, action Action, dt time.Duration) {
	b.FrameIndex += dt.Seconds()

	b.State = action

	i := int(math.Floor(b.FrameIndex / GlobalConfig.BatAnimationRate.Seconds()))
	sprite := EntitySprites.BatSprites[b.State][i%len(EntitySprites.BatSprites[b.State])]

	b.Sprite.Update(sprite, vec)
	b.HitBox = b.HitBox.Moved(vec)
}

func (b *Bat) Draw(target pixel.Target) {
	b.Sprite.Draw(target)
}
