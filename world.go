package main

import (
	"BatSimulator2020/mapdecoder"
	"encoding/hex"
	"errors"
	"github.com/faiface/pixel"
	"image/color"
	"math"
)

const (
	TileSize   = 16
	WorldScale = 2

	FlippedHorizontallyFlag uint = 0x80000000
	FlippedVerticallyFlag   uint = 0x40000000
	FlippedDiagonallyFlag   uint = 0x20000000
)

type World struct {
	Height, Width   int
	BackgroundColor color.RGBA
	Layers          []Layer
	Orientation     string
	TileSets        []TileSet
	Type            string
}

type Layer struct {
	Id      int
	Name    string
	Sprites [][]Sprite
	Objects []Object
	Opacity float64
	Type    string
	Visible bool

	Batch *pixel.Batch
}

type Sprite struct {
	Sprite *pixel.Sprite
	Matrix pixel.Matrix
}

type TileSet struct {
	FirstGId              uint
	Name                  string
	TileHeight, TileWidth int
}

type Object struct {
	GId           int
	Id            int
	Height, Width float64
	Name          string
	Point         bool
	Position      pixel.Vec
	Bounds        []pixel.Line
	Collideable   bool
}

func LoadMap(tMap *mapdecoder.Map) (*World, error) {
	w := &World{
		Height:      tMap.Height,
		Width:       tMap.Width,
		Layers:      make([]Layer, len(tMap.Layers), len(tMap.Layers)),
		Orientation: tMap.Orientation,
		TileSets:    make([]TileSet, len(tMap.TileSets), len(tMap.TileSets)),
		Type:        tMap.Type,
	}

	if tMap.BackgroundColor != "" {
		b, err := hex.DecodeString(tMap.BackgroundColor[1:])
		if err != nil {
			return nil, err
		}

		switch len(b) {
		case 4:
			w.BackgroundColor = color.RGBA{
				R: b[1],
				G: b[2],
				B: b[3],
				A: 255,
			}
		case 5:
			w.BackgroundColor = color.RGBA{
				R: b[1],
				G: b[2],
				B: b[3],
				A: b[0],
			}
		default:
			return nil, errors.New("improper format of hex color in LoadMap()")
		}
	}

	for i, layer := range tMap.Layers {
		w.Layers[i].Type = layer.Type
		w.Layers[i].Name = layer.Name
		w.Layers[i].Visible = layer.Visible
		w.Layers[i].Id = layer.Id
		w.Layers[i].Opacity = layer.Opacity
		w.Layers[i].Sprites = make([][]Sprite, w.Width, w.Width)

		if len(layer.Objects) != 0 {
			w.Layers[i].Objects = make([]Object, len(layer.Objects), len(layer.Objects))

			for j, object := range layer.Objects {
				w.Layers[i].Objects[j].Id = object.Id
				w.Layers[i].Objects[j].Name = object.Name
				w.Layers[i].Objects[j].Width = object.Width
				w.Layers[i].Objects[j].Height = object.Height
				w.Layers[i].Objects[j].Collideable = object.Properties[0].Value.(bool)
				w.Layers[i].Objects[j].Point = object.Point
				w.Layers[i].Objects[j].GId = object.GId
				w.Layers[i].Objects[j].Position = pixel.V(object.X, object.Y)

				if len(object.Polyline) != 0 {
					w.Layers[i].Objects[j].Bounds = make([]pixel.Line, len(object.Polyline)-1, len(object.Polyline)-1)

					for k := 0; k < len(object.Polyline)-1; k++ {
						w.Layers[i].Objects[j].Bounds[k] = pixel.L(pixel.V(object.Polyline[k].X, object.Polyline[k].Y), pixel.V(object.Polyline[k+1].X, object.Polyline[k+1].Y))
					}
				}
			}
		}

		spritesheet, err := LoadPicture(tMap.TileSets[0].Image)
		if err != nil {
			panic(err)
		}

		var sprites []*pixel.Sprite
		sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(0, 0, 0, 0)))

		// Save sprites from sprite sheet to array
		for y := spritesheet.Bounds().Max.Y - TileSize; y >= spritesheet.Bounds().Min.Y; y -= TileSize {
			for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += TileSize {
				sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+TileSize, y+TileSize)))
			}
		}

		if len(layer.Data) != 0 {
			w.Layers[i].Sprites = w.TileFlipCheck(layer.Data, sprites)
		}

		w.Layers[i].Batch = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	}

	return w, nil
}

func (w *World) TileFlipCheck(data []uint, sprites []*pixel.Sprite) (spriteMap [][]Sprite) {
	spriteMap = make([][]Sprite, w.Height, w.Height)
	for y := range spriteMap {
		spriteMap[y] = make([]Sprite, w.Width, w.Width)
	}

	var tileIndex uint = 0

	for y := w.Height - 1; y >= 0; y-- {
		for x := 0; x < w.Width; x++ {
			globalTileId := data[tileIndex]

			// Read out the flags
			flippedHorizontally := globalTileId&FlippedHorizontallyFlag == FlippedHorizontallyFlag
			flippedVertically := globalTileId&FlippedVerticallyFlag == FlippedVerticallyFlag
			flippedDiagonally := globalTileId&FlippedDiagonallyFlag == FlippedDiagonallyFlag

			// Clear the flags
			globalTileId &= ^(FlippedHorizontallyFlag |
				FlippedVerticallyFlag |
				FlippedDiagonallyFlag)

			// Resolve the tile
			for i := len(w.TileSets) - 1; i >= 0; i-- {
				tileset := w.TileSets[i]

				if tileset.FirstGId <= globalTileId {
					spriteMap[y][x].Sprite = sprites[globalTileId-w.TileSets[i].FirstGId]
					break
				}
			}

			matrix := pixel.IM

			if flippedDiagonally {
				matrix = matrix.Rotated(pixel.ZV, 90*math.Pi/180).ScaledXY(pixel.ZV, pixel.V(1, -1))
			}

			if flippedHorizontally {
				matrix = matrix.ScaledXY(pixel.ZV, pixel.V(-1, 1))
			}

			if flippedVertically {
				matrix = matrix.ScaledXY(pixel.ZV, pixel.V(1, -1))
			}

			spriteMap[y][x].Matrix = matrix.Scaled(pixel.ZV, WorldScale).Moved(pixel.V(float64(x)*TileSize*WorldScale, float64(y)*TileSize*WorldScale))
			tileIndex++
		}
	}
	return
}

func (w *World) Draw(target pixel.Target) {
	for _, layer := range w.Layers {
		layer.Batch.Clear()

		for _, spritex := range layer.Sprites {
			for _, sprite := range spritex {
				if sprite.Sprite != nil {
					sprite.Sprite.Draw(layer.Batch, sprite.Matrix)
				}
			}
		}

		layer.Batch.Draw(target)
	}
}
