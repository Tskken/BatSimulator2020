package main

import (
	"BatSimulator2020/mapdecoder"
	"encoding/hex"
	"errors"
	"github.com/Tskken/QuadGo"
	"github.com/faiface/pixel"
	"image/color"
	"math"
	"path/filepath"
)

const (
	TileSize   = 16
	WorldScale = 2.5

	FlippedHorizontallyFlag uint = 0x80000000
	FlippedVerticallyFlag   uint = 0x40000000
	FlippedDiagonallyFlag   uint = 0x20000000
)

type World struct {
	Height, Width   int
	BackgroundColor color.RGBA
	Layers          []*Layer
	Orientation     string
	TileSets        []*TileSet
	Type            string
}

type Layer struct {
	Id      int
	Name    string
	Sprites map[pixel.Rect]*Sprite
	Objects []*Object
	Opacity float64
	Type    string
	Visible bool

	Batch *pixel.Batch
}

func (w *World) NewLayer(m *mapdecoder.Map, index int) *Layer {
	l := &Layer{
		Type:    m.Layers[index].Type,
		Name:    m.Layers[index].Name,
		Visible: m.Layers[index].Visible,
		Id:      m.Layers[index].Id,
		Opacity: m.Layers[index].Opacity,
	}

	if l.Type != "tilelayer" {
		if len(m.Layers[index].Objects) != 0 {
			l.Objects = make([]*Object, len(m.Layers[index].Objects), len(m.Layers[index].Objects))
			for i, obj := range m.Layers[index].Objects {
				o := w.NewObject(obj)
				l.Objects[i] = o
				QGo.Insert(QuadGo.NewBounds(o.Rect.Min.X, o.Rect.Min.Y, o.Rect.Max.X, o.Rect.Max.Y), o)
			}
		}
	} else {
		sprites, spritesheet := LoadSpiteMap(m)

		if len(m.Layers[index].Data) != 0 {
			l.Sprites = w.TileFlipCheck(m.Layers[index].Data, sprites)
		}

		l.Batch = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	}

	return l
}

type Sprite struct {
	Sprite *pixel.Sprite
	Matrix pixel.Matrix
}

func LoadSpiteMap(m *mapdecoder.Map) ([]*pixel.Sprite, pixel.Picture) {
	spritesheet, err := LoadPicture(m.TileSets[0].Image)
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for y := spritesheet.Bounds().Max.Y - TileSize; y >= spritesheet.Bounds().Min.Y; y -= TileSize {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += TileSize {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+TileSize, y+TileSize)))
		}
	}

	return sprites, spritesheet
}

type TileSet struct {
	FirstGId              uint
	Name                  string
	TileHeight, TileWidth int
}

func NewTileSet(tile *mapdecoder.TileSet) *TileSet {
	return &TileSet{
		FirstGId:   uint(tile.FirstGId),
		Name:       tile.Name,
		TileWidth:  tile.TileWidth,
		TileHeight: tile.TileHeight,
	}
}

func GetDecodedMap() *mapdecoder.Map {
	mapConfigPath, err := filepath.Abs("./assets/maps/cave_map_v1.json")
	if err != nil {
		panic(err)
	}

	tileSetRootPath, err := filepath.Abs("../BatSimulator2020/assets/tilesets")
	if err != nil {
		panic(err)
	}

	m, err := mapdecoder.LoadMap(tileSetRootPath, mapConfigPath)
	if err != nil {
		panic(err)
	}

	return m
}

func LoadMap() *World {
	m := GetDecodedMap()

	w := &World{
		Height:      m.Height,
		Width:       m.Width,
		Layers:      make([]*Layer, len(m.Layers), len(m.Layers)),
		Orientation: m.Orientation,
		TileSets:    make([]*TileSet, len(m.TileSets), len(m.TileSets)),
		Type:        m.Type,
	}

	if m.BackgroundColor != "" {
		b, err := hex.DecodeString(m.BackgroundColor[1:])
		if err != nil {
			panic(err)
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
			panic(errors.New("improper format of hex color in LoadMap()"))
		}
	}

	for i, tile := range m.TileSets {
		w.TileSets[i] = NewTileSet(tile)
	}

	for i := range m.Layers {
		w.Layers[i] = w.NewLayer(m, i)
	}

	return w
}

func (w *World) TileFlipCheck(data []uint, sprites []*pixel.Sprite) (spriteMap map[pixel.Rect]*Sprite) {
	spriteMap = make(map[pixel.Rect]*Sprite)

	var tileIndex uint = 0

	for y := w.Height - 1; y >= 0; y-- {
		for x := 0; x < w.Width; x++ {
			point := pixel.R(float64(x)*TileSize*WorldScale, float64(y)*TileSize*WorldScale, (float64(x)*TileSize*WorldScale)+TileSize, (float64(y)*TileSize*WorldScale)+TileSize)

			globalTileId := data[tileIndex]

			// Read out the flags
			flippedHorizontally := globalTileId&FlippedHorizontallyFlag == FlippedHorizontallyFlag
			flippedVertically := globalTileId&FlippedVerticallyFlag == FlippedVerticallyFlag
			flippedDiagonally := globalTileId&FlippedDiagonallyFlag == FlippedDiagonallyFlag

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

			// Clear the flags
			globalTileId &= ^(FlippedHorizontallyFlag |
				FlippedVerticallyFlag |
				FlippedDiagonallyFlag)

			// Resolve the tile
			for i := len(w.TileSets) - 1; i >= 0; i-- {
				tileset := w.TileSets[i]
				if tileset.FirstGId <= globalTileId {
					spriteMap[point] = &Sprite{
						Sprite: sprites[globalTileId-w.TileSets[i].FirstGId],
						Matrix: matrix.Scaled(pixel.ZV, WorldScale).Moved(point.Min),
					}
					break
				}
			}
			tileIndex++
		}
	}
	return
}

func (w *World) Draw(target pixel.Target) {
	for _, layer := range w.Layers {
		if layer.Type == "tilelayer" {
			layer.Batch.Clear()

			for _, sprite := range layer.Sprites {
				sprite.Sprite.Draw(layer.Batch, sprite.Matrix)
			}

			layer.Batch.Draw(target)
		}
	}
}
