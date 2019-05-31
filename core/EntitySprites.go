package core

import (
	"github.com/faiface/pixel"
	"path/filepath"
)

const (
	SpriteWidth  = 48
	SpriteHeight = 64
	Scale        = 2
)

var EntitySprites = struct {
	BatSprites map[Action][]*pixel.Sprite
}{
	BatSprites: loadBatSprites(),
}

func loadBatSprites() map[Action][]*pixel.Sprite {
	path, err := filepath.Abs("./assets/sprites/bat-NESW.png")
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
	for y := spritesheet.Bounds().Max.Y; y > spritesheet.Bounds().Min.Y; y -= SpriteHeight {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += SpriteWidth {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+SpriteWidth, y-SpriteHeight)))
		}
	}

	return map[Action][]*pixel.Sprite{
		Up: {
			sprites[0],
			sprites[1],
			sprites[2],
		},
		Right: {
			sprites[3],
			sprites[4],
			sprites[5],
		},
		Down: {
			sprites[6],
			sprites[7],
			sprites[8],
		},
		Left: {
			sprites[9],
			sprites[10],
			sprites[11],
		},

		Idle: {
			sprites[6],
			sprites[7],
			sprites[8],
		},
	}
}
