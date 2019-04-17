package main

import (
	"github.com/faiface/pixel"
	"path/filepath"
	"time"
)

const (
	AnimationTime = time.Second / 8
	SpriteWidth = 32
	SpriteHeight = 32
	Scale = 2
)

type Animation struct {
	SpriteMap map[Action][]*pixel.Sprite
	Index     int
	Action    Action

	AnimationTimer <-chan time.Time
}

func NewAnimation() Animation {
	path, err := filepath.Abs("./assets/sprites/bat-sprite.png")
	if err != nil {
		panic(err)
	}

	// Load sprite sheet from file
	spritesheet, err := LoadPicture(path)
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+32, y+32)))
		}
	}

	return Animation{
		SpriteMap: map[Action][]*pixel.Sprite{
			Up: {
				sprites[5],
				sprites[9],
				sprites[13],
			},
			Down: {
				sprites[7],
				sprites[11],
				sprites[15],
			},
			Left: {
				sprites[4],
				sprites[8],
				sprites[12],
			},
			Right: {
				sprites[6],
				sprites[10],
				sprites[14],
			},
			Idle: {
				sprites[7],
				sprites[11],
				sprites[15],
			},
		},
		Action:         Idle,
		AnimationTimer: time.Tick(AnimationTime),
	}
}

func (a *Animation) Update(action Action) {
	if a.Action == action {
		a.Index++
		if a.Index >= len(a.SpriteMap[a.Action]) {
			a.Index = 0
		}
	} else {
		a.Action = action
	}
}
