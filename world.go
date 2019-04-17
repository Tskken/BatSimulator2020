package main

import (
	"github.com/faiface/pixel"
	"path/filepath"
)

const (
	TileSize = 16
)

type World struct {
	Objects ObjectCollection
	Batch *pixel.Batch
}

func NewWorld() World {
	path, err := filepath.Abs("./assets/textures/cave.png")
	if err != nil {
		panic(err)
	}

	spritesheet, err := LoadPicture(path)
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += TileSize {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += TileSize {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+TileSize, y+TileSize)))
		}
	}

	return World{
		Objects: map[Collideable][]Object{
			CollideTrue:{
				{
					Position:pixel.V(0,0),
					Bounds:pixel.R(0,0,0,0),
					Sprite:nil,
				},
			},
		},

	}
}
