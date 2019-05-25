package main

import (
	"github.com/faiface/pixel"
	"path/filepath"
)

const (
	SpriteWidth  = 32
	SpriteHeight = 32
	Scale        = 2
)

var EntitySprites = struct {
	BatSprites map[Action][]*pixel.Sprite
}{
	BatSprites: loadBatSprites(),
}

func loadBatSprites() map[Action][]*pixel.Sprite {
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
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += SpriteWidth {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += SpriteHeight {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+SpriteWidth, y+SpriteHeight)))
		}
	}

	return map[Action][]*pixel.Sprite{
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
	}
}
