package main

import (
	"github.com/faiface/pixel"
	"log"
	"path/filepath"
)

const (
	TileSize = 16
	WorldScale = 4
)

var (
	SpriteHeightCount float64 = 0
	SpriteWidthCount float64 = 0
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

	SpriteHeightCount = spritesheet.Bounds().Max.Y / float64(TileSize)
	SpriteWidthCount = spritesheet.Bounds().Max.X / float64(TileSize)

	log.Println(SpriteHeightCount)
	log.Println(SpriteWidthCount)

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += TileSize {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += TileSize {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+TileSize, y+TileSize)))
		}
	}

	var objects []Object

	lastX := 0.0

	for _, sprite := range sprites {
			objects = append(objects, NewObject(sprite.Frame().W(), sprite.Frame().H(), pixel.V(lastX, 0), sprite))
			lastX += sprite.Frame().W()+25
	}

	return World{
		Objects: map[bool][]Object{
			true:{},
			false:objects,
		},
		Batch:pixel.NewBatch(&pixel.TrianglesData{}, spritesheet),
	}
}

func (w *World) Draw(target pixel.Target) {
	w.Batch.Clear()

	w.Objects.Draw(w.Batch, WorldScale)

	w.Batch.Draw(target)
}
